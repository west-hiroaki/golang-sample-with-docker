# go ファイル達

## hello/main.go

疎通テスト用。
実行すると、Hello, World を print します。

## exercise_web_crawler/main.go

A Tour of Go の問題に対する自分なりの回答。
https://go-tour-jp.appspot.com/concurrency/10

# 実行手順

## コンテナ起動

```
$ docker-compose build
$ docker-compose up
```

## コンテナに入る

```
$ docker exec -it app_go bash
```

## go ファイル実行

```
root@1dc8dc8b2169:/go/src/app# go run hello/main.go
Hello, World
```
