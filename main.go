package main

import (
	"fmt"
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
		fmt.Printf("%s\n", configPath)
		fmt.Println("Please use `Ctrl-D` to exit this program.")
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		panic(err.Error())
	}
}