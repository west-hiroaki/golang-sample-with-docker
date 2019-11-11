.PHONY: build
build: ## go 環境のコンテナを生成します。
	docker-compose build

.PHONY: build.local
build.local: ## ローカル環境を構築します。
	go get "github.com/gin-gonic/gin"
	go get "github.com/jinzhu/gorm"
	go get "github.com/mattn/go-sqlite3"
	go get "golang.org/x/lint/golint"

.PHONY: up
up: ## gin サーバを起動します。
	docker-compose up

.PHONY: data.clean.local
data.clean.local: ## ローカルで 作成されたsplite3 データを削除します。
	rm test.sqlite3

.PHONY: clean.local
clean.local: data.clean.local  ## ローカルでコンパイルされた go 実行ファイルを削除します。
	rm main

.PHONY: format
format: ## アプリのコードに対してコード整形を行います。
	gofmt -w src/main.go
	gofmt -w src/app/*

.PHONY: lint
lint: build.local  ## ローカルで golint を実行します。
	golint src/main.go
	find ./src/app -name '*.go' | xargs -n 1 golint

.PHONY: ssh
ssh: ## 起動しているコンテナに ssh ログインします。
	docker exec -it app_test_gin bash

.PHONY: help
.DEFAULT_GOAL := help
help: ## ヘルプを表示させます。
	@grep -E '^[a-zA-Z_\.-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
