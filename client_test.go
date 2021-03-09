package minecraft

import (
	"errors"
	"strings"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	testHost     = "127.0.0.1:25575" // Game port is 25565; RCON port is 25575.
	testPassword = "minecraft"
)

func TestNewClient(t *testing.T) {
	client, err := NewClient(testHost)
	if err != nil {
		t.Fatal(err)
	}
	defer client.Close()
}

func TestConnectionFailure(t *testing.T) {
	_, err := NewClient("127.0.0.1:25576")
	assert.NotNil(t, err)
}

func TestAuthenticate(t *testing.T) {
	client, err := NewClient(testHost)
	if err != nil {
		t.Fatal(err)
	}
	defer client.Close()

	if err := client.Authenticate(testPassword); err != nil {
		t.Fatal(err)
	}
}

func TestAuthenticationFailure(t *testing.T) {
	client, err := NewClient(testHost)
	if err != nil {
		t.Fatal(err)
	}
	defer client.Close()

	err = client.Authenticate("invalid_password")
	assert.Equal(t, err, errAuthenticationFailure)
}

func TestSendCommand(t *testing.T) {
	client, err := NewClient(testHost)
	if err != nil {
		t.Fatal(err)
	}
	defer client.Close()

	if err := client.Authenticate(testPassword); err != nil {
		t.Fatal(err)
	}

	out, err := client.SendCommand("seed")
	if err != nil {
		t.Fatal(err)
	}

	if !strings.HasPrefix(out, "Seed: [") {
		t.Fatal()
	}
}

func TestSendCommandAsync(t *testing.T) {
	client, err := NewClient(testHost)
	if err != nil {
		t.Fatal(err)
	}
	defer client.Close()

	if err := client.Authenticate(testPassword); err != nil {
		t.Fatal(err)
	}

	var wg sync.WaitGroup
	var errCh = make(chan error, 0)
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			out, err := client.SendCommand("seed")
			if err != nil {
				errCh <- err
			}

			if !strings.HasPrefix(out, "Seed: [") {
				errCh <- errors.New("bad seed")
			}
			wg.Done()
		}()
	}
	wg.Wait()

	select {
	case err := <-errCh:
		t.Fatalf("failed to send async command: %s", err)
	default:
		return
	}
}
