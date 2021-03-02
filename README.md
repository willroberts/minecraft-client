# minecraft-client

A client for the Minecraft RCON API.

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