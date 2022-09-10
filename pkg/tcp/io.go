package tcp

import (
	"fmt"
	"net"

	"github.com/pkg/errors"
)

// send write data to the TCP connection or returns particular kind of the error.
func send(conn net.Conn, data []byte) error {
	n, err := conn.Write(data)
	dataLen := len(data)

	if n != dataLen {
		return ConnectionWriteError{
			Msg: fmt.Sprintf("%d bytes have been sent only instead of %d", n, dataLen),
		}
	}

	if err != nil {
		errMsg := fmt.Sprintf(
			"can't receive %d bytes from %v", len(data), conn.RemoteAddr(),
		)

		return errors.Wrap(ConnectionWriteError{Msg: errMsg}, err.Error())
	}

	return nil
}

// receive reads numOfBytes from TCP connection or returns particular kind of the error.
func receive(conn net.Conn, numOfBytes int) ([]byte, error) {
	buf := make([]byte, numOfBytes)

	n, err := conn.Read(buf)
	if err != nil {
		errMsg := fmt.Sprintf(
			"can't receive %d bytes from %v", len(buf), conn.RemoteAddr(),
		)

		return nil, errors.Wrap(ConnectionReadError{Msg: errMsg}, err.Error())
	}

	if n != len(buf) {
		return nil, ConnectionReadError{
			Msg: fmt.Sprintf("%d bytes have been received instead of %d", n, len(buf)),
		}
	}

	return buf, nil
}
