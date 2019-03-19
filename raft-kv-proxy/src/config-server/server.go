package config_server

import "github.com/turingkv/raft-kv/src/raft-leveldb"

type Config struct {
	ServerMetaDataDir string
	RouterMetaDataDir string
}

type ConfigServer struct {

	/*
		服务器信息存储
	 */
	serversData *raft_leveldb.LeveldbStore

	/*
		key路由表信息存储
	*/
	routerData *raft_leveldb.LeveldbStore

	//版本
	version string

}
