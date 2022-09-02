package tcp

import (
	"context"
	"encoding/binary"
	"github.com/pkg/errors"
	"github.com/vmihailenco/msgpack/v5"
	"net"
	"time"
)

type Client struct {
	dialer        net.Dialer
	conn          net.Conn
	addr          string
	cancelDialCtx func()
}

// NewClient returns TCP client to send custom messages.
// The argument addr must be in ip:port format.
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
		return errors.Wrapf(err, "can't connect to %s TCP address", c.addr)
	}

	c.conn = conn

	return nil
}

// SendMessage sends msg of the particular type.
func (c *Client) SendMessage(msg any, tp MsgType) error {
	rawMsg, err := msgpack.Marshal(msg)
	if err != nil {
		return errors.Wrapf(err, "can't marshal to MsgPack message %+v", msg)
	}
	// The first 2 bytes in big-endian order (the most popular endianess at
	// the network transmissions).
	msgLen := uint16(len(rawMsg))
	// Third 1 byte as identifier of the message type.
	output := make([]byte, 3)
	binary.BigEndian.PutUint16(output, msgLen)

	output[2] = byte(tp)
	output = append(output, rawMsg...)

	n, err := c.conn.Write(rawMsg)
	if n != len(output) {
		return errors.Errorf("%d bytes have been sent only instead of %d", n, len(output))
	}

	return nil
}

func (c *Client) ReceiveMsg(timeout time.Duration) error {
	c.conn.SetReadDeadline()
}
