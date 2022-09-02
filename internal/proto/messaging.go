package proto

import (
	"encoding/binary"
	"io"
	"net"
)

tyep MsgType byte

const (
	MsgTypeRegister MsgType = iota
	MsgTypeTourFailed
	MsgTypeSVerified
	MSgType
)

func sendRawMessage(msg []byte, conn net.Conn) error {
	// The first 2 bytes in big-endian order (the most popular endianess at
	// the network transmissions).
	msgLen := uint16(len(msg))
	lenBytes := make([]byte, 2)
	binary.BigEndian.PutUint16(lenBytes, msgLen)
	
	rawMsg := append(lenBytes, msg...)
	io.
		EOF
	
	conn.Write(rawMsg)
}
