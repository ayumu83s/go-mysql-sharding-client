# go-mysql-sharding-client
database management prompt for sharded databases

### Install
```
git clone git@github.com:ayumu83s/go-mysql-sharding-client.git
dep ensure
```

### Usage
```$xslt
go run main --config /path/to/config.toml
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