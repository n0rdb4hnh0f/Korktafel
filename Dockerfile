# --- ビルドステージ ---
FROM golang:1.22-alpine AS builder

# SQLiteのビルドに必要なツールをインストール
RUN apk add --no-cache gcc musl-dev

WORKDIR /app

# 依存関係をコピーしてダウンロード
COPY go.mod go.sum ./
RUN go mod download

# ソースコードをコピー
COPY . .

# CGOを有効にしてビルド（SQLiteを使用するため必要）
RUN CGO_ENABLED=1 GOOS=linux go build -o main .

# --- 実行ステージ ---
FROM alpine:latest

# タイムゾーン設定などのためのライブラリ
RUN apk add --no-cache ca-certificates tzdata

WORKDIR /app

# ビルドしたバイナリをコピー
COPY --from=builder /app/main .

# 実行
EXPOSE 8080
CMD ["./main"]
