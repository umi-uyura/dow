package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/goodsign/monday"
)

func main() {
	var (
		withDate = flag.Bool("w", false, "Display with date")
		separator = flag.String("s", "-", "Specify date separator")
		repNonDigit = regexp.MustCompile(`[^\d]`)
		repOtherSep = regexp.MustCompile(`[/\.]`)
	)

	flag.Parse()

	if len(*separator) > 1 {
		fmt.Println("Separator is one character.")
		os.Exit(1)
	}

	date := time.Now()

	if flag.NArg() > 0 {
		datestring := flag.Args()[0]

		sep := repNonDigit.FindString(datestring)
		if len(sep) > 0 {
			*separator = sep
		}

		datestring = repOtherSep.ReplaceAllString(datestring, "-")

		d, err := time.Parse("2006-1-2", datestring)
		if err == nil {
			date = d
		} else {
			fmt.Println("Invalid date format.")
			fmt.Println("Valid date formats: 'yyyy-mm-dd', 'yyyy/mm/dd', 'yyyy.mm.dd'.")
			os.Exit(1)
		}
	}

	var locale monday.Locale = monday.LocaleEnUS
	lang := os.Getenv("LANG")

	if len(lang) > 0 {
		lang1 := strings.Split(lang, ".")
		for _, l := range monday.ListLocales() {
			if lang1[0] == string(l) {
				locale = l
				break
			}
		}
	}

	if !*withDate {
		fmt.Println(monday.Format(date, "Mon", locale))
	} else {
		dateFormat := fmt.Sprintf("2006%s1%s2(Mon)", *separator, *separator)
		fmt.Println(monday.Format(date, dateFormat, locale))
	}
}
