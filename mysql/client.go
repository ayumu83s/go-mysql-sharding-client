package mysql

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"github.com/BurntSushi/toml"
	"os"
	"strings"
	"time"
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

const ViewShardHeader = "database"

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

func (c *Client) Executor(query	string) {

	if query == "exit" {
		fmt.Println("Bye!")
		c.Disconnect()
		os.Exit(0)
		return
	}

	var maxValueLength map[string]int
	var result []map[string]string
	var columns []string
	var scanArgs []interface{}
	var values []sql.RawBytes
	execTimes := make([]float64, len(c.databases))
	execRows := make([]int, len(c.databases))

	for i, database := range c.databases {
		begin := time.Now()
		rows, err := database.db.Query(query)
		end := time.Now()
		if err != nil {
			fmt.Printf("%s\n", err.Error())
			return
		}
		execTimes[i] = end.Sub(begin).Seconds()

		if columns == nil {
			columns, err = rows.Columns()
			if err != nil {
				panic(err.Error())
			}
			values = make([]sql.RawBytes, len(columns))
			if maxValueLength == nil {
				maxValueLength = make(map[string]int, len(columns)+1)
				maxValueLength[ ViewShardHeader ] = len(ViewShardHeader)
			}

			scanArgs = make([]interface{}, len(values))
			for i := range values {
				scanArgs[i] = &values[i]
				maxValueLength[ columns[i] ] = len(columns[i])
			}
		}

		for rows.Next() {
			err = rows.Scan(scanArgs...)
			if err != nil {
				panic(err.Error())
			}

			var value string
			data := make(map[string]string)
			for i, col := range values {
				if col == nil {
					value = "NULL"
				} else {
					value = string(col)
				}
				if maxValueLength[ columns[i] ] < len(value) {
					maxValueLength[ columns[i] ] = len(value)
				}
				data[columns[i]] = value
			}
			data[ViewShardHeader] = database.config.Database
			if maxValueLength[ ViewShardHeader ] < len(database.config.Database) {
				maxValueLength[ ViewShardHeader ] = len(database.config.Database)
			}
			result = append(result, data)
			execRows[i]++
		}
		rows.Close()
	}
	columns = append(columns, ViewShardHeader)
	viewHeader(maxValueLength, columns)
	viewBody(maxValueLength, columns, result)
	for i, execTime := range execTimes {
		fmt.Printf("%s > %d rows in set (%.2f sec)\n", c.databases[i].config.Database, execRows[i], execTime)
	}
}

func viewHeader(maxValueLength map[string]int, columns []string) {
	headStr := "|"
	for _, columnName := range columns {
		columnNameLen := len(columnName)
		margin := maxValueLength[columnName] - columnNameLen
		headStr += " "
		headStr += strings.Repeat(" ", margin)
		headStr += columnName
		headStr += " |"
	}
	viewBorder(maxValueLength, columns)
	fmt.Printf("%s\n", headStr)
	viewBorder(maxValueLength, columns)
}

func viewBorder(maxValueLength map[string]int, columns []string) {
	border := "+"
	for _, columnName := range columns {
		columnNameLen := len(columnName)
		border += strings.Repeat("-", (columnNameLen + 2)) + "+"
	}
	fmt.Printf("%s\n", border)
}

func viewBody(maxValueLength map[string]int, columns []string, result []map[string]string) {
	for _, row := range result {
		rowStr := "|"
		for _, column := range columns {
			value := row[column]
			valueLen := len(value)
			margin := maxValueLength[column] - valueLen

			rowStr += " "
			rowStr += strings.Repeat(" ", margin)
			rowStr += value
			rowStr += " |"
		}
		fmt.Printf("%s\n", rowStr)
	}
	viewBorder(maxValueLength, columns)
}

func (c *Client) Disconnect() {
	for i := range c.databases {
		c.databases[i].db.Close()
	}
}
