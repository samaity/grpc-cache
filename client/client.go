package client

import (
	"context"
	"fmt"
	"github.com/samaity/grpc-cache/api/cachepb"
	"google.golang.org/grpc"
)

func runClient() error {
	// connect
	conn, err := grpc.Dial("localhost:5051", grpc.WithInsecure())
	if err != nil {
		return fmt.Errorf("failed to dial server: %v", err)
	}
	cc := cachepb.NewCacheClient(conn)
	// Put
	_, err = cc.Put(context.Background(), &cachepb.PutReq{Key: "gopher", Val: []byte("con")})
	if err != nil {
		return fmt.Errorf("failed to store: %v", err)
	}
	//Get
	resp, err := cc.Get(context.Background(), &cachepb.GetReq{Key: "gopher"})
	if err != nil {
		return fmt.Errorf("failed to get: %v", err)
	}
	fmt.Printf("Got cached value %s\n", resp.Val)
    return nil
	
}


func clientMain() {

    if err := runClient(); err != nil {
        fmt.Fprintf(os.Stderr, "failed: %v\n", err)
        os.Exit(1)
    }
}