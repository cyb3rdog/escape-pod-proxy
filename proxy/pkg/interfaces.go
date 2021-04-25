// Cyb3rVector EscapePod Proxy
//  by cyb3rdog
//
// (c) 2021 Vaclav Macha
// All rights reserved.
//
// interfaces - base interfaces the packages are interfaced with
// i.e. the IntentFactory could be MongoDB client or HttpClient
//
package interfaces

import (
	"context"

	"github.com/cyb3rdog/escape-pod-proxy/proto/lang/go/podextension"
)

type ProxyTarget interface {
	Push(context.Context, *podextension.UnaryReq) (string, map[string]string, error)
}

type IntentFactory interface {
	InsertIntent(intent_json string) (string, error)
	DeleteIntent(intent_id string) (bool, error)
	SelectIntents(filter_query string) (map[string]string, error)
}
