FROM --platform=linux/x86_64 golang:1.20

WORKDIR /usr/local/go/src/go_crm/

# airをインストールしてホットリロードできるようにする
RUN go install github.com/cosmtrek/air@v1.42.0

COPY application/. /usr/local/go/src/go_crm/

RUN go mod download
