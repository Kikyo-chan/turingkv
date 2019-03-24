package data

import (
	"encoding/json"
	"fmt"
	"github.com/turingkv/raft-kv/src/utils"
	"io/ioutil"
	"os"
	"strconv"
	"testing"
)

type T_Value struct {
	Value  string `json:"value"`
}


func TestGenData(t *testing.T) {

	value := T_Value{
		Value: "test value kv",
	}

	for i := 0 ; i < 100 ; i ++ {

		value.Value = utils.RandStringBytes(4096)
		byte_, err := json.Marshal(value)
		if err != nil {
			fmt.Println("error:", err)
		}
		//生成json文件
		err = ioutil.WriteFile(fmt.Sprintf("test.%s.json", strconv.Itoa(i)), byte_, os.ModePerm)
		if err != nil {
			return
		}

	}

}