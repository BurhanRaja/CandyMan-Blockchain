package network

type ServerOpts struct {
	Transports []Transport
}

type Server struct {
	ServerOpts
	rpcCh chan RPC
	quitCh chan struct{}
}