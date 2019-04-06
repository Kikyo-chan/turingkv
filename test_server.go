package main

import (
    "golang.org/x/net/context"
    "google.golang.org/grpc"
    "google.golang.org/grpc/reflection"
    pb "github.com/turingkv/kvrpc"
    "net"
)

type RServer struct{}

func (s *RServer) PostKV(ctx context.Context, in *pb.KVRequest) (*pb.Status, error) {

    return &pb.Status{Isok: "yes"}, nil
}

func (s *RServer) GetV(ctx context.Context, in *pb.VRequest) (*pb.ValueReply, error) {

    //return &pb.ValueReply{Value: storage.Get(in.Key)}
    return &pb.ValueReply{Value: ""}, nil
}

func main(){
      lis, _ := net.Listen("tcp", ":8000")
      s := grpc.NewServer()
      pb.RegisterApiServer(s, &RServer{})
      reflection.Register(s)
      s.Serve(lis)
}
