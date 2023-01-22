#!/usr/bin/env bash

echo "dow"
go run *.go
echo "* Display today day of week"
echo

echo "dow 2023-1-1"
go run *.go 2023-1-1
echo "Sun"
echo

echo "dow 2023-01-01"
go run *.go 2023-01-01
echo "Sun"
echo

echo "dow 2023/1/1"
go run *.go 2023/1/1
echo "Sun"
echo

echo "dow 2023.1.1"
go run *.go 2023.1.1
echo "Sun"
echo

echo "dow -w 2023-1-1"
go run *.go -w 2023-1-1
echo "2023-1-1(Sun)"
echo

echo "dow -w 2023-01-01"
go run *.go -w 2023-01-01
echo "2023-1-1(Sun)"
echo

echo "dow -w -s / 2023-1-1"
go run *.go -w -s / 2023-1-1
echo "2023/1/1(Sun)"
echo

echo "dow -w -s . 2023-1-1"
go run *.go -w -s . 2023-1-1
echo "2023.1.1(Sun)"
echo

echo "dow -w -s '#' 2023-1-1"
go run *.go -w -s '#' 2023-1-1
echo "2023#1#1(Sun)"
echo

echo "dow 20230101"
go run *.go 20230101 2> /dev/null
echo "* Error"
echo

echo "dow 2023#1#1"
go run *.go 2023#1#1 2> /dev/null
echo "* Error"
echo

echo "dow -w -s -- 2023-1-1"
go run *.go -w -s -- 2023-1-1 2> /dev/null
echo "* Error"
echo

echo "LANG=ja_JP dow"
LANG=ja_JP go run *.go
echo "* Display today day of week and Japanese"
echo

echo "LANG=ja_JP dow 2023-1"
LANG=ja_JP go run *.go 2023-1-1
echo "日"
echo

echo "LANG=ja_JP dow -w 2023-1-1"
LANG=ja_JP go run *.go -w 2023-1-1
echo "2023-1-1(日)"
echo
