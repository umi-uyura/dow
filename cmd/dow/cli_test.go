package main

import (
	"bytes"
	"os"
	"strings"
	"testing"
	"time"
)

var TEST_PG = os.Args[0]

func getTodayWeekDayName() string {
	weekday := time.Now().Weekday()
	weekdayString := weekday.String()

	return weekdayString[:3]
}

func TestDow(t *testing.T) {
	todayWeekDayName := getTodayWeekDayName()
	
	tests := []struct {
		args []string
		expected string
	}{
		{[]string{TEST_PG}, todayWeekDayName},
		{[]string{TEST_PG, "2023-1-1"}, "Sun"},
		{[]string{TEST_PG, "2023-1-2"}, "Mon"},
		{[]string{TEST_PG, "2023-1-3"}, "Tue"},
		{[]string{TEST_PG, "2023-1-4"}, "Wed"},
		{[]string{TEST_PG, "2023-1-5"}, "Thu"},
		{[]string{TEST_PG, "2023-1-6"}, "Fri"},
		{[]string{TEST_PG, "2023-1-7"}, "Sat"},
		{[]string{TEST_PG, "2023-12-31"}, "Sun"},
	}

	for _, tt := range tests {
		var (
			out bytes.Buffer
			err bytes.Buffer
			in bytes.Buffer
		)
		
		cli := &CLI{
			Stdout: &out,
			Stderr: &err,
			Stdin:  &in,
		}

		var runName string
		if len(tt.args[1:]) == 0 {
			runName = "today"
		} else {
			runName = strings.Join(tt.args[1:], "/")
		}

		t.Run(runName, func(t *testing.T) {
			exitcode := cli.Run(tt.args)
			if exitcode != EXITCODE_OK {
				t.Errorf("Expected exitcode %d, but got %d", EXITCODE_OK, exitcode)
			}

			actual := out.String()
			if actual != tt.expected {
				t.Errorf("Expected %q, but got %q", tt.expected, actual)
			}
		})
	}
}

func TestDowZeroPaddingDate(t *testing.T) {
	todayWeekDayName := getTodayWeekDayName()
	
	tests := []struct {
		args []string
		expected string
	}{
		{[]string{TEST_PG}, todayWeekDayName},
		{[]string{TEST_PG, "2023-01-01"}, "Sun"},
		{[]string{TEST_PG, "2023-01-02"}, "Mon"},
		{[]string{TEST_PG, "2023-01-03"}, "Tue"},
		{[]string{TEST_PG, "2023-01-04"}, "Wed"},
		{[]string{TEST_PG, "2023-01-05"}, "Thu"},
		{[]string{TEST_PG, "2023-01-06"}, "Fri"},
		{[]string{TEST_PG, "2023-01-07"}, "Sat"},
		{[]string{TEST_PG, "2023-12-31"}, "Sun"},
	}

	for _, tt := range tests {
		var (
			out bytes.Buffer
			err bytes.Buffer
			in bytes.Buffer
		)
		
		cli := &CLI{
			Stdout: &out,
			Stderr: &err,
			Stdin:  &in,
		}

		var runName string
		if len(tt.args[1:]) == 0 {
			runName = "today"
		} else {
			runName = strings.Join(tt.args[1:], "/")
		}

		t.Run(runName, func(t *testing.T) {
			exitcode := cli.Run(tt.args)
			if exitcode != EXITCODE_OK {
				t.Errorf("Expected exitcode %d, but got %d", EXITCODE_OK, exitcode)
			}

			actual := out.String()
			if actual != tt.expected {
				t.Errorf("Expected %q, but got %q", tt.expected, actual)
			}
		})
	}
}

func TestDowOtherSeparator(t *testing.T) {
	tests := []struct {
		args []string
		expected string
	}{
		{[]string{TEST_PG, "2023.1.1"}, "Sun"},
		{[]string{TEST_PG, "2023/1/2"}, "Mon"},
		{[]string{TEST_PG, "2023-1.3"}, "Tue"},
		{[]string{TEST_PG, "2023.1-4"}, "Wed"},
		{[]string{TEST_PG, "2023-1/5"}, "Thu"},
		{[]string{TEST_PG, "2023/1-6"}, "Fri"},
		{[]string{TEST_PG, "2023.1/7"}, "Sat"},
		{[]string{TEST_PG, "2023/12.31"}, "Sun"},
	}

	for _, tt := range tests {
		var (
			out bytes.Buffer
			err bytes.Buffer
			in bytes.Buffer
		)
		
		cli := &CLI{
			Stdout: &out,
			Stderr: &err,
			Stdin:  &in,
		}

		runName := strings.Join(tt.args[1:], "/")

		t.Run(runName, func(t *testing.T) {
			exitcode := cli.Run(tt.args)
			if exitcode != EXITCODE_OK {
				t.Errorf("Expected exitcode %d, but got %d", EXITCODE_OK, exitcode)
			}

			actual := out.String()
			if actual != tt.expected {
				t.Errorf("Expected %q, but got %q", tt.expected, actual)
			}
		})
	}
}

func TestDowInvalidDateFormat(t *testing.T) {
	const ERROR_MESSAGE = "Invalid date format.\nValid date formats: 'yyyy-mm-dd', 'yyyy/mm/dd', 'yyyy.mm.dd'.\n"

	tests := []struct {
		args []string
		expected string
	}{
		{[]string{TEST_PG, "20230101"}, ERROR_MESSAGE},
		{[]string{TEST_PG, "2023#1#2"}, ERROR_MESSAGE},
	}

	for _, tt := range tests {
		var (
			out bytes.Buffer
			err bytes.Buffer
			in bytes.Buffer
		)
		
		cli := &CLI{
			Stdout: &out,
			Stderr: &err,
			Stdin:  &in,
		}

		runName := strings.Join(tt.args[1:], "/")

		t.Run(runName, func(t *testing.T) {
			exitcode := cli.Run(tt.args)
			if exitcode != EXITCODE_NG {
				t.Errorf("Expected exitcode %q, but got %q", EXITCODE_NG, exitcode)
			}

			actual := out.String()
			if actual != tt.expected {
				t.Errorf("Expected %q, but got %q", tt.expected, actual)
			}
		})
	}
}

func TestDowWithDateOption(t *testing.T) {
	tests := []struct {
		args []string
		expected string
	}{
		{[]string{TEST_PG, "-w", "2023-1-1"}, "2023-1-1(Sun)"},
		{[]string{TEST_PG, "-w", "2023-1-2"}, "2023-1-2(Mon)"},
		{[]string{TEST_PG, "-w", "2023-1-3"}, "2023-1-3(Tue)"},
		{[]string{TEST_PG, "-w", "2023-1-4"}, "2023-1-4(Wed)"},
		{[]string{TEST_PG, "-w", "2023-1-5"}, "2023-1-5(Thu)"},
		{[]string{TEST_PG, "-w", "2023-1-6"}, "2023-1-6(Fri)"},
		{[]string{TEST_PG, "-w", "2023-1-7"}, "2023-1-7(Sat)"},
		{[]string{TEST_PG, "-w", "2023-12-31"}, "2023-12-31(Sun)"},
	}

	for _, tt := range tests {
		var (
			out bytes.Buffer
			err bytes.Buffer
			in bytes.Buffer
		)
		
		cli := &CLI{
			Stdout: &out,
			Stderr: &err,
			Stdin:  &in,
		}

		runName := strings.Join(tt.args[1:], "/")

		t.Run(runName, func(t *testing.T) {
			exitcode := cli.Run(tt.args)
			if exitcode != EXITCODE_OK {
				t.Errorf("Expected exitcode %d, but got %d", EXITCODE_OK, exitcode)
			}

			actual := out.String()
			if actual != tt.expected {
				t.Errorf("Expected %q, but got %q", tt.expected, actual)
			}
		})
	}
}

func TestDowWithDateAndSeparaterOption(t *testing.T) {
	tests := []struct {
		args []string
		expected string
		exitcode int
	}{
		{[]string{TEST_PG, "-w", "-s", "-", "2023-1-1"}, "2023-1-1(Sun)", EXITCODE_OK},
		{[]string{TEST_PG, "-w", "-s", "/", "2023-1-2"}, "2023/1/2(Mon)", EXITCODE_OK},
		{[]string{TEST_PG, "-w", "-s", ".", "2023-1-3"}, "2023.1.3(Tue)", EXITCODE_OK},
		{[]string{TEST_PG, "-w", "-s", "#", "2023-1-4"}, "2023#1#4(Wed)", EXITCODE_OK},
		{[]string{TEST_PG, "-w", "-s", "", "2023-1-5"}, "Separator is one character.\n", EXITCODE_NG},
		{[]string{TEST_PG, "-w", "-s", "//", "2023-1-6"}, "Separator is one character.\n", EXITCODE_NG},
		{[]string{TEST_PG, "-s", ".", "2023-1-7"}, "Sat", EXITCODE_OK},
	}

	for _, tt := range tests {
		var (
			out bytes.Buffer
			err bytes.Buffer
			in bytes.Buffer
		)
		
		cli := &CLI{
			Stdout: &out,
			Stderr: &err,
			Stdin:  &in,
		}

		runName := strings.Join(tt.args[1:], "/")

		t.Run(runName, func(t *testing.T) {
			exitcode := cli.Run(tt.args)
			if exitcode != tt.exitcode {
				t.Errorf("Expected exitcode %d, but got %d", tt.exitcode, exitcode)
			}

			actual := out.String()
			if actual != tt.expected {
				t.Errorf("Expected %q, but got %q", tt.expected, actual)
			}
		})
	}
}

func TestDowLangJaJp(t *testing.T) {
	tests := []struct {
		args []string
		expected string
		exitcode int
	}{
		{[]string{TEST_PG, "2023-1-1"}, "日", EXITCODE_OK},
		{[]string{TEST_PG, "-w", "2023-1-2"}, "2023-1-2(月)", EXITCODE_OK},
		{[]string{TEST_PG, "-w", "-s", ".", "2023-1-3"}, "2023.1.3(火)", EXITCODE_OK},
	}

	for _, tt := range tests {
		var (
			out bytes.Buffer
			err bytes.Buffer
			in bytes.Buffer
		)
		
		cli := &CLI{
			Stdout: &out,
			Stderr: &err,
			Stdin:  &in,
		}

		runName := strings.Join(tt.args[1:], "/")

		t.Run(runName, func(t *testing.T) {
			t.Setenv("LANG", "ja_JP")
			
			exitcode := cli.Run(tt.args)
			if exitcode != tt.exitcode {
				t.Errorf("Expected exitcode %d, but got %d", tt.exitcode, exitcode)
			}

			actual := out.String()
			if actual != tt.expected {
				t.Errorf("Expected %q, but got %q", tt.expected, actual)
			}
		})
	}
}
