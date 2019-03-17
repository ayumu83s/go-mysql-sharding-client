package mysql

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"github.com/BurntSushi/toml"
	"os"
)

type Config struct {
	Databases []DBConfig
	DatabaseCommon DBConfigCommon
}

type DBConfig struct {
	Database string
	Host string
	Port int
	User string
	Password string
}

type DBConfigCommon struct {
	User string
	Password string

}

type database struct {
	db *sql.DB
	config DBConfig
}

type Client struct {
	databases []database
}

func loadConfig(configPath string) (*Config, error) {
	_, err := os.Stat(configPath)
	if os.IsNotExist(err) {
		return nil, err
	}

	var config Config
	_, err = toml.DecodeFile(configPath, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}

func NewClient(configPath string) (*Client, error) {
	config, err := loadConfig(configPath)
	if err != nil {
		return nil, err
	}

	databases := make([]database, len(config.Databases))
	for i, dbConfig := range config.Databases {
		if dbConfig.User == "" {
			dbConfig.User = config.DatabaseCommon.User
		}
		if dbConfig.Password == "" {
			dbConfig.Password = config.DatabaseCommon.Password
		}

		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
			dbConfig.User,
			dbConfig.Password,
			dbConfig.Host,
			dbConfig.Port,
			dbConfig.Database,
		)
		db, err := sql.Open("mysql", dsn)
		if err != nil {
			panic(err.Error())
		}
		if err = db.Ping(); err != nil {
			panic(err.Error())
		}
		databases[i] = database{
			db: db,
			config: dbConfig,
		}
		fmt.Printf("Connection ok(%s)\n", dbConfig.Database)
	}
	client := &Client{
		databases: databases,
	}
	return client, nil
}

func (c *Client) Disconnect() {
	for i := range c.databases {
		c.databases[i].db.Close()
	}
}
