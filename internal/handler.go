package internal

import (
	"code_sim/internal/query"
	"code_sim/internal/update"
	"code_sim/pb_gen"
	"context"
)

func Search(ctx context.Context, req *pb_gen.CodeSimSearchRequest) (*pb_gen.CodeSimSearchResponse, error) {
	return query.Search(ctx, req)
}

func Upload(stream pb_gen.CodeSim_UploadServer) error {
	return update.Upload(stream)
}
