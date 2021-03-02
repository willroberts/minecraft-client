package minecraft

import "testing"

var (
	testHost     = "127.0.0.1:25575" // Game port is 25565; RCON port is 25575.
	testPassword = "minecraft"
)

func TestNewClient(t *testing.T) {
	client, err := NewClient(testHost)
	if err != nil {
		t.Fatalf("%v", err)
	}
	defer client.Close()
}

func TestAuthenticate(t *testing.T) {
	client, err := NewClient(testHost)
	if err != nil {
		t.Fatalf("%v", err)
	}
	defer client.Close()

	if err := client.Authenticate(testPassword); err != nil {
		t.Fatalf("%v", err)
	}
}
