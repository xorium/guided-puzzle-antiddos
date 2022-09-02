package server

type Config struct {
	GuideIndex int `default:"0" required:"true"`
	// Number of the keys is the tour guides number too.
	// Current tour guide key is at index GuideIndex.
	GuidesKeys []string `default:"7r9XA2c0,0Sq2Q7jV,U4WwkrUq,Hg8a385F"`
	// Address format is IP:Port, the addresses number is the same as the keys.
	// Current tour guide address is at index GuideIndex.
	GuidesAddrs []string `default:"localhost:5001,localhost:5002,localhost:5003,localhost:5004"`
	// Address of the main server in format IP:Port.
	GuidesMainServer string `default:"localhost:5000"`
}
