package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/urfave/cli/v2"

	"github.com/gen2brain/beeep"

	"mvdan.cc/sh/v3/syntax"
)

var (
	command       string = ""
	cronName      string = ""
	notifySuccess bool
)

func main() {

	app := &cli.App{
		Name:  "cronmon",
		Usage: "runs a job and notifies according to your wishes 🔔",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "command",
				Required:    true,
				Destination: &command,
				Aliases:     []string{"c"},
			},
			&cli.StringFlag{
				Name:        "cron-name",
				Required:    true,
				Destination: &cronName,
				Aliases:     []string{"cn"},
			},
			&cli.BoolFlag{
				Name:        "notify-success",
				DefaultText: "does not notify when the command was successfully executed",
				Destination: &notifySuccess,
				Aliases:     []string{"success"},
				Value:       false,
			},
		},
		Action: func(*cli.Context) error {
			r := strings.NewReader(command)
			f, err := syntax.NewParser().Parse(r, "")
			if err != nil {
				log.Fatal(err)
			}
			var buf bytes.Buffer
			err = syntax.NewPrinter().Print(&buf, f)
			if err != nil {
				log.Fatal(err)
			}

			ex := exec.Command("bash", "-c", buf.String())
			ex.Stdout = os.Stdout
			ex.Stderr = os.Stderr
			err = ex.Run()
			if err != nil {
				if exitError, ok := err.(*exec.ExitError); ok {
					exitCode := exitError.ExitCode()
					fmt.Println("Command exited with non-zero status:", exitCode)
					beeep.Alert("Cronmon", fmt.Sprintf("running CronJob '%s' failed", cronName), "")
				}
				return nil
			}
			if notifySuccess {
				beeep.Alert("Cronmon", fmt.Sprintf("running CronJob '%s' was successful", cronName), "")
			}
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
