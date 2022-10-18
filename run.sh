echo "package main; func flows() string { return \`" > flows.go
#cat $1 >> flows.go
cat examples/inject_debug_flows.json >> flows.go
echo "\`}" >> flows.go
GOOS=js GOARCH=wasm go build -o docs/red.wasm
go run red.go flows.go
python3 -m http.server 8080 &
open http://localhost:8080/docs
