package update

import (
	"code_sim/es"
	"code_sim/internal/converter"
	"code_sim/pb_gen"
	"code_sim/util"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"
)

const (
	loadFileBatchSize = 10
)

func Upload(stream pb_gen.CodeSim_UploadServer) error {
	return sourceCodeUploader.doUpload(stream)
}

var sourceCodeUploader = &Uploader{}

type Uploader struct {
}

func (u *Uploader) doUpload(stream pb_gen.CodeSim_UploadServer) error {
	res, err := u.receiveStream(stream)
	if err != nil {
		log.Printf("upload chunk, failed receive stream")
		return err
	}
	extractedDir, err := u.storeToES(res.Project, res.FileType, res.FPath)
	if err != nil {
		log.Printf("upload chunk, failed while storeToES, err=[%v]", err)
		return err
	}
	defer func() {
		if res != nil && res.FPath != "" {
			_ = os.Remove(res.FPath)
		}
		if extractedDir != "" {
			_ = os.RemoveAll(extractedDir)
		}
	}()
	return nil
}

type receiveStreamRes struct {
	Project  *pb_gen.CodeSimProject
	FileType pb_gen.CodeSimUploadFileType
	FPath    string
}

func (u *Uploader) receiveStream(stream pb_gen.CodeSim_UploadServer) (*receiveStreamRes, error) {
	fo, err := os.CreateTemp("", fmt.Sprintf("upload_temp_%d", time.Now().UnixNano()))
	if err != nil {
		err = stream.SendAndClose(&pb_gen.CodeSimUploadResponse{
			Message: "Failed to create temp file",
			Status:  pb_gen.CodeSimUploadStatus_code_sim_upload_status_fail,
		})
		return nil, err
	}
	log.Printf("Uploading file to %s", fo.Name())
	// close fo on exit and check for its returned error
	defer func() {
		if err := fo.Close(); err != nil {
			log.Printf("Upload close fo failed, err=[%+v]", err)
		}
	}()
	var nonNilRes *pb_gen.CodeSimUploadRequest
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			err = stream.SendAndClose(&pb_gen.CodeSimUploadResponse{
				Message: "Upload Received",
				Status:  pb_gen.CodeSimUploadStatus_code_sim_upload_status_OK,
			})
			if err != nil {
				err = errors.New("failed to send status code")
				return nil, err
			}
			break
		}
		if nonNilRes == nil {
			nonNilRes = res
		}
		resStr, _ := json.Marshal(res)
		log.Printf("upload chunk, res=[%s]", resStr)
		// write a chunk
		if _, err := fo.Write(res.GetContentChunk()); err != nil {
			log.Printf("upload chunk, failed while writing to fo, err=[%v]", err)
			err = stream.SendAndClose(&pb_gen.CodeSimUploadResponse{
				Message: "Failed to write chunk",
				Status:  pb_gen.CodeSimUploadStatus_code_sim_upload_status_fail,
			})
			return nil, err
		}
	}
	return &receiveStreamRes{
		Project:  nonNilRes.GetProject(),
		FileType: nonNilRes.GetFileType(),
		FPath:    fo.Name(),
	}, nil
}

func (u *Uploader) storeToES(project *pb_gen.CodeSimProject, fileType pb_gen.CodeSimUploadFileType, fPath string) (string, error) {
	dir, err := u.extract(fileType, fPath)
	if err != nil {
		log.Printf("storeToES failed while extracting, err=[%v]", err)
		return dir, err
	}
	absPaths, relPaths, err := u.getPaths(dir)
	if err != nil {
		log.Printf("storeToES failed while getPaths, err=[%v]", err)
		return dir, err
	}
	log.Printf("storeToES save plainTextDocs=[%+v]", relPaths)
	err = u.saveDocs(project, absPaths, relPaths)
	if err != nil {
		log.Printf("storeToES failed while saving plain text docs, err=[%v]", err)
		return dir, err
	}
	return dir, nil
}

func (u *Uploader) getPaths(extractedDir string) (absPaths []string, relPaths []string, err error) {
	projRootDir := ""
	absPaths = make([]string, 0, 8)
	log.Printf("storeToES ready to walk dir=[%s]", extractedDir)
	err = filepath.WalkDir(extractedDir, func(path string, d fs.DirEntry, err error) error {
		log.Printf("path=[%s], fileName=[%s]", path, d.Name())
		if d.IsDir() {
			if projRootDir != "" {
				return nil
			}
			rel, _ := filepath.Rel(extractedDir, path)
			if rel == "." || rel == "" { // path is extractedDir itself, skip
				return nil
			}
			list := filepath.SplitList(rel)
			if len(list) == 1 {
				projRootDir = path
			}
			return nil
		}
		absPaths = append(absPaths, path)
		return nil
	})
	if err != nil {
		log.Printf("storeToES failed while Walking Dir, dir=[%s], err=[%v]", extractedDir, err)
		return nil, nil, err
	}
	relPaths = make([]string, 0, len(absPaths))
	for _, p := range absPaths {
		relP, err := filepath.Rel(projRootDir, p)
		if err != nil {
			log.Printf("storeToES rel path get failed, err=[%v]", err)
			return nil, nil, err
		}
		relPaths = append(relPaths, relP)
	}
	return absPaths, relPaths, nil
}

func (u *Uploader) extract(fileType pb_gen.CodeSimUploadFileType, filepath string) (string, error) {
	var mkdirTemp = func() (string, error) {
		return os.MkdirTemp("", fmt.Sprintf("upload_extracted_temp_%d", time.Now().UnixNano()))
	}
	var wrapExtract = func(f func(tempDir string) (string, error)) (string, error) {
		tempDir, err := mkdirTemp()
		if err != nil {
			return "", err
		}
		return f(tempDir)
	}
	switch fileType {
	case pb_gen.CodeSimUploadFileType_code_sim_upload_file_type_unknown:
		log.Printf("extract with unknown filetype")
		return "", errors.New("extract with unknown filetype")
	case pb_gen.CodeSimUploadFileType_code_sim_upload_file_type_zip:
		return wrapExtract(func(tempDir string) (string, error) {
			err := util.Unzip(filepath, tempDir)
			if err != nil {
				log.Printf("extract, Unzip failed, err=[%v]", err)
				return tempDir, err
			}
			return tempDir, nil
		})
	default:
		log.Printf("extract with unsupported filetype, fileType=[%d]", fileType)
		return "", errors.New("extract with unsupported filetype")
	}
}

type docGenerator func(esProjFileIdentifier *es.ProjectFileIdentifier, plainText, relPath string) es.Document

var docGenerators = map[es.IndexName]docGenerator{
	//es.CodePlainTextIndex: func(esProjFileIdentifier *es.ProjectFileIdentifier, plainText, relPath string) es.Document {
	//	return es.NewCodePlainText(plainText, esProjFileIdentifier)
	//},
	//es.CodeTransformedTextIndex: func(esProjFileIdentifier *es.ProjectFileIdentifier, plainText, relPath string) es.Document {
	//	b := path.Base(relPath)
	//	i := strings.LastIndex(b, ".")
	//	if i == -1 {
	//		return nil
	//	}
	//	suffix := b[i+1:]
	//	codeType, err := transformer.GetSupportedCodeType(suffix)
	//	if err != nil {
	//		log.Printf("encounter file suffix is not supported code type, suffix is %s, relPath=[%s]", suffix, relPath)
	//		return nil
	//	}
	//	transformed, err := transformer.Transform(plainText, codeType)
	//	if err != nil {
	//		return nil
	//	}
	//	return es.NewCodeTransformedText(transformed, esProjFileIdentifier)
	//},
	es.CodeIndex: func(esProjFileIdentifier *es.ProjectFileIdentifier, plainText, relPath string) es.Document {
		return es.NewCodePlainText(plainText, esProjFileIdentifier)
	},
}

func (u *Uploader) saveDocs(project *pb_gen.CodeSimProject, absPaths, relPaths []string) error {
	c := 0
	currBatch := 0
	relPathsBatch := make([]string, 0, loadFileBatchSize)
	absPathsBatch := make([]string, 0, loadFileBatchSize)
	for i, relP := range relPaths {
		c++
		relPathsBatch = append(relPathsBatch, relP)
		absPathsBatch = append(absPathsBatch, absPaths[i])
		if c%loadFileBatchSize != 0 && i != len(relPaths)-1 {
			continue
		}
		// 凑够了
		currBatch++
		log.Printf("saveDocs loading batch=[%d], batchSize=[%d]", currBatch, len(relPathsBatch))
		batchDocs := make(map[es.IndexName][]es.Document)
		for indexName := range docGenerators {
			batchDocs[indexName] = make([]es.Document, 0, loadFileBatchSize)
		}
		// 生成全部的batch docs
		for j, eachRelP := range relPathsBatch {
			bs, err := ioutil.ReadFile(absPathsBatch[j])
			if err != nil {
				log.Printf("saveDocs failed reading file, path=[%s]", absPathsBatch[j])
				return err
			}
			for indexName, gen := range docGenerators {
				doc := gen(converter.ConvertToES(&pb_gen.CodeSimProjectFile{
					ProjectInfo:  project,
					RelativePath: eachRelP,
				}), string(bs), eachRelP)
				if doc != nil {
					batchDocs[indexName] = append(batchDocs[indexName], doc)
				}
			}
		}
		// 对于每种indexName对应的batchDocs，将他们使用BulkIndex方法批量存入es中。
		for indexName, docs := range batchDocs {
			if len(docs) == 0 {
				log.Printf("saveDocs skipped indexName=[%s], docs len is 0", indexName)
				continue
			}
			log.Printf("saveDocs ready to call BulkIndex indexName=[%s], docs=[%+v], currBatch=[%d]", indexName, docs, currBatch)
			err := es.BulkIndex(indexName, docs)
			if err != nil {
				log.Printf("saveDocs indexName=[%s], failed when using BulkIndex, err=[%v]", indexName, err)
				return err
			}
			log.Printf("saveDocs BulkIndex success, indexName=[%s], currBatch=[%d]", indexName, currBatch)
		}
		log.Printf("saveDocs saved batch=[%d]", currBatch)
		relPathsBatch = relPathsBatch[0:0]
		absPathsBatch = absPathsBatch[0:0]
	}
	return nil
}
