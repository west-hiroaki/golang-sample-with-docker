# Description

ホストの go ファイルを go run で実行するサンプル。
[A Tour of Go - Exercise: Web Crawler](https://go-tour-jp.appspot.com/concurrency/10) に対する自分なりの回答。

# Procedure

## コンテナ生成

```
$ docker-compose build
```

## go ファイル実行

```
$ docker-compose up
Recreating app_simple_go_run ... done
Attaching to app_simple_go_run
app_simple_go_run | found: https://golang.org/ "The Go Programming Language"
app_simple_go_run | not found: https://golang.org/cmd/
app_simple_go_run | found: https://golang.org/pkg/ "Packages"
app_simple_go_run | not found: https://golang.org/cmd/
app_simple_go_run | found: https://golang.org/pkg/fmt/ "Package fmt"
app_simple_go_run | found: https://golang.org/pkg/os/ "Package os"
```
