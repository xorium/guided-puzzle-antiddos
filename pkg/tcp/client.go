package tcp

import (
	"context"
	"encoding/binary"
	"fmt"
	"net"
	"time"

	"github.com/pkg/errors"
	"github.com/vmihailenco/msgpack/v5"
)

// Client is a TCP agent trying to get some resource from the server.
// It can execute guided tour protocol if resource server asks for.
type Client struct {
	dialer net.Dialer
	conn   net.Conn
	addr   string
}

// NewClient returns TCP crypto to send custom messages.
// The argument listenAddr must be in ip:port format.
func NewClient(addr string) *Client {
	return &Client{
		addr: addr,
		dialer: net.Dialer{
			Timeout:       8 * time.Second,
			FallbackDelay: time.Second,
			KeepAlive:     10 * time.Second,
		},
	}
}

// Connect tries to connect to the TCP server.
func (c *Client) Connect(timeout time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	conn, err := c.dialer.DialContext(ctx, "tcp", c.addr)
	if err != nil {
		return errors.Wrapf(
			ConnectionEstablishingError{
				Msg: fmt.Sprintf("can't connect to %s TCP address", c.addr),
			},
			err.Error(),
		)
	}

	c.conn = conn

	return nil
}

// SendMessage sends msg of the particular type appending at the beginning
// 2 bytes which of raw marshalled msg size (big-endian ordered).
func (c *Client) SendMessage(msg any) error {
	rawMsg, err := msgpack.Marshal(msg)
	if err != nil {
		return errors.Wrap(
			CodecError{
				Msg: fmt.Sprintf("can't marshal to MsgPack message %+v", msg),
			}, err.Error(),
		)
	}
	// The first 2 bytes in big-endian order (the most popular endianess at
	// the network transmissions).
	msgLen := uint16(len(rawMsg))
	output := make([]byte, 2)
	binary.BigEndian.PutUint16(output, msgLen)

	output = append(output, rawMsg...)

	return send(c.conn, output)
}

// ReceiveRawMsg receives raw bytes of an any message where first 2 bytes is
// big-endian ordered message raw size.
func (c *Client) ReceiveRawMsg() ([]byte, error) {
	buf, err := receive(c.conn, 2)
	if err != nil {
		return nil, err
	}

	msgLen := binary.BigEndian.Uint16(buf)

	return receive(c.conn, int(msgLen))
}
