package match

import (
	"code_sim/pb_gen"
	"log"
)

func SearchCode(req *pb_gen.CodeSimSearchRequest) (*pb_gen.CodeSimSearchResponse, error) {
	if !req.GetMatchTextIsCode() {
		return searchPlainText(req)
	}
	switch req.GetCodeType() {
	case pb_gen.CodeSimSearchRequest_python:
		return searchPythonCode(req)
	default:
		log.Fatalf("SearchCode encountered unsupported code type, code type = [%d]", req.GetCodeType())
	}
	return nil, nil
}

func searchPlainText(req *pb_gen.CodeSimSearchRequest) (*pb_gen.CodeSimSearchResponse, error) {
	return nil, nil
}

func searchPythonCode(req *pb_gen.CodeSimSearchRequest) (*pb_gen.CodeSimSearchResponse, error) {
	return nil, nil
}
