package server

import (
	"fmt"
	"net"
	"github.com/samaity/grpc-cache/api/cachepb"
)


type CacheServer struct {
	store map[string][]byte

}

func (s *CacheServer) Put(_ context.Context, req *cachepb.PutReq) (*cachepb.PutResp, error) {

	s.store[req.Key] = req.Val
	return &cachepb.PutResp{}, nil
}
func (s *CacheServer) Get(_ context.Context, *cachepb.GetReq) (*cachepb.GetResp, error) {
	val := s.store[req.Key] 
	return &cachepb.GetResp{Val: val}, nil
}

func runServer() error {
    srv := grpc.NewServer()
	cachepb.RegisterCacheServer(srv, &CacheServer{make(map[string][]byte)})
	lis, err := net.Listen("tcp", "localhost:5051")
    if err != nil {
        return err
    }
    return srv.Serve(lis)
}


func serverMain() {
    if err := runServer(); err != nil {
        fmt.Fprintf(os.Stderr, "Failed to run cache server: %s\n", err)
        os.Exit(1)
    }
}