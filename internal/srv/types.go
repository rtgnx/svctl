package srv

type ServerConfig struct {
	TLSCertFile string
	TLSKeyFile  string
	UseTLS      bool
	Addr        string
}
