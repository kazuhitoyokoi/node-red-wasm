go run red.go
GOOS=js GOARCH=wasm go build -o red.wasm
python3 -m http.server 8080 &
open http://localhost:8080
