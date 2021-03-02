package minecraft

import (
	"bytes"
	"encoding/binary"
)

type messageType int32

type response struct {
	Length int32
	ID     int32
	Type   messageType
	Body   []byte
}

const (
	msgResponse     messageType = iota // 0: response.
	_                                  // 1: unused.
	msgCommand                         // 2: command.
	msgAuthenticate                    // 3: login.

	headerSize = 10 // 4-byte request ID, 4-byte message type, 2-byte terminator.
)

// encode serializes an RCON command.
func encode(msgType messageType, msg []byte, requestID int32) ([]byte, error) {
	buf := new(bytes.Buffer)
	for _, v := range []interface{}{
		int32(len(msg) + headerSize), // Request length.
		requestID,
		msgType,
		[]byte(msg),  // Payload.
		[]byte{0, 0}, // Terminator.
	} {
		if err := binary.Write(buf, binary.LittleEndian, v); err != nil {
			return nil, err
		}
	}
	return buf.Bytes(), nil
}

// decode deserialize an RCON response.
func decode(msg []byte) (response, error) {
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

	resp := response{
		Length: responseLength,
		ID:     responseID,
		Type:   responseType,
	}

	remainingBytes := responseLength - headerSize
	if remainingBytes > 0 {
		data := make([]byte, remainingBytes)
		if err := binary.Read(reader, binary.LittleEndian, &data); err != nil {
			return response{}, err
		}
		resp.Body = data
	}

	return resp, nil
}
