package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/umi-uyura/dow"
)

const (
	EXITCODE_OK = iota
	EXITCODE_NG
)

type CLI struct {
	Stdout io.Writer
	Stderr io.Writer
	Stdin  io.Reader
}

func (cli *CLI) Run(args []string) int {
	var (
		regexOtherSep = regexp.MustCompile(`[/\.]`)
		withDate bool
		separator string
	)

	command := flag.NewFlagSet("dow", flag.ExitOnError)
	command.SetOutput(cli.Stdout)
	command.BoolVar(&withDate, "w", false, "Display with date")
	command.StringVar(&separator, "s", "-", "Specify date separator")

	err := command.Parse(args[1:])
	if err != nil {
		fmt.Fprintf(cli.Stdout, "ERROR: %v", err)
		return EXITCODE_NG
	}

	if len(separator) > 1 || len(separator) == 0 {
		fmt.Fprintln(cli.Stdout, "Separator is one character.")
		return EXITCODE_NG
	}

	date := time.Now()

	if command.NArg() > 0 {
		dateparam := command.Args()[0]
		datestring := regexOtherSep.ReplaceAllString(dateparam, "-")

		d, err := time.Parse("2006-1-2", datestring)
		if err == nil {
			date = d
		} else {
			fmt.Fprintln(cli.Stdout, "Invalid date format.")
			fmt.Fprintln(cli.Stdout, "Valid date formats: 'yyyy-mm-dd', 'yyyy/mm/dd', 'yyyy.mm.dd'.")
			return EXITCODE_NG
		}
	}

	lang := os.Getenv("LANG")

	if !withDate {
		fmt.Fprint(cli.Stdout, dow.GetLocaleName(date, lang))
	} else {
		dayofweek := dow.GetLocaleNameWithDate(date, lang)
		if separator != "-" {
			dayofweek = strings.Replace(dayofweek, "-", separator, -1)
		}
		fmt.Fprint(cli.Stdout, dayofweek)
	}

	return EXITCODE_OK
}
