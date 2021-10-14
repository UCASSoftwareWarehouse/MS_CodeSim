package query

import (
	"code_sim/config"
	"code_sim/es"
	"code_sim/internal/converter"
	"code_sim/pb_gen"
	"context"
	"encoding/json"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

var codeTypeMap = map[pb_gen.CodeSimSearchRequest_CodeType]es.CodeIndexField{
	pb_gen.CodeSimSearchRequest_python: es.CodeIndexFieldJava,
	pb_gen.CodeSimSearchRequest_golang: es.CodeIndexFieldGolang,
}

func Search(ctx context.Context, req *pb_gen.CodeSimSearchRequest) (*pb_gen.CodeSimSearchResponse, error) {
	reqStr, _ := json.Marshal(req)
	log.Printf("Search req=[%s]", reqStr)

	for _, codeType := range req.GetCodeTypes() {
		if _, ok := codeTypeMap[codeType]; !ok {
			log.Printf("encountered none supported codeType %v", codeType)
			return nil, status.Errorf(codes.InvalidArgument,
				fmt.Sprintf("encountered none supported codeType %v", codeType))
		}
	}
	if config.Conf.MaxSearchDepth > 0 && int(req.GetOffset()) > config.Conf.MaxSearchDepth {
		return nil, status.Errorf(codes.InvalidArgument,
			fmt.Sprintf("Only Showing a Maximum of %d files", config.Conf.MaxSearchDepth))
	}
	fields := []es.CodeIndexField{es.CodeIndexFieldPlain}
	for _, codeType := range req.GetCodeTypes() {
		fields = append(fields, codeTypeMap[codeType])
	}
	res := es.GetQuery().MatchCode(&es.MatchCodeParams{
		TargetCode:   req.GetMatchText(),
		TargetFields: fields,
		From:         int(req.GetOffset()),
		Size:         int(req.GetLimit()),
		WithSource:   req.GetWithSource(),
	})
	return packSearchResponse(res)
}

func packSearchResponse(searchRes []*es.MatchCodeEachRes) (*pb_gen.CodeSimSearchResponse, error) {
	files := make([]*pb_gen.CodeSimProjectFile, 0, len(searchRes))
	for _, each := range searchRes {
		f, err := converter.ExtractProjectFileFromESID(each.ID)
		if err != nil {
			return nil, status.Error(codes.Internal, "failed converting esID 2 project file")
		}
		f.Content = []byte(each.PlainText)
		files = append(files, f)
	}
	return &pb_gen.CodeSimSearchResponse{Files: files}, nil
}
