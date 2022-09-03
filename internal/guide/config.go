package guide

type Config struct {
	// AESKey is a string with 32 ASCII characters len (AES-256 bit using)
	AESKey string `required:"true"`
}
