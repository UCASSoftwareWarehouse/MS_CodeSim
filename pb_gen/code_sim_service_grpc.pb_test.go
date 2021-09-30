package pb_gen

import (
	"context"
	"google.golang.org/grpc"
	"testing"
)

func TestNewCodeSimClient(t *testing.T) {
	conn, err := grpc.Dial("localhost:4401", grpc.WithInsecure())
	if err != nil {
		t.Fatalf("err=[%v]", err)
	}
	defer conn.Close()
	cli := NewCodeSimClient(conn)
	res, err := cli.HelloWorld(context.Background(), &CodeSimHelloWorldRequest{
		HelloText:     "Hello!!!",
	})
	cli.Search(context.Background(), &CodeSimSearchRequest{
		MatchText:       "if __name__ == '__main__: '",
		MatchTextIsCode: true,
		CodeType:        CodeSimSearchRequest_python,
		Limit:           10,
		Offset:          0,
	})
	t.Logf("res=[%v], err=[%v]", res, err)
}
