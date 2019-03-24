package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
)

type KV struct {
	Name   string
	Value  string
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func PostKVToAnNode(node_ip string, node_port int, key string, json_value string) string {
	//转发set请求到leader
	jsonStr := []byte(`{"value":"`+json_value+`"}`)

	url := fmt.Sprintf("http://%s:%d/keys/%s/", node_ip, node_port , key)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		return "forward request to leader error"
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "forward request to leader error %s"
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return string(body)

}

func GetValueByKeyFromNode(node_ip string, node_port int, key string) string {

	resp, err := http.Get(fmt.Sprintf("http://%s:%d/keys/%s/", node_ip, node_port, key))
	if err != nil {
		fmt.Errorf("forward request to leader error %s", err.Error())
	}
	defer resp.Body.Close()
	s,err:=ioutil.ReadAll(resp.Body)
	kv := &KV{}
	err = json.Unmarshal([]byte(string(s)), &kv)
	if err != nil {
		fmt.Errorf("forward request to leader error %s", err.Error())
	}

	return kv.Value

}