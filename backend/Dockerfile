FROM golang:1.25-alpine

RUN apk update && apk add git

WORKDIR /app

# ホットリロードツール "Air" のインストール
RUN go install github.com/air-verse/air@v1.61.0

COPY go.mod go.sum ./
RUN go mod download

COPY . .

CMD ["air", "-c", ".air.toml"]