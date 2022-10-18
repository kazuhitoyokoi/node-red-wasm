#echo "package main; func flows() string { return \`" > flows.go
#cat $1 >> flows.go
#echo "\`}" >> flows.go

echo "package main; func flows() string { return \`" > flows.go
cat examples/console.json >> flows.go
echo "\`}" >> flows.go
GOOS=js GOARCH=wasm go build -o docs/console.wasm

echo "package main; func flows() string { return \`" > flows.go
cat examples/worldmap.json >> flows.go
echo "\`}" >> flows.go
GOOS=js GOARCH=wasm go build -o docs/worldmap.wasm

#go run red.go flows.go
python3 -m http.server 8080 &
open http://localhost:8080/docs
