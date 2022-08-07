# minecraft-client

[![License Badge]][License]
[![Build Badge]][Build]
[![GoDoc Badge]][GoDoc]

A client for the Minecraft RCON protocol.

## Library Usage

```go
// Create a new client and connect to the server.
client, err := minecraft.NewClient(minecraft.ClientOptions{
	Hostport: "127.0.0.1:25575",
	Timeout: 5 * time.Second, // Optional, this is the default value.
})
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

## Upgrading from 1.x to 2.x

The client is now instantiated with a ClientOptions struct in order to allow future client changes to be nonbreaking.

Change `NewClient("hostport")` to `NewClient(ClientOptions{Hostport: "hostport"})` to upgrade.

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

[License]: https://www.gnu.org/licenses/gpl-3.0
[License Badge]: https://img.shields.io/badge/License-GPLv3-blue.svg
[Build]: https://github.com/willroberts/minecraft-client/actions/workflows/build.yaml
[Build Badge]: https://github.com/willroberts/minecraft-client/actions/workflows/build.yaml/badge.svg
[GoDoc]: https://pkg.go.dev/github.com/willroberts/minecraft-client
[GoDoc Badge]: https://pkg.go.dev/badge/github.com/willroberts/minecraft-client