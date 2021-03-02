# minecraft-client

A client for the Minecraft RCON API.

## Usage

```go
client, err := NewClient("127.0.0.1:25575")
if err != nil {
	log.Fatal(err)
}
defer client.Close()

if err := client.Authenticate("my_password"); err != nil {
	log.Fatal(err)
}

resp, err := client.SendCommand("seed")
if err != nil {
	log.Fatal(err)
}

log.Println(resp) // "Seed: [-2474125574890692308]"
```

## Limitations

Response bodies over 4KB will be truncated.

## Starting a server for testing

```
docker pull itzg/minecraft-server
docker run --name=minecraft-server -p 25575:25575 -d -e EULA=TRUE itzg/minecraft-server
```

## Running Tests

After starting the test server in Docker:

```
go test -v
```

## Reference

- https://wiki.vg/Rcon