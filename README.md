# go-mysql-sharding-client
database management prompt for sharded databases

### Install
Get release binary file for your environment

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