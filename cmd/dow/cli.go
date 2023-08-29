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

type CLI struct {
	Stdout io.Writer
	Stderr io.Writer
	Stdin  io.Reader
}

func (cli *CLI) Run() {
	var (
		withDate = flag.Bool("w", false, "Display with date")
		separator = flag.String("s", "-", "Specify date separator")
		regexOtherSep = regexp.MustCompile(`[/\.]`)
	)

	flag.Parse()

	if len(*separator) > 1 {
		fmt.Fprintln(cli.Stdout, "Separator is one character.")
		os.Exit(1)
	}

	date := time.Now()

	if flag.NArg() > 0 {
		dateparam := flag.Args()[0]
		datestring := regexOtherSep.ReplaceAllString(dateparam, "-")

		d, err := time.Parse("2006-1-2", datestring)
		if err == nil {
			date = d
		} else {
			fmt.Fprintln(cli.Stdout, "Invalid date format.")
			fmt.Fprintln(cli.Stdout, "Valid date formats: 'yyyy-mm-dd', 'yyyy/mm/dd', 'yyyy.mm.dd'.")
			os.Exit(1)
		}
	}

	lang := os.Getenv("LANG")

	if !*withDate {
		fmt.Fprintln(cli.Stdout, dow.GetLocaleName(date, lang))
	} else {
		dayofweek := dow.GetLocaleNameWithDate(date, lang)
		if *separator != "-" {
			dayofweek = strings.Replace(dayofweek, "-", *separator, -1)
		}
		fmt.Fprintln(cli.Stdout, dayofweek)
	}
}
