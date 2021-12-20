package converter

import (
	"code_sim/es"
	"code_sim/pb_gen"
	"fmt"
	"log"
	"strings"
)

type splitFrom int

const (
	splitFromFront splitFrom = 1
	splitFromBack  splitFrom = 2

	tagSplitWord         = ":"
	projectNameSplitWord = "@"
)

// ExtractProjectFromESID esID has a format of
func ExtractProjectFromESID(esID string) (*pb_gen.CodeSimProject, error) {
	f, err := ExtractProjectFileFromESID(esID)
	if err != nil {
		return nil, err
	}
	return f.GetProjectInfo(), nil
}

func ExtractProjectFileFromESID(esID string) (*pb_gen.CodeSimProjectFile, error) {
	tag, codeUniqueStr, err := getFirstSplitWord(esID, tagSplitWord, splitFromBack)
	if err != nil {
		log.Printf("failed converting esID 2 project file, esID=[%s]", esID)
		return nil, fmt.Errorf("invalid esID, failed to extract tag, esID=[%s]", esID)
	}
	projectName, relPath, err := getFirstSplitWord(codeUniqueStr, projectNameSplitWord, splitFromFront)
	if err != nil {
		log.Printf("failed converting esID 2 project file, esID=[%s]", esID)
		return nil, fmt.Errorf("invalid esID, failed to extract projectName and relPath, esID=[%s]", esID)
	}
	return &pb_gen.CodeSimProjectFile{
		ProjectInfo: &pb_gen.CodeSimProject{
			ProjectName: projectName,
			Tag:         tag,
		},
		RelativePath: relPath,
	}, nil
}

func ConvertToES(f *pb_gen.CodeSimProjectFile) *es.ProjectFileIdentifier {
	codeUniquePath := fmt.Sprintf("%s"+projectNameSplitWord+"%s", f.GetProjectInfo().GetProjectName(), f.GetRelativePath())
	ID := fmt.Sprintf("%s:%s", codeUniquePath, f.GetProjectInfo().GetTag())
	return &es.ProjectFileIdentifier{
		CodeUniquePath: codeUniquePath,
		ProjectIdentifier: es.ProjectIdentifier{
			ProjectName:    f.GetProjectInfo().GetProjectName(),
			Tag:            f.GetProjectInfo().GetTag(),
		},
		ID:             ID,
	}
}

func getFirstSplitWord(s, subStr string, sf splitFrom) (split string, restPiece string, err error) {
	strs := strings.Split(s, subStr)
	if len(strs) == 1 {
		return "", "", fmt.Errorf("string s=[%s] doesn't contain subStr=[%s]", s, subStr)
	}
	var restPieces []string
	if sf == splitFromFront {
		split = strs[0]
		restPieces = strs[1:]
	} else {
		split = strs[len(strs)-1]
		restPieces = strs[:len(strs)-1]
	}
	restPiece = strings.Join(restPieces, subStr)
	return
}
