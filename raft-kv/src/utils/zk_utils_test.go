package utils

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/lestrrat/go-file-rotatelogs"
	"math/rand"
	"testing"
	"time"
)

func TestServiceDiscovery(t *testing.T) {

	servers := []string{"127.0.0.1:2181"}
	client, err := NewClient(servers, "/api", 1)
	if err != nil {
		panic(err)
	}

	defer client.Close()
	//node1 := &ServiceNode{"user", "127.0.0.1", 4000}

	//if err :s= client.Register(node1); err != nil {
	//	panic(err)
	//}


	for {
		nodes, err := client.GetNodes("group_0")
		if err != nil {
			panic(err)
		}
		for _, node := range nodes {
			fmt.Println(node.Host, node.Port)
		}
		fmt.Println("Waiting...")
		time.Sleep(time.Second * 1)
	}

}


func TestServiceDiscovery2(t *testing.T) {

	servers := []string{"127.0.0.1:2181"}
	client, err := NewClient(servers, "/api", 1)
	if err != nil {
		panic(err)
	}

	defer client.Close()
	//node1 := &ServiceNode{"user", "127.0.0.1", 4000}

	//if err :s= client.Register(node1); err != nil {
	//	panic(err)
	//}


	for {
		nodes, err := client.GetNodes("group_1")
		if err != nil {
			panic(err)
		}
		for _, node := range nodes {
			fmt.Println(node.Host, node.Port)
		}

		fmt.Println(rand.Intn(2))

		fmt.Println("Waiting...")
		time.Sleep(time.Second * 1)
	}

}

func TestStringsSplit(t *testing.T) {

	//fmt.Println(strings.Split("127.0.0.1:2181,127.0.0.1:2182", ","))
	if logf, err := rotatelogs.New(
		"turing-kv.log" + ".%Y%m%d",
		rotatelogs.WithMaxAge(time.Hour*24*30),
		rotatelogs.WithRotationTime(time.Hour*24),
	); err != nil {
		log.WithError(err).Error("create rotatelogs, use default io.writer instead")
	} else {
		log.SetOutput(logf)
	}

	log.Infof("test %s", "hello")

}

