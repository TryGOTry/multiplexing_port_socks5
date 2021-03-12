package protocols
// NewSOCKS5Protocol initializes a Protocol with a SOCKS5 signature.
func NewSOCKS5Protocol(targetAddress string) *Protocol {
	return &Protocol{
		Name:            "SOCKS5",
		Target:          targetAddress,
		MatchStartBytes: [][]byte{{0x05}},
	}
}

