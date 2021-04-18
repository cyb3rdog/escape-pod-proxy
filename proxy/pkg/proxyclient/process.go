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
	"context"

	"github.com/cyb3rdog/escape-pod-proxy/proto/lang/go/podextension"
)

// Extension server entry point with UnaryRequest to proxied as a new event pushed to event stream
func (client *ProxyClient) Unary(context context.Context, request *podextension.UnaryReq) (*podextension.UnaryResp, error) {

	key, params, err := client.proxy.Push(context, request)
	if err != nil {
		return nil, err
	}

	response := &podextension.UnaryResp{
		IntentName: key,
		Parameters: params,
	}
	return response, nil
}
