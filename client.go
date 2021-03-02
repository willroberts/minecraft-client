package minecraft

import (
	"errors"
	"net"
)

const maxResponseSize = 4110 // https://wiki.vg/Rcon#Fragmentation

// Client manages a connection to a Minecraft server.
type Client struct {
	conn        net.Conn
	idGenerator *idGenerator
}

// NewClient creates a TCP connection to a Minecraft server.
func NewClient(hostport string) (*Client, error) {
	conn, err := net.Dial("tcp", hostport)
	if err != nil {
		return nil, err
	}

	return &Client{conn: conn, idGenerator: &idGenerator{}}, nil
}

// Close disconnects from the server.
func (c *Client) Close() error {
	return c.conn.Close()
}

// Authenticate starts a logged-in RCON session.
func (c *Client) Authenticate(password string) error {
	resp, err := c.sendMessage(msgAuthenticate, password)
	if err != nil {
		return err
	}

	if resp.Type != msgCommand {
		return errors.New("failed to authenticate")
	}

	return nil
}

// SendCommand sends an RCON command to the server.
func (c *Client) SendCommand(command string) (string, error) {
	resp, err := c.sendMessage(msgCommand, command)
	if err != nil {
		return "", err
	}
	return string(resp.Body), nil
}

func (c *Client) sendMessage(msgType messageType, msg string) (response, error) {
	requestID := c.idGenerator.GenerateID()

	encoded, err := encode(msgType, []byte(msg), requestID)
	if err != nil {
		return response{}, err
	}

	if _, err := c.conn.Write(encoded); err != nil {
		return response{}, err
	}

	respBytes := make([]byte, maxResponseSize)
	if _, err := c.conn.Read(respBytes); err != nil {
		return response{}, err
	}

	resp, err := decode(respBytes)
	if err != nil {
		return response{}, err
	}

	if resp.ID != requestID {
		return response{}, errors.New("invalid response ID")
	}

	return resp, nil
}
