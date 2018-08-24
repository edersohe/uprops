### Use

```sh
uprops "$(cat base.properties)" "$(cat custom.properties)" > baseUpdateWithCustom.properties
```

### Build binaries

```
env GOOS=linux GOARCH=amd64 go build -o uprops-linux-amd64 uprops.go
env GOOS=darwin GOARCH=amd64 go build -o uprops-darwin-amd64 uprops.go
env GOOS=windows GOARCH=amd64 go build -o uprops-windows-amd64 uprops.go
```