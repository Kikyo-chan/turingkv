package zk_utils

import (
	"fmt"
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
		nodes, err := client.GetNodes("group_-1")
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

