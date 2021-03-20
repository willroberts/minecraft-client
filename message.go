package minecraft

import (
	"bytes"
	"encoding/binary"
)

// MessageType is an int32 representing the type of message being sent or received.
type MessageType int32

// Message contains fields for RCON messages.
type Message struct {
	Length int32
	ID     int32
	Type   MessageType
	Body   string
}

const (
	// MsgResponse is returned by the server.
	MsgResponse MessageType = iota
	_
	// MsgCommand is used when sending commands to the server.
	MsgCommand
	// MsgAuthenticate is used when logging into the server.
	MsgAuthenticate

	headerSize = 10 // 4-byte request ID, 4-byte message type, 2-byte terminator.
)

// EncodeMessage serializes an RCON command.
// Format: [4-byte message size | 4-byte message ID | 4-byte message type | variable length message | 2-byte terminator].
func EncodeMessage(msg Message) ([]byte, error) {
	buf := new(bytes.Buffer)
	for _, v := range []interface{}{
		msg.Length,
		msg.ID,
		msg.Type,
		[]byte(msg.Body),
		[]byte{0, 0}, // 2-byte terminator.
	} {
		if err := binary.Write(buf, binary.LittleEndian, v); err != nil {
			return nil, err
		}
	}
	return buf.Bytes(), nil
}

// DecodeMessage deserialize an RCON response.
// Format: [4-byte message size | 4-byte message ID | 4-byte message type | variable length message].
func DecodeMessage(msg []byte) (Message, error) {
	reader := bytes.NewReader(msg)

	var responseLength int32
	if err := binary.Read(reader, binary.LittleEndian, &responseLength); err != nil {
		return Message{}, err
	}

	var responseID int32
	if err := binary.Read(reader, binary.LittleEndian, &responseID); err != nil {
		return Message{}, err
	}

	var responseType MessageType
	if err := binary.Read(reader, binary.LittleEndian, &responseType); err != nil {
		return Message{}, err
	}

	resp := Message{
		Length: responseLength,
		ID:     responseID,
		Type:   responseType,
	}

	remainingBytes := responseLength - headerSize
	if remainingBytes > 0 {
		data := make([]byte, remainingBytes)
		if err := binary.Read(reader, binary.LittleEndian, &data); err != nil {
			return Message{}, err
		}
		resp.Body = string(data)
	}

	return resp, nil
}
