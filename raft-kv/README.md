- Leader 启动

type Opts struct {
        BindAddress string `long:"bind" env:"BIND" default:"10.0.2.5:3000" description:"ip:port to bind for a node"`
        JoinAddress string `long:"join" env:"JOIN" default:"" description:"ip:port to join for a node"`
        Bootstrap   bool   `long:"bootstrap" env:"BOOTSTRAP" description:"bootstrap a cluster"`
        DataDir     string `long:"datadir" env:"DATA_DIR" default:"/tmp/data/" description:"Where to store system data"`
}

config := node.Config{
        BindAddress:    opts.BindAddress,
        NodeIdentifier: opts.BindAddress,
        JoinAddress:    opts.JoinAddress,
        DataDir:        opts.DataDir,
        Bootstrap:      true,
}

- Follower 启动

10.0.2.4:3000 为要加入集群的服务器地址

type Opts struct {
        BindAddress string `long:"bind" env:"BIND" default:"10.0.2.4:3000" description:"ip:port to bind for a node"`
        JoinAddress string `long:"join" env:"JOIN" default:"10.0.2.5.8080" description:"ip:port to join for a node"`
        Bootstrap   bool   `long:"bootstrap" env:"BOOTSTRAP" description:"bootstrap a cluster"`
        DataDir     string `long:"datadir" env:"DATA_DIR" default:"/tmp/data/" description:"Where to store system data"`
}

config := node.Config{
        BindAddress:    opts.BindAddress,
        NodeIdentifier: opts.BindAddress,
        JoinAddress:    opts.JoinAddress,
        DataDir:        opts.DataDir,
        Bootstrap:      opts.Bootstrap,
}
