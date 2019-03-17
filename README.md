# go-mysql-sharding-client
database management prompt for sharded databases

### Install
Get release binary file for your environment

### Usage
```$xslt
mysql-shard --config /path/to/config.toml
```

```sql
$ ./mysql-shard --config ./config.toml
Connection ok(shard1)
Connection ok(shard2)
Please use `exit` or `Ctrl-D` to exit this program.
sharding > SELECT * FROM user ;
+----+------+-----+------------+------------+----------+
| id | name | age | created_on | updated_on | database |
+----+------+-----+------------+------------+----------+
|  1 | aaa  |  10 | 1552577181 |       NULL | shard1   |
|  2 | ccc  |  30 | 1552577181 |       NULL | shard1   |
|  3 | eee  |  50 | 1552577181 |       NULL | shard1   |
|  4 | ggg  |  70 | 1552577181 |       NULL | shard1   |
|  5 | iii  |  90 | 1552577181 |       NULL | shard1   |
|  1 | bbb  |  20 | 1552577167 |       NULL | shard2   |
|  2 | ddd  |  40 | 1552577167 |       NULL | shard2   |
|  3 | fff  |  60 | 1552577167 |       NULL | shard2   |
|  4 | hhh  |  80 | 1552577167 |       NULL | shard2   |
|  5 | jjj  | 100 | 1552577167 |       NULL | shard2   |
|  6 | kkk  | 110 | 1552801266 |       NULL | shard2   |
+----+------+-----+------------+------------+----------+
shard1 > 5 rows in set (0.00 sec)
shard2 > 6 rows in set (0.00 sec)
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
