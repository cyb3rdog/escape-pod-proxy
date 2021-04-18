// Cyb3rVector EscapePod Proxy
//  by cyb3rdog
//
// (c) 2021 Vaclav Macha
// All rights reserved.
//
// mongoclient - handles the communication with mongo db
//
package mongoclient

import (
	"log"

	"github.com/digital-dream-labs/hugh/config"
)

type options struct {
	name string
	host string
	port string
	user string
	pass string
}

// Option is the list of options
type Option func(*options)

// WithViper loads a config from environment variables.
func WithViper(args ...string) Option {
	return func(o *options) {
		if err := o.viperize(args...); err != nil {
			log.Fatal(err)
		}
	}
}

func (o *options) viperize(args ...string) error {
	v, err := config.New("DB", args...)
	if err != nil {
		return err
	}

	for i, j := 0, 1; j < len(args); i, j = i+2, j+2 {
		key, val := args[i], args[j]
		switch key {
		case "env-prefix":
			v.SetEnvPrefix(val)
		}
	}

	if v.IsSet("name") {
		o.name = v.GetString("name")
	}
	if v.IsSet("host") {
		o.host = v.GetString("host")
	}
	if v.IsSet("port") {
		o.port = v.GetString("port")
	}
	if v.IsSet("user") {
		o.user = v.GetString("user")
	}
	if v.IsSet("pass") {
		o.pass = v.GetString("pass")
	}

	return nil
}
