package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/turingkv/raft-kv/src/node"
	"github.com/hashicorp/raft"

	"github.com/stretchr/testify/assert"
)

// test package with one real node

var raftNode *node.RStorage

func setupNode() *node.RStorage {
	dataDir := "/tmp/test_node/"
	os.RemoveAll(dataDir)

	config := node.Config{
		BindAddress:    "127.0.0.1:6666",
		NodeIdentifier: "127.0.0.1:6666",
		JoinAddress:    "127.0.0.1:6666",
		DataDir:        dataDir,
		Bootstrap:      true,
	}
	storage, err := node.NewRStorage(&config)
	if err != nil {
		log.Panic(err)
	}
	return storage
}

func performRequest(r http.Handler, method, path string, body io.Reader) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, body)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func assertValue(t *testing.T, w *httptest.ResponseRecorder, expectedValue string) {
	assert.Equal(t, http.StatusOK, w.Code, "Response code should be 200")

	var response map[string]string
	json.Unmarshal([]byte(w.Body.String()), &response)
	// Grab the value & whether or not it exists
	value, _ := response["value"]
	assert.Equal(t, expectedValue, value, "Values should be equal")
}

func getLeaderNode() *node.RStorage {
	raftNode := setupNode()
	startedAt := time.Now().Unix()
	for true == true {
		// wait until node will become leader
		if raftNode.RaftNode.State() == raft.Leader {
			break
		}
		time.Sleep(time.Millisecond * 100)
		if time.Now().Unix()-startedAt > 5 {
			log.Panicln("Node can't become a leader!")
		}
	}
	return raftNode
}

func TestGetValue(t *testing.T) {
	router := setupRouter(raftNode)
	testKey := "test-key"
	testValue := "test-value"
	url := fmt.Sprintf("/keys/%s/", testKey)

	// kv storage must be empty before the test
	assert.Equal(t, "", raftNode.Get(testKey), "KV storage must be empty before the test")

	// check that GET with empty storage returns 200
	w := performRequest(router, "GET", url, nil)
	assertValue(t, w, "")

	// set value and then get it with http request
	err := raftNode.Set(testKey, testValue)
	assert.Nil(t, err, "Can't write to the node")
	time.Sleep(time.Millisecond * 100) // wait for value to be applied

	w = performRequest(router, "GET", url, nil)
	assertValue(t, w, testValue)
}

func TestSetValueViaHTTP(t *testing.T) {
	router := setupRouter(raftNode)
	testKey := "test-key2"
	testValue := "test-value2"
	url := fmt.Sprintf("/keys/%s/", testKey)

	// kv storage must be empty before the test
	assert.Equal(t, "", raftNode.Get(testKey), "KV storage must be empty before the test")

	// check that GET with empty storage returns 200
	w := performRequest(router, "GET", url, nil)
	assertValue(t, w, "")

	// set value and then get it with http request

	jsonStr := []byte(fmt.Sprintf("{\"value\": \"%s\"}", testValue))
	w = performRequest(router, "POST", url, bytes.NewBuffer(jsonStr))

	time.Sleep(time.Millisecond * 100) // wait for value to be applied

	w = performRequest(router, "GET", url, nil)
	assertValue(t, w, testValue)
}

func TestSetKey(t *testing.T)  {

	jsonStr := []byte(`{"value":"testkey"}`)
	url := fmt.Sprintf("http://%s/keys/%s/", "127.0.0.1:8080", "some-key")
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		 fmt.Errorf("forward request to leader error %s", err.Error())
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		 fmt.Errorf("forward request to leader error %s", err.Error())
	}
	defer resp.Body.Close()

}

func TestGetKeyValue(t *testing.T) {

	resp, err := http.Get(fmt.Sprintf("http://%s/keys/%s/", "127.0.0.1:8080", "some-key"))
	if err != nil {
		fmt.Errorf("forward request to leader error %s", err.Error())
	}
	defer resp.Body.Close()
	s,err:=ioutil.ReadAll(resp.Body)
	fmt.Printf(string(s))

}


func init() {
	raftNode = getLeaderNode()
}
