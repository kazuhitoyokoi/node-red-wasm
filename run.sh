node codegen.js ./examples/console.json > flows.go
GOOS=js GOARCH=wasm go build -o docs/console.wasm

node codegen.js ./examples/worldmap.json > flows.go
GOOS=js GOARCH=wasm go build -o docs/worldmap.wasm

#go run red.go flows.go
python3 -m http.server 8080 &
open http://localhost:8080/docs
