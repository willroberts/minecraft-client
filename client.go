package minecraft

import (
	"errors"
	"net"
)

// Client manages a connection to a Minecraft server.
type Client struct {
	conn          net.Conn
	lastRequestID int32
}

// NewClient creates a TCP connection to a Minecraft server.
func NewClient(hostport string) (*Client, error) {
	conn, err := net.Dial("tcp", hostport)
	if err != nil {
		return nil, err
	}

	return &Client{conn: conn}, nil
}

// Close disconnects from the server.
func (c *Client) Close() error {
	return c.conn.Close()
}

// Authenticate starts a logged-in RCON session.
// TODO: Obfuscate password.
func (c *Client) Authenticate(password string) error {
	return c.SendMessage(msgAuthenticate, password)
}

// SendMessage sends an RCON command to the server.
func (c *Client) SendMessage(msgType messageType, msg string) error {
	encoded, err := encode(msgType, []byte(msg), c.lastRequestID+1)
	if err != nil {
		return err
	}
	c.lastRequestID++

	// Send the message to the server.
	if _, err := c.conn.Write(encoded); err != nil {
		return err
	}

	// Read the response.
	resp := make([]byte, 14)
	if _, err := c.conn.Read(resp); err != nil {
		return err
	}

	_, id, msgType, err := decode(resp)
	if err != nil {
		return err
	}

	if id != c.lastRequestID || msgType != msgCommand {
		return errors.New("failed to authenticate")
	}

	return nil
}
