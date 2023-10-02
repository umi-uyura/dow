dow (Day Of Week)
=================

Display day of week for specified date.


Install
-------

CLI Tool

```shell
go install github.com/umi-uyura/dow/cmd/dow@latest
```

Library

```shell
go get github.com/umi-uyura/dow
```


Usage
-----

### CLI Tool

```shell
dow [options] [date]
```

* Support date format
  - `yyyy-mm-dd`
  - `yyyy/mm/dd`
  - `yyyy.mm.dd`
* When execute without arguments, display current day of week.
* Display day of week in language of locale, set `LANG` environment variable.

Options

* `-w` Outputs the name of the day of the week along with the date
* `-s <separator>` Specifies the date separator

#### Example

```shell
$ dow 2023-1-1
Sun

$ dow -w 2023-1-2
2023-1-2(Mon)

$ dow -w -s / 2023-1-3
2023/1/3(Tue)

$ LANG=ja_JP dow 2023/1/4
水

$ LANG=zh_CN dow 2023.1.5
四
```

### Library

See: [dow package - github.com/umi-uyura/dow - Go Packages](https://pkg.go.dev/github.com/umi-uyura/dow#section-documentation)


License
-------

MIT License
