package minecraft

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncodeMessage(t *testing.T) {
	msg := Message{
		Length: int32(len("seed") + headerSize),
		ID:     1,
		Type:   MsgCommand,
		Body:   []byte("seed"),
	}

	b, err := EncodeMessage(msg)
	if err != nil {
		t.Fatal(err)
	}

	expected := []byte{
		// Request length: 14 bytes.
		14, 0, 0, 0,
		// Request ID: 1.
		1, 0, 0, 0,
		// Request type: 2 (msgCommand).
		2, 0, 0, 0,
		// Message: "seed".
		115, 101, 101, 100,
		// Terminator.
		0, 0,
	}

	assert.Equal(t, b, expected)
}

func TestDecodeMessage(t *testing.T) {
	b := []byte{
		// Response length: 38 bytes.
		38, 0, 0, 0,
		// Request ID: 2.
		2, 0, 0, 0,
		// Response type: 0 (msgResponse).
		0, 0, 0, 0,
		// Message: "Seed: [-2474125574890692308]".
		83, 101, 101, 100, 58, 32, 91, 45, 50, 52, 55, 52, 49, 50, 53, 53, 55, 52, 56, 57, 48, 54, 57, 50, 51, 48, 56, 93,
	}
	resp, err := DecodeMessage(b)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, resp.Length, int32(38))
	assert.Equal(t, resp.ID, int32(2))
	assert.Equal(t, resp.Type, MsgResponse)
	assert.Equal(t, string(resp.Body), "Seed: [-2474125574890692308]")
}
