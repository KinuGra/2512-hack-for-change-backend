# 環境構築

```
brew install go
git clone <repo>
go mod tidy
go run cmd/app/main.go
```

ローカルの場合は`curl http://localhost:8080/ping
`で`pong`が返ってくることを確認

# プロジェクト作成

```
$ go mod init github.com/KinuGra/2512-hack-for-change-backend
$ go get -u github.com/gin-gonic/gin
```
