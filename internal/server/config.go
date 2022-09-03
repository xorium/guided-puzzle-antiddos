package server

type Config struct {
	// AESKey is a string with 32 ASCII characters len (AES-256 bit using)
	AESKey string `required:"true"`
	// MaxTourLen is the maximum tour length when the server switched in anti-DDOS mode
	MaxTourLen int `default:"10"`
	// RPSAllowed is a maximum RPS allowed until to switch to anti-DDOS mode
	RPSAllowed int `default:"10"`
}
