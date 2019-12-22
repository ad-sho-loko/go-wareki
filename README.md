# go-wareki

Goにて日本の和暦をサポートするためのライブラリです。改暦以降*の西暦と和暦の変換に対応しています。
標準ライブラリである`time.Time`と実装をなるべく統一するようにしています。

`go-wareki` is a library for wareki(Japanese-era).

\* 日本におけるグレゴリオ暦の採用以降（a.k.a 明治6年以降）

## Install

```bash
go get -u github.com/ad-sho-loko/wareki
```

## Quick Start

```
package main

import "github.com/ad-sho-loko/wareki"

func main(){
	short, _ := wareki.Parse(wareki.JISX0301Short, "R01.12.24")
	mid, _  := wareki.Parse(wareki.JISX0301Mid, "令01.12.24")
	long, _ := wareki.Parse(wareki.JISX0301Long, "令和01.12.24")

	wareki.Format(short, wareki.JISX0301Short) // R01.12.24
	wareki.Format(mid, wareki.JISX0301Mid)     // 令01.12.24
	wareki.Format(long, wareki.JISX0301Long)   // 令和01.12.24
}
```

