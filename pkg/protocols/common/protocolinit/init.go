package protocolinit

import (
	"Ni/pkg/js/compiler"
	"Ni/pkg/protocols/common/protocolstate"
	"Ni/pkg/protocols/dns/dnsclientpool"
	"Ni/pkg/protocols/http/httpclientpool"
	"Ni/pkg/protocols/http/signerpool"
	"Ni/pkg/protocols/network/networkclientpool"
	"Ni/pkg/protocols/whois/rdapclientpool"
	"Ni/pkg/types"
)

// Init initializes the client pools for the protocols
func Init(options *types.Options) error {

	if err := protocolstate.Init(options); err != nil {
		return err
	}
	if err := dnsclientpool.Init(options); err != nil {
		return err
	}
	if err := httpclientpool.Init(options); err != nil {
		return err
	}
	if err := signerpool.Init(options); err != nil {
		return err
	}
	if err := networkclientpool.Init(options); err != nil {
		return err
	}
	if err := rdapclientpool.Init(options); err != nil {
		return err
	}
	if err := compiler.Init(options); err != nil {
		return err
	}
	return nil
}

func Close() {
	protocolstate.Dialer.Close()
}
