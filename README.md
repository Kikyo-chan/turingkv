# turingkv介绍

turingkv 是一个基于Raft一致性算法的分布式kv存储系统，使用 leveldb 作为存储引擎

**运行单机测试**

- 编译，进入raft-kv根目录

```
sh build.sh
```

- 运行

```
sh run.sh
```

- 设置key值

```
curl 'http://leader地址:leader api端口/keys/some-key/' -H 'Content-Type: application/json' -d '{"value": "some-value"}'
```

- 获取key值

```
curl 'http://leader地址:leader api端口/keys/some-key/'
```

**使用docker运行测试**

- 获取镜像

```
docker pull cxspace/turingkv:v1
```

- 启动容器

```
docker run -it -d b546998a9c04(使用docker images查看镜像ID) /bin/bash
```

- 进入容器

```
docker exec -it b546998a9c04(使用docker ps查看容器ID) /bin/bash
```

- 进入项目根目录

```
cd /root/go/src/github.com/turingkv/raft-kv
```

- 启动系统

```
sh run.sh
```

- 运行测试

```
sh test_case.sh
```

- 停止系统

```
sh stop.sh
```

