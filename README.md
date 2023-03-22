# Go_WebAssembly



# Run
1. go to assets
2. run: `cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" .`
3. go to cmd/wasm
4. Compile with command: `GOOS=js GOARCH=wasm go build -o  ../../assets/json.wasm`
5. go to cmd/server
6. run: `go run server.go`
7. access localhost:8080