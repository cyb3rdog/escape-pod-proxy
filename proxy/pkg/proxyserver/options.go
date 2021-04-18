// Cyb3rVector EscapePod Proxy
//  by cyb3rdog
//
// (c) 2021 Vaclav Macha
// All rights reserved.
//
// proxyserver - handles the incomming connection from the clients and
// notifies them about the intent events raised from the escape-pod.
//
package proxyserver

import (
	"log"

	interfaces "github.com/cyb3rdog/escape-pod-proxy/proxy/pkg"
	"github.com/digital-dream-labs/hugh/config"
)

// Option is the list of options
type Option func(*options)

type options struct {
	prefix  string
	factory interfaces.IntentFactory
}

// WithIntentFactory adds the factory for altering intents
func WithIntentFactory(factory interfaces.IntentFactory) Option {
	return func(o *options) {
		o.factory = factory
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
