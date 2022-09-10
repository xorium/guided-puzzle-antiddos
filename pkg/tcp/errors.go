package tcp

type ConnectionReadError struct {
	Msg string
}

func (e ConnectionReadError) Error() string {
	if e.Msg == "" {
		return "connection read error"
	}

	return e.Msg
}

type ConnectionWriteError struct {
	Msg string
}

func (e ConnectionWriteError) Error() string {
	if e.Msg == "" {
		return "connection write error"
	}

	return e.Msg
}

type ConnectionEstablishingError struct {
	Msg string
}

func (e ConnectionEstablishingError) Error() string {
	if e.Msg == "" {
		return "connection establishing error"
	}

	return e.Msg
}

type CodecError struct {
	Msg string
}

func (e CodecError) Error() string {
	if e.Msg == "" {
		return "encoding/decoding error"
	}

	return e.Msg
}
