package main

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/jessevdk/go-flags"
	"github.com/lestrrat/go-file-rotatelogs"
	pb "github.com/turingkv/kvrpc"
	"github.com/turingkv/raft-kv/src/node"
	"github.com/turingkv/raft-kv/src/server"
	"github.com/turingkv/raft-kv/src/utils"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
    "google.golang.org/grpc/reflection"
	"net"
    "os"
	"strconv"
	"strings"
	"time"
)

var storage *node.RStorage
var config *node.Config
var err error

func init() {
	//参数解析
	var opts Opts
	p := flags.NewParser(&opts, flags.Default)
	if _, err := p.ParseArgs(os.Args[1:]); err != nil {
		log.Panicln(err)
	}

	//日志配置
	if logf, err := rotatelogs.New(
		opts.LogPath+".%Y%m%d",
		rotatelogs.WithMaxAge(time.Hour*24*30),
		rotatelogs.WithRotationTime(time.Hour*24),
	); err != nil {
		log.WithError(err).Error("create rotatelogs, use default io.writer instead")
	} else {
		log.SetOutput(logf)
	}

	log.Infof("'%s' is used to store files of the node", opts.DataDir)

	config = &node.Config{
		BindAddress:    opts.BindAddress,
		NodeIdentifier: opts.BindAddress,
		JoinAddress:    opts.JoinAddress,
		DataDir:        opts.DataDir,
		Bootstrap:      opts.Bootstrap,
		ApiPort:        opts.ApiPort,
        RpcPort:        opts.RpcPort,
		GroupId:        opts.GroupId,
	}

	storage, err = node.NewRStorage(config)
	if err != nil {
		log.Panic(err)
	}

	msg := fmt.Sprintf("[INFO] Started node=%s", storage.RaftNode)
	log.Info(msg)


	go printStatus(storage)

    if config.JoinAddress != "" {
        for 1 == 1 {
            time.Sleep(time.Second * 1)
            err := storage.JoinCluster(config.JoinAddress)
            if err != nil {
                log.Info("Can't join the cluster: %+v", err)
            } else {
                break
            }
        }
    }

}

type RServer struct{}

func (s *RServer) PostKV(ctx context.Context, in *pb.KVRequest) (*pb.Status, error) {

	err = storage.Set(in.Key, in.Value)
	if err != nil {
		return &pb.Status{Isok: "no"}, err
	}

	return &pb.Status{Isok: "yes"}, nil
}

func (s *RServer) GetV(ctx context.Context, in *pb.VRequest) (*pb.ValueReply, error) {

	return &pb.ValueReply{Value: storage.Get(in.Key)}, nil
}

func RegisterService(){
      log.Info(config.RpcPort)
      lis, _ := net.Listen("tcp", config.RpcPort)
      s := grpc.NewServer()
      pb.RegisterApiServer(s, &RServer{})
      reflection.Register(s)
      s.Serve(lis)
}

type Opts struct {
	BindAddress string `long:"bind" env:"BIND" default:"127.0.0.1:3000" description:"ip:port to bind for a node"`
	JoinAddress string `long:"join" env:"JOIN" default:"" description:"ip:port to join for a node"`
	ApiPort     string `long:"api_port" env:"API_PORT" default:":8080" description:":port for a api port"`
    RpcPort     string `long:"rpc_port" env:"RPC_PORT" default:":8000" description:":port for a rpc port"`
	Bootstrap   bool   `long:"bootstrap" env:"BOOTSTRAP" description:"bootstrap a cluster"`
	DataDir     string `long:"data_dir" env:"DATA_DIR" default:"/tmp/data/" description:"Where to store system data"`
	GroupId     int    `long:"group_id" env:"GROUP_ID" default:"0" description:"Raft Group Id"`
	ZkAddress   string `long:"zk_address" env:"ZK_ADDRESS" default:"127.0.0.1:2181" description:"zkServerAddress"`
	LogPath     string `long:"log_path" env:"LOG_PATH" default:"logs/turing-kv.log" description:"logPath"`
}

func main() {

    var opts Opts
    p := flags.NewParser(&opts, flags.Default)
    if _, err := p.ParseArgs(os.Args[1:]); err != nil {
        log.Panicln(err)
    }

    zkServers := strings.Split(opts.ZkAddress, ",")
    client, err := utils.NewClient(zkServers, "/api", 10)
    if err != nil {
        log.Error("%v", err)
    }

    defer client.Close()
    port, err := strconv.Atoi(opts.RpcPort[1:])
    if err != nil {
        log.Error("%v", err)
    }

    node_ := &utils.ServiceNode{"group_" + strconv.Itoa(opts.GroupId), strings.Split(opts.BindAddress, ":")[0], port}
    if err := client.Register(node_); err != nil {
        log.Error(err)
    }
    

    go RegisterService()
     
    time.Sleep(time.Second * 2)
	// Start an HTTP server
	server.RunHTTPServer(storage, config.ApiPort)

}


func printStatus(s *node.RStorage) {
	for {
		log.Infof("state=%s leader=%s", s.RaftNode.State(), s.RaftNode.Leader())
		time.Sleep(time.Second * 2)
	}
}
