package main

import (
	"github.com/tokenbankteam/eos-go/p2p"
)

func main() {

	proxy := p2p.Proxy{
		Routes: []*p2p.Route{
			{From: ":19876", To: "patrick.testnets.tokenbankteam.com:9876"},
		},
		Handlers: []p2p.Handler{
			//p2p.StringLoggerHandler,
			p2p.LoggerHandler,
		},
	}

	proxy.Start()

}
