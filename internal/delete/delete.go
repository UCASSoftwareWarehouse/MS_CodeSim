package delete

import (
	"code_sim/es"
	"code_sim/pb_gen"
	"context"
	"fmt"
)

func ByProject(ctx context.Context, req *pb_gen.CodeSimDeleteRequest) (*pb_gen.CodeSimDeleteResponse, error) {
	terms := map[string]string{
		"project_name": req.GetProjectName(),
		"tag": req.GetTag(),
	}
	err := es.BulkDeleteByTerms(es.CodeIndex, terms)
	if err != nil {
		err = fmt.Errorf("Delete By Project failed，err=[%s]", err)
		return nil, err
	}
	return &pb_gen.CodeSimDeleteResponse{}, nil
}