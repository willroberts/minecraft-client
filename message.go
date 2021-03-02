package minecraft

import (
	"bytes"
	"encoding/binary"
)

type messageType int32

const (
	msgResponse     messageType = iota // multi-packet response.
	msgUnused                          // no known use case.
	msgCommand                         // run a command.
	msgAuthenticate                    // authenticate.
)

var terminator = []byte{0, 0}

func encode(msgType messageType, msg []byte, requestID int32) ([]byte, error) {
	buf := new(bytes.Buffer)
	for _, v := range []interface{}{
		int32(len(msg) + 10), // Request length.
		requestID,
		msgType,
		[]byte(msg), // Payload.
		terminator,
	} {
		if err := binary.Write(buf, binary.LittleEndian, v); err != nil {
			return nil, err
		}
	}
	return buf.Bytes(), nil
}

type response struct {
	Length int32
	ID     int32
	Type   messageType
	Body   []byte
}

func decode(msg []byte) (response, error) {
	// Decode the response.
	reader := bytes.NewReader(msg)
	var responseLength int32
	if err := binary.Read(reader, binary.LittleEndian, &responseLength); err != nil {
		return response{}, err
	}

	var responseID int32
	if err := binary.Read(reader, binary.LittleEndian, &responseID); err != nil {
		return response{}, err
	}

	var responseType messageType
	if err := binary.Read(reader, binary.LittleEndian, &responseType); err != nil {
		return response{}, err
	}

	// TODO: Read response payload for non-auth messages.

	return response{
		Length: responseLength,
		ID:     responseID,
		Type:   responseType,
	}, nil
}
