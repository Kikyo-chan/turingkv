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

