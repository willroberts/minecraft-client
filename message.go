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

func decode(msg []byte) (int32, int32, messageType, error) {
	// Decode the response.
	reader := bytes.NewReader(msg)
	var responseLength int32
	if err := binary.Read(reader, binary.LittleEndian, &responseLength); err != nil {
		return 0, 0, 0, err
	}

	var responseID int32
	if err := binary.Read(reader, binary.LittleEndian, &responseID); err != nil {
		return 0, 0, 0, err
	}

	var responseType messageType
	if err := binary.Read(reader, binary.LittleEndian, &responseType); err != nil {
		return 0, 0, 0, err
	}

	// TODO: Read response payload for non-auth messages.

	return responseLength, responseID, responseType, nil
}
