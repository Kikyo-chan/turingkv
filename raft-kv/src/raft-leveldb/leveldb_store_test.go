package raft_leveldb

import (
	"fmt"
	"github.com/hashicorp/raft"
	"testing"
)

func TestStoreLog(t *testing.T){
	//db, err := leveldb.OpenFile(DB_PATH, nil)
	//if err != nil {
	//	return
	//}
	//defer db.Close()
	logStore, err := NewLeveldbStore(DB_PATH)
	if err != nil {
		fmt.Println(err.Error())
	}
	logStore.StoreLog(&raft.Log{0,0, 0, []byte("a")})
	logStore.StoreLog(&raft.Log{1,1, 0, []byte("b")})
	logStore.StoreLog(&raft.Log{2,1, 0, []byte("c")})
	logStore.StoreLog(&raft.Log{3,1, 0, []byte("d")})

}

func TestScanAllKV(t *testing.T)  {



	logStore, err := NewLeveldbStore(DB_PATH)
	if err != nil {
		fmt.Println(err.Error())
	}

	logStore.Set([]byte("testkey"), []byte("test value"))

	for k, v := range logStore.ScanAllKV() {
		fmt.Println("key: " + k + " value: " + v)
	}

}

func TestGetLog(t *testing.T) {

	raftLog := &raft.Log{}
	logStore, err := NewLeveldbStore(DB_PATH)
	if err != nil {
		fmt.Println(err.Error())
	}
	logStore.GetLog(3, raftLog)


	fmt.Println(raftLog.Index)
	fmt.Println(raftLog.Term)
	fmt.Println(raftLog.Type)
	fmt.Println(raftLog.Data)

}

func TestGetLogIndex(t *testing.T)  {

	logStore, err := NewLeveldbStore(DB_PATH)
	if err != nil {
		fmt.Println(err.Error())
	}

	firstIndex, err := logStore.FirstIndex()
	fmt.Println(firstIndex)

	lastIndex, err := logStore.LastIndex()
	fmt.Println(lastIndex)
}

func TestIterLog(t *testing.T)  {
	logStore, err := NewLeveldbStore(DB_PATH)
	if err != nil {
		fmt.Println(err.Error())
	}

	iter := logStore.db.NewIterator(nil,nil)
	for iter.Next() {
		fmt.Println(iter.Key())
	}

}

func TestDeleteRange(t *testing.T)  {
	logStore, err := NewLeveldbStore(DB_PATH)
	if err != nil {
		fmt.Println(err.Error())
	}
	logStore.DeleteRange(1,2)
}

func TestGetUint64(t *testing.T)  {
	stateStore, err := NewLeveldbStore(DB_PATH)
	if err != nil {
		fmt.Println(err.Error())
	}
	data, err := stateStore.GetUint64([]byte("a"))
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(data)
}