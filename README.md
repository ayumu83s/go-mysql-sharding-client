# go-mysql-sharding-client
database management prompt for sharded databases

### Install
```
wget https://github.com/ayumu83s/go-mysql-sharding-client/releases/download/v1.0.0/go-mysql-sharding-client_1.0.1_linux_amd64.tar.gz
```

### Usage
```$xslt
mysql-shard --config /path/to/config.toml
```
### Config Sample
```$xslt
[[Databases]]
database = "shard1"
host = "localhost"
port = 3306

[[Databases]]
database = "shard2"
host = "localhost"
port = 3306
user = "user"
password = "password"

[DatabaseCommon]
user = "root"
password = ""
```