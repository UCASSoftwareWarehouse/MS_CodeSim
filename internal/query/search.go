package query

import (
	"code_sim/es"
	"code_sim/internal/converter"
	"code_sim/pb_gen"
	"code_sim/transformer"
	"context"
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
	esIDs2Score := es.MatchCode(req.GetMatchText(), es.CodePlainTextIndex, int(req.GetOffset()), int(req.GetLimit()))
	if esIDs2Score == nil {
		return nil, status.Error(codes.Internal, "failed matching code elasticsearch")
	}
	return packSearchResponse(esIDs2Score)
}

func searchCode(ctx context.Context, req *pb_gen.CodeSimSearchRequest, codeType transformer.CodeType) (*pb_gen.CodeSimSearchResponse, error) {
	transformed, err := transformer.Transform(req.GetMatchText(), codeType)
	if err != nil {
		log.Printf("searchCode Transform encountered unsupported codeType, codeType=[%v]", codeType)
		return nil, err
	}
	esIDs2Score := es.MatchCode(transformed, es.CodeTransformedTextIndex, int(req.GetOffset()), int(req.GetLimit()))
	return packSearchResponse(esIDs2Score)
}

func packSearchResponse(esIDs2Score map[string]float64) (*pb_gen.CodeSimSearchResponse, error) {
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
	return &pb_gen.CodeSimSearchResponse{Files: files}, nil
}
