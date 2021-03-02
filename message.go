package minecraft

import (
	"bytes"
	"encoding/binary"
	"log"
)

type messageType int32

const (
	msgResponse     messageType = iota // 0: response.
	_                                  // 1: unused.
	msgCommand                         // 2: command.
	msgAuthenticate                    // 3: login.

	headerSize = 10 // 4-byte request ID, 4-byte message type, 2-byte terminator.
)

var terminator = []byte{0, 0}

func encode(msgType messageType, msg []byte, requestID int32) ([]byte, error) {
	buf := new(bytes.Buffer)
	for _, v := range []interface{}{
		int32(len(msg) + headerSize), // Request length.
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
	reader := bytes.NewReader(msg)

	// TODO: Consider parsing directly into a &response{} object.
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

	if responseLength-headerSize > 0 {
		log.Println("there was more data to be read!")
		var discarded int16 // Terminator.
		if err := binary.Read(reader, binary.LittleEndian, &discarded); err != nil {
			return response{}, err
		}
	}

	return response{
		Length: responseLength,
		ID:     responseID,
		Type:   responseType,
	}, nil
}
