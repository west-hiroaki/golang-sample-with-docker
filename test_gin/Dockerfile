FROM golang:1.13-alpine3.10

# install OS dependencies and golang's dep
RUN apk add --no-cache bash git make curl gcc g++ musl-dev \
    && curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

# install needs golang package
RUN go get -u \
      github.com/gin-gonic/gin \
      github.com/jinzhu/gorm \
      github.com/mattn/go-sqlite3 \
      golang.org/x/lint/golint

CMD ["go", "run", "src/main.go"]
