# Description

DB を使用したシンプルな TODO 管理アプリ。

# Used package

* [Gin Web Framework](https://github.com/gin-gonic/gin)
  * 軽量な Web フレームワーク
* [GORM](https://github.com/jinzhu/gorm)
  * ORM library
* [go-sqlite3](https://github.com/mattn/go-sqlite3)
  * sqlite ドライバ
* [Golint](https://github.com/golang/lint)
  * golang のコードチェッカー

# Procedure

## コンテナ起動

```
$ make build
$ make up
```

8080 ポートでWebサーバが起動するので、ブラウザでアクセス可能になる

## ブラウザでアクセス

http://localhost:8080/todo/

## コンテナに入る場合

コンテナを起動した状態で

```
$ make ssh
```

# Tips

## VSCode を使っている場合の環境設定手順

環境変数 GOPATH を設定しておかないと、エディタ上でサブディレクトリのパッケージが反応してくれない

`settings.json` の `go.gopath` にこのファイルを置いたローカルパスを指定する必要あり
```
{
    "go.gopath": "/hogehoge/golang-sample-with-docker/test_gin",
    :
    :
```
