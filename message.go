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

func encode(msgType messageType, msg []byte, requestID int32) ([]byte, error) {
	buf := new(bytes.Buffer)
	// Request length,
	if err := binary.Write(buf, binary.LittleEndian, int32(len(msg)+10)); err != nil {
		return nil, err
	}
	// Request ID.
	if err := binary.Write(buf, binary.LittleEndian, requestID); err != nil {
		return nil, err
	}
	// Message type.
	if err := binary.Write(buf, binary.LittleEndian, msgType); err != nil {
		return nil, err
	}
	// Payload.
	if err := binary.Write(buf, binary.LittleEndian, []byte(msg)); err != nil {
		return nil, err
	}
	// Terminator.
	if err := binary.Write(buf, binary.LittleEndian, []byte("\x00\x00")); err != nil {
		return nil, err
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

	return responseLength, responseID, responseType, nil
}
