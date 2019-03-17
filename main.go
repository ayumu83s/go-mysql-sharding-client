package main

import (
	"fmt"
	"github.com/ayumu83s/go-mysql-sharding-client/mysql"
	"github.com/c-bata/go-prompt"
	"gopkg.in/urfave/cli.v1"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "mysql-shard"
	app.Usage = "database management prompt for sharded databases"
	app.Version = "1.0.0"
	app.Flags = []cli.Flag {
		cli.StringFlag{
			Name: "config",
			Usage: "config path(toml format)",
		},
	}
	app.Action = func(c *cli.Context) error {
		configPath := c.String("config")
		client, err := mysql.NewClient(configPath)
		if err != nil {
			return err
		}
		fmt.Println("Please use `exit` or `Ctrl-D` to exit this program.")
		defer func() {
			client.Disconnect()
			fmt.Println("Bye!")
		}()

		p := prompt.New(
			client.Executor,
			mysql.Completer,
			prompt.OptionTitle("sql-prompt"),
			prompt.OptionPrefix("sharding > "),
			prompt.OptionInputTextColor(prompt.Yellow),
		)
		p.Run()
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		panic(err.Error())
	}
}