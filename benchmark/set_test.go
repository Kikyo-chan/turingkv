package benchmark

import (
	"fmt"
	"github.com/turingkv/raft-kv/src/utils"
	"io/ioutil"
	"os/exec"
	"strconv"
	"testing"
)

func TestSetData(t *testing.T) {

	for i := 0 ; i < 10 ; i++ {
		cmd := exec.Command("/bin/bash", "-c", `ab -n 100 -c 10 -T application/json -p data/test."`+strconv.Itoa(i)+`".json  http://127.0.0.1:9988/keys/"`+utils.RandStringBytes(1024)+`"/`)

		//创建获取命令输出管道
		stdout, err := cmd.StdoutPipe()
		if err != nil {
			fmt.Printf("Error:can not obtain stdout pipe for command:%s\n", err)
			return
		}

		//执行命令
		if err := cmd.Start(); err != nil {
			fmt.Println("Error:The command is err,", err)
			return
		}

		//读取所有输出
		bytes, err := ioutil.ReadAll(stdout)
		if err != nil {
			fmt.Println("ReadAll Stdout:", err.Error())
			return
		}

		if err := cmd.Wait(); err != nil {
			fmt.Println("wait:", err.Error())
			return
		}
		fmt.Printf("stdout:\n\n %s", bytes)
	}

}
