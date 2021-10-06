package internal

import (
	"code_sim/internal/data"
	"code_sim/internal/query"
	"code_sim/pb_gen"
	"context"
)

func Search(ctx context.Context, req *pb_gen.CodeSimSearchRequest) (*pb_gen.CodeSimSearchResponse, error) {
 	return query.SearchCode(ctx, req)
}

func Upload(stream pb_gen.CodeSim_UploadServer) error {
	return data.SourceCodeUploader.DoUpload(stream)
}