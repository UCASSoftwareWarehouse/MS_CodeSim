package query

import (
	"code_sim/es"
	"code_sim/internal/converter"
	"code_sim/pb_gen"
	"code_sim/transformer"
	"context"
	"encoding/json"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"sort"
)

const (
	maxSearchDepth = 100
)

var codeTypeMap = map[pb_gen.CodeSimSearchRequest_CodeType]transformer.CodeType{
	pb_gen.CodeSimSearchRequest_python: transformer.Python,
	pb_gen.CodeSimSearchRequest_golang: transformer.Golang,
}

func Search(ctx context.Context, req *pb_gen.CodeSimSearchRequest) (*pb_gen.CodeSimSearchResponse, error) {
	reqStr, _ := json.Marshal(req)
	log.Printf("Search req=[%s]", reqStr)
	if _, ok := codeTypeMap[req.GetCodeType()]; !ok {
		log.Printf("encountered none supported codeType %v, use plain text search", req.GetCodeType())
		return searchPlainText(ctx, req)
	}
	if req.GetOffset() > maxSearchDepth {
		return nil, status.Errorf(codes.InvalidArgument,
			fmt.Sprintf("Only Showing a Maximum of %d files", maxSearchDepth))
	}
	return searchCode(ctx, req, codeTypeMap[req.GetCodeType()])
}

func searchPlainText(ctx context.Context, req *pb_gen.CodeSimSearchRequest) (*pb_gen.CodeSimSearchResponse, error) {
	esIDs2Score := es.GetQuery().MatchCodeIDs(req.GetMatchText(), es.CodePlainTextIndex, int(req.GetOffset()), int(req.GetLimit()))
	if esIDs2Score == nil {
		return nil, status.Error(codes.Internal, "failed matching code elasticsearch")
	}
	return packSearchResponse(esIDs2Score, req.GetWithSource())
}

func searchCode(ctx context.Context, req *pb_gen.CodeSimSearchRequest, codeType transformer.CodeType) (*pb_gen.CodeSimSearchResponse, error) {
	transformed, err := transformer.Transform(req.GetMatchText(), codeType)
	if err != nil {
		log.Printf("searchCode Transform encountered unsupported codeType, codeType=[%v]", codeType)
		return nil, err
	}
	esIDs2Score := es.GetQuery().MatchCodeIDs(transformed, es.CodeTransformedTextIndex, int(req.GetOffset()), int(req.GetLimit()))
	// aggregate plain result
	if len(esIDs2Score) < int(req.GetLimit()) {
		log.Printf("transformed res not enough, aggragate plain text result, transformed res length=[%d], offset=[%d], limit=[%d]", len(esIDs2Score), req.GetOffset(), req.GetLimit())
		plainEsIDs2Score := es.GetQuery().MatchCodeIDs(req.GetMatchText(), es.CodePlainTextIndex, int(req.GetOffset()), int(req.GetLimit()))
		for esID, score := range plainEsIDs2Score {
			if len(esIDs2Score) >= int(req.GetLimit()) {
				break
			}
			if _, ok := esIDs2Score[esID]; ok {
				continue
			}
			esIDs2Score[esID] = -1. / score
		}
	}
	log.Printf("searchCode result esID2Score=[%+v]", esIDs2Score)
	return packSearchResponse(esIDs2Score, req.GetWithSource())
}

func packSearchResponse(esIDs2Score map[string]float64, withSource bool) (*pb_gen.CodeSimSearchResponse, error) {
	files := make([]*pb_gen.CodeSimProjectFile, 0, len(esIDs2Score))
	for esID := range esIDs2Score {
		f, err := converter.ExtractProjectFileFromESID(esID)
		if err != nil {
			return nil, status.Error(codes.Internal, "failed converting esID 2 project file")
		}
		files = append(files, f)
	}
	sort.Slice(files, func(i, j int) bool {
		info1 := converter.ConvertToES(files[i])
		info2 := converter.ConvertToES(files[j])
		return esIDs2Score[info1.ID] > esIDs2Score[info2.ID]
	})

	if withSource {
		packWithSource(esIDs2Score, files)
	}

	return &pb_gen.CodeSimSearchResponse{Files: files}, nil
}

func packWithSource(esIDs2Score map[string]float64, files []*pb_gen.CodeSimProjectFile) {
	esIDs := make([]string, 0, len(esIDs2Score))
	for esID := range esIDs2Score {
		esIDs = append(esIDs, esID)
	}
	esID2Code, err := GetCodeByIDs(esIDs)
	if err != nil {
		log.Printf("packWithSource GetCodeByIDs failed, err=[%v]", err)
		return
	}
	for _, file := range files {
		esInfo := converter.ConvertToES(file)
		if code, ok := esID2Code[esInfo.ID]; ok {
			file.Content = code
			continue
		}
		log.Printf("file need to pack with source but cannot retrive source, projectFileInfo=[%+v]", file)
	}
}