dow (Day Of Week)
=================

Display day of week for specified date.


Install
-------

```
go install github.com/umi-uyura/dow@latest
```


Usage
-----

```
dow [date]
```

Support date format

- `yyyy-mm-dd`
- `yyyy/mm/dd`
- `yyyy.mm.dd`

When execute without arguments, display current day of week.

Display day of week in language of locale, set `LANG` environment variable.


### Example

```shell
$ dow 2023-1-1
Sun

$ LANG=ja_JP dow 2023/1/2
月

$ LANG=zh_CN dow 2023.1.3
二
```
