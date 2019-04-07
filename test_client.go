package main

import (
    "golang.org/x/net/context"
    "google.golang.org/grpc"
    pb "github.com/turingkv/kvrpc"
    "log"
)

func main(){
    conn, err := grpc.Dial("127.0.0.1:8000", grpc.WithInsecure())
        if err != nil {
        log.Fatalf("did not connect: %v", err)
    }
    defer conn.Close()
    c := pb.NewApiClient(conn)

    r, err := c.PostKV(context.Background(), &pb.KVRequest{Key:"key", Value:"hello turingturing"})
    if err != nil {
        log.Fatalf("could not greet: %v", err)
    }
    log.Printf("Set Value Status: %s", r.Isok)

    r_, err := c.GetV(context.Background(), &pb.VRequest{Key:"key"})
    if err != nil {
        log.Fatalf("could not greet: %v", err)
    }
    log.Printf("Get Value: %s", r_.Value)
}
