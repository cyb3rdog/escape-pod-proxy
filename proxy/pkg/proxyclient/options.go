// Cyb3rVector EscapePod Proxy
//  by cyb3rdog
//
// (c) 2021 Vaclav Macha
// All rights reserved.
//
// proxyclient - handles the incomming connection from the escapepod binary
//
package proxyclient

import (
	"log"

	interfaces "github.com/cyb3rdog/escape-pod-proxy/proxy/pkg"
	"github.com/digital-dream-labs/hugh/config"
)

type options struct {
	proxy  interfaces.ProxyTarget
	prefix string
}

// Option is the list of options
type Option func(*options)

// WithProxyTo sets the proxy target server
func WithProxyTo(s interfaces.ProxyTarget) Option {
	return func(o *options) {
		o.proxy = s
	}
}

// WithViper loads a config from environment variables.
func WithViper(args ...string) Option {
	return func(o *options) {
		if err := o.viperize(args...); err != nil {
			log.Fatal(err)
		}
	}
}

func (o *options) viperize(args ...string) error {
	v, err := config.New("", args...)
	if err != nil {
		return err
	}

	for i, j := 0, 1; j < len(args); i, j = i+2, j+2 {
		key, val := args[i], args[j]
		switch key {
		case "env-prefix":
			o.prefix = val
			v.SetEnvPrefix(val)
		}
	}
	return nil
}
