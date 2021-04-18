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
	"fmt"

	"github.com/cyb3rdog/escape-pod-proxy/proto/lang/go/podextension"
)

// Unary checks the function map for an appropriate processor and either processes it or returns an error
func (client *ProxyClient) Unary(context context.Context, request *podextension.UnaryReq) (*podextension.UnaryResp, error) {

	fmt.Println(request)

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
