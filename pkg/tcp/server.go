package tcp

import (
	"fmt"
	"net"
	"time"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func init() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
}

func defaultConnHandler(c net.Conn) {
	log.Info().Str("addr", c.RemoteAddr().String()).Msg("has been connected")
}

func closeByTimeout(c net.Conn, timeout time.Duration) {
	time.AfterFunc(timeout, func() { _ = c.Close() })
}

// Server is a simple TCP server which accepts connections, passes them to the
// handlers with timeout of handling.
type Server struct {
	listenAddr string
	handler    func(conn net.Conn)
}

// NewServer returns TCP server.
// The argument listenAddr must be in ip:port format.
func NewServer(listenAddr string, handler func(net.Conn)) *Server {
	return &Server{
		listenAddr: listenAddr,
		handler:    defaultConnHandler,
	}
}

// Listen starts accepting connections and pass them to the s.handler with
// maximum allowed time timeout.
func (s *Server) Listen(timeout time.Duration) error {
	defer func() {
		if r := recover(); r != nil {
			log.Panic().Msg(fmt.Sprintf("%v", r))
		}
	}()

	if timeout <= 0 {
		timeout = 30 * time.Second
	}

	l, err := net.Listen("tcp", s.listenAddr)
	defer func() { _ = l.Close() }()

	if err != nil {
		return errors.Wrap(err, "can't start listening on "+s.listenAddr)
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			return errors.Wrap(err, "can't accept new connection on "+s.listenAddr)
		}

		go closeByTimeout(conn, timeout)
		go s.handler(conn)
	}
}
