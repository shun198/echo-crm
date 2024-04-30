FROM --platform=linux/x86_64 golang:1.20

WORKDIR /usr/local/go/src/my_project/

# airをインストールしてホットリロードできるようにする
RUN go install github.com/cosmtrek/air@v1.42.0

COPY application/. /usr/local/go/src/my_project/

RUN go mod download
