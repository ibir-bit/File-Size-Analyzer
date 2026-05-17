package main

import (
	"fmt"
	"os"

	"code"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "hexlet-path-size",
		Usage: "print size of a file or directory; supports -r (recursive), -H (human-readable), -a (include hidden)",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "recursive",
				Aliases: []string{"r"},
				Usage:   "recursive size of directories",
			},
			&cli.BoolFlag{
				Name:    "human",
				Aliases: []string{"H"},
				Usage:   "human-readable sizes (auto-select unit)",
			},
			&cli.BoolFlag{
				Name:    "all",
				Aliases: []string{"a"},
				Usage:   "include hidden files and directories",
			},
		},
		Action: func(c *cli.Context) error {
			if c.NArg() == 0 {
				return fmt.Errorf("не указан путь к файлу или директории")
			}

			path := c.Args().First()
			recursive := c.Bool("recursive")
			human := c.Bool("human")
			all := c.Bool("all")

			sizeStr, err := code.GetPathSize(path, recursive, human, all)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Ошибка: %v\n", err)
				os.Exit(1)
			}

			fmt.Printf("%s\t%s\n", sizeStr, path)

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка запуска приложения: %v\n", err)
		os.Exit(1)
	}
}
