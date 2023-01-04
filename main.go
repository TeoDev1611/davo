package main

import (
	"os"
	"time"

	"github.com/TeoDev1611/davo/core"
	"github.com/i582/cfmt/cmd/cfmt"
	"github.com/urfave/cli/v2"
)

func init() {
	cli.VersionPrinter = func(cCtx *cli.Context) {
		cfmt.Fprintf(cCtx.App.Writer, "Davo 🥬! Version is {{%s}}::green\n", cCtx.App.Version)
	}
}

func main() {
	app := &cli.App{
		Name:     "davo",
		Version:  "v0.1.0",
		Compiled: time.Now(),
		Authors: []*cli.Author{
			{
				Name:  "Teo",
				Email: "teodev1611@gmail.com",
			},
		},
		Usage:     "Davo 🥬! A easy way for download binaries to the system",
		Copyright: "Copyright (c) 2023 Teo. All Rights Reserved.",
		Action: func(ctx *cli.Context) error {
			core.DavoSetup()
			return nil
		},
		Commands: []*cli.Command{
			{
				Name:    "download",
				Aliases: []string{"install"},
				Usage:   "Download the binary from GitHub with Davo🥬!",
				Action: func(ctx *cli.Context) error {
					core.DownloadNow(ctx.Args().First())
					return nil
				},
			},
		},
	}
	app.Suggest = true

	if err := app.Run(os.Args); err != nil {
		cfmt.Printf("Davo 🥬! {{ERROR}}::red|bold\n{{%s}}::cyan\n", err.Error())
		os.Exit(2)
	}
}
