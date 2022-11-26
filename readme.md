# Equal Experts Backend Tech Test

The code is quite simple and self explanatory. It's split in command and position parsing and a navigator package which contains the logic to move the rover over the coordinates. The test coverage is pretty high, close to 100%. Currently it supports one command per run. Can be easily extended to supports multiple commands. Please use the instructions below to run and test.

## HOW TO USE

You do not need any external dependencies in order to run this program. The only thing you need is a working Go installation. I have used `go1.18 darwin/arm64` but any other latest version should work.

### Build

```
$ go build -o navigator cmd/main.go
```

### Test 

```
$ go test -v ./...
```

### Run

```
$ ./ticker_server -start=1,2,EAST -command=FLFFFRFLB
(6, 4, NORTH)
```
