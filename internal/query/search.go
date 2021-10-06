package query

import (
	"code_sim/es"
	"code_sim/internal/converter"
	"code_sim/pb_gen"
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

func SearchCode(ctx context.Context, req *pb_gen.CodeSimSearchRequest) (*pb_gen.CodeSimSearchResponse, error) {
	if !req.GetMatchTextIsCode() {
		return searchPlainText(ctx, req)
	}
	if req.GetOffset() > maxSearchDepth {
		return nil, status.Errorf(codes.InvalidArgument,
			fmt.Sprintf("Only Showing a Maximum of %d files", maxSearchDepth))
	}
	switch req.GetCodeType() {
	case pb_gen.CodeSimSearchRequest_python:
		return searchPythonCode(ctx, req)
	default:
		log.Printf("SearchCode encountered unsupported code type, code type = [%d]", req.GetCodeType())
		return nil, status.Errorf(codes.InvalidArgument,
			"invalid code type")
	}
}

func searchPlainText(ctx context.Context, req *pb_gen.CodeSimSearchRequest) (*pb_gen.CodeSimSearchResponse, error) {
	esIDs2Score := es.MatchCode(req.GetMatchText(), es.CodePlainTextIndex, int(req.GetOffset()), int(req.GetLimit()))
	if esIDs2Score == nil {
		return nil, status.Error(codes.Internal, "failed matching code elasticsearch")
	}
	files := make([]*pb_gen.CodeSimProjectFile, 0, len(esIDs2Score))
	for esID := range esIDs2Score {
		f, err := converter.ExtractProjectFileFromESID(esID)
		if err != nil {
			return nil, status.Error(codes.Internal, "failed converting esID 2 project file")
		}
		files = append(files, f)
	}
	sort.Slice(files, func(i, j int) bool {
		_, _, ID1 := converter.GenEsInfo(files[i])
		_, _, ID2 := converter.GenEsInfo(files[j])
		return esIDs2Score[ID1] > esIDs2Score[ID2]
	})
	return &pb_gen.CodeSimSearchResponse{Files: files}, nil
}

func searchPythonCode(ctx context.Context, req *pb_gen.CodeSimSearchRequest) (*pb_gen.CodeSimSearchResponse, error) {
	return nil, nil
}
