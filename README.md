# minecraft-client

[![GoDoc Badge]][GoDoc]
[![Travis Badge]][Travis]
[![License Badge]][License]

A client for the Minecraft RCON protocol.

## Library Usage

```go
// Create a new client and connect to the server.
client, err := minecraft.NewClient("127.0.0.1:25575")
if err != nil {
	log.Fatal(err)
}
defer client.Close()

// Send some commands.
if err := client.Authenticate("password"); err != nil {
	log.Fatal(err)
}
resp, err := client.SendCommand("seed")
if err != nil {
	log.Fatal(err)
}
log.Println(resp.Body) // "Seed: [1871644822592853811]"
```

## Shell Utility

If you are looking for a tool rather than a library, try the shell command:

```bash
$ cd cmd/shell
$ go run main.go --hostport 127.0.0.1:25575 --password minecraft
Starting RCON shell. Use 'exit', 'quit', or Ctrl-C to exit.
> list
There are 0 of a max of 20 players online:
> seed
Seed: [5853448882787620410]
```

## Limitations

Response bodies over 4KB will be truncated.

## Starting a server for testing

```
$ docker pull itzg/minecraft-server
$ docker run --name=minecraft-server -p 25575:25575 -d -e EULA=TRUE itzg/minecraft-server
```

## Running Tests

To run unit tests:

```
$ go test -v
```

To run integration tests after starting the test server in Docker:

```
$ go test -v --tags=integration
```

## Reference

- https://wiki.vg/Rcon

[GoDoc]: https://pkg.go.dev/github.com/willroberts/minecraft-client
[GoDoc Badge]: https://pkg.go.dev/badge/github.com/willroberts/minecraft-client
[Travis]: https://travis-ci.org/willroberts/minecraft-client
[Travis Badge]: https://api.travis-ci.org/willroberts/minecraft-client.svg?branch=main
[License]: https://www.gnu.org/licenses/gpl-3.0
[License Badge]: https://img.shields.io/badge/License-GPLv3-blue.svg