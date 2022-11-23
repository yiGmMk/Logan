# compile
```
cd ServerGo
export GOPATH=`pwd`
go mod vendor
CGO_ENABLED=0 GOARCH=amd64 go build  -o server logan/server/cmd
```

# run

```
mkdir logfile
./server -c src/logan/server/cmd/logan.toml
```


repo<https://github.com/longbai/Logan/tree/add_go_server>


curl -X POST "http://localhost:8887/logan/web/upload.json" -F '{
  "client": "Web",
  "webSource": "browser",
  "deviceId": "test-logan-web",
  "environment": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/107.0.0.0 Safari/537.36 Edg/107.0.1418.42",
  "customInfo": "{\"userId\":123456,\"biz\":\"Live Better\"}",
  "logPageNo": 1,
  "fileDate": "2022-11-23",
  "logArray": "%7B%22l%22%3A%22hjAls6nLBlEqJwz4EqclX4zXP4oMz8uVGs%2F4WbPKzt7dHmrAnJ%2FsIacssRmwjEB0%22%2C%22iv%22%3A%225viydf5lxfjj1bhv%22%2C%22k%22%3A%22LYX8iLlULYvX2%2Buc%2Fji4p9TG%2BTWKjOtU2YQul%2BAbOZW9e4F8qh9dY%2B3l8Rtrr1srthobR0r3fKKa%2F5ZjR1lsa3cdOyhHIeiCWMYQTr0w7PQrwdSUTO%2FccP1Vo8x5AJ6i%2Bf8inJ64DlyUkyAoxKJP1NiK%2B8PmKzqnuugktJwbeR8%3D%22%2C%22v%22%3A1%7D"
}'


curl -X POST "http://localhost:8887/logan/web/upload.json" -F 'client=Web' -F 'logArray=%7B%22l%22%3A%22hjAls6nLBlEqJwz4EqclX4zXP4oMz8uVGs%2F4WbPKzt7dHmrAnJ%2FsIacssRmwjEB0%22%2C%22iv%22%3A%225viydf5lxfjj1bhv%22%2C%22k%22%3A%22LYX8iLlULYvX2%2Buc%2Fji4p9TG%2BTWKjOtU2YQul%2BAbOZW9e4F8qh9dY%2B3l8Rtrr1srthobR0r3fKKa%2F5ZjR1lsa3cdOyhHIeiCWMYQTr0w7PQrwdSUTO%2FccP1Vo8x5AJ6i%2Bf8inJ64DlyUkyAoxKJP1NiK%2B8PmKzqnuugktJwbeR8%3D%22%2C%22v%22%3A1%7D'
