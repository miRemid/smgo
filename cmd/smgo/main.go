package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sort"
	"time"

	"github.com/miRemid/smgo"

	"github.com/urfave/cli/v2"
)

var client *smgo.SmClient

func init() {
	client = smgo.NewSmClient()
	client.SetTimeout(time.Second * time.Duration(5))
}

func printJSON(model interface{}) error {
	data, err := json.MarshalIndent(model, "", "  ")
	if err != nil {
		return err
	}
	fmt.Println(string(data))
	return nil
}

func main() {
	app := &cli.App{
		Name:  "smgo",
		Usage: "upload to or delete from sm.ms",
		Commands: []*cli.Command{
			{
				Name:    "upload",
				Aliases: []string{"u"},
				Usage:   "upload image to sm.ms",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "token",
						Aliases: []string{"t"},
						Usage:   "set account token",
					},
					&cli.IntFlag{
						Name:    "time",
						Aliases: []string{"tm"},
						Usage:   "set timeout",
						Value:   5,
					},
				},
				Action: func(c *cli.Context) error {
					client.SetToken(c.String("token"))
					if t := c.Int("time"); t != 0 {
						client.SetTimeout(time.Second * time.Duration(t))
					}
					files := c.Args().Slice()
					imgs, err := client.Uploads(files...)
					if err != nil {
						return err
					}
					for _, img := range imgs {
						if err := printJSON(img); err != nil {
							return err
						}
					}
					return nil
				},
			},
			{
				Name:    "delete",
				Aliases: []string{"d"},
				Usage:   "delete image from sm.ms",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "token",
						Aliases: []string{"t"},
						Usage:   "set account token",
					},
					&cli.IntFlag{
						Name:    "time",
						Aliases: []string{"tm"},
						Usage:   "set timeout",
						Value:   5,
					},
				},
				Action: func(c *cli.Context) error {
					client.SetToken(c.String("token"))
					if t := c.Int("time"); t != 0 {
						client.SetTimeout(time.Second * time.Duration(t))
					}
					res, err := client.Delete(c.Args().First())
					if err != nil {
						return err
					}
					return printJSON(res)
				},
			},
			{
				Name:    "profile",
				Aliases: []string{"p"},
				Usage:   "print profile infomation",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "token",
						Aliases:  []string{"t"},
						Usage:    "set account token",
						Required: true,
					},
					&cli.IntFlag{
						Name:    "time",
						Aliases: []string{"tm"},
						Usage:   "set timeout",
						Value:   5,
					},
				},
				Action: func(c *cli.Context) error {
					client.SetToken(c.String("token"))
					if t := c.Int("time"); t != 0 {
						client.SetTimeout(time.Second * time.Duration(t))
					}
					res, err := client.Profile()
					if err != nil {
						return err
					}
					return printJSON(res)
				},
			},
			{
				Name:    "history",
				Aliases: []string{"h"},
				Usage:   "print upload history",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "token",
						Aliases: []string{"t"},
						Usage:   "set account token",
					},
					&cli.IntFlag{
						Name:    "time",
						Aliases: []string{"tm"},
						Usage:   "set timeout",
						Value:   5,
					},
				},
				Action: func(c *cli.Context) error {
					client.SetToken(c.String("token"))
					if t := c.Int("time"); t != 0 {
						client.SetTimeout(time.Second * time.Duration(t))
					}
					imgs, err := client.History()
					if err != nil {
						return err
					}
					return printJSON(imgs)
				},
			},
			{
				Name:    "clear",
				Aliases: []string{"c"},
				Usage:   "clear upload history",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "token",
						Aliases: []string{"t"},
						Usage:   "set account token",
					},
					&cli.IntFlag{
						Name:    "time",
						Aliases: []string{"tm"},
						Usage:   "set timeout",
						Value:   5,
					},
				},
				Action: func(c *cli.Context) error {
					client.SetToken(c.String("token"))
					if t := c.Int("time"); t != 0 {
						client.SetTimeout(time.Second * time.Duration(t))
					}
					imgs, err := client.Clear()
					if err != nil {
						return err
					}
					return printJSON(imgs)
				},
			},
		},
	}
	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
