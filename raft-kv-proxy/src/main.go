package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jessevdk/go-flags"
	"github.com/turingkv/raft-kv-proxy/src/hash"
	"log"
	"os"
)

const MOD  = 2048

type Opts struct {
	ApiPort   string `long:"api_port" env:"API_PORT" default:":8080" description:":port for a api port"`
	GroupCount  int   `long:"group_count" env:"GROUP_COUNT" default:"0" description:"Raft Group Count"`
}

func setKeyView(opts Opts) func(*gin.Context) {
	view := func(c *gin.Context) {
		key := c.Param("key")
		slot := hash.Crc32IEEE([]byte(key), MOD)
		fmt.Println("SLOT", slot)
		base := uint32(MOD / opts.GroupCount)
		groupId := slot / base
		fmt.Println(groupId)
	}
	return view
}

func main()  {

	var opts Opts
	p := flags.NewParser(&opts, flags.Default)
	if _, err := p.ParseArgs(os.Args[1:]); err != nil {
		log.Panicln(err)
	}

	router := gin.Default()

	router.GET("/", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.POST("/keys/:key/", setKeyView(opts))

	router.Run(opts.ApiPort)
}
