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
	"context"
	"crypto/rand"
	"fmt"
	"log"
	"time"

	"github.com/cyb3rdog/escape-pod-proxy/proto/lang/go/cybervector"
	"github.com/cyb3rdog/escape-pod-proxy/proto/lang/go/podextension"
)

func (server *ProxyServer) GenerateGuid() (string, error) {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	uuid := fmt.Sprintf("%x-%x-%x-%x-%x",
		b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
	return uuid, nil
}

// Pushes the Intent through the proxy channel to all subscribed clients
func (server *ProxyServer) Push(ctx context.Context, request *podextension.UnaryReq) (string, map[string]string, error) {

	for uuid, stream := range server.subscribers {

		response := cybervector.ProxyMessaage{
			MessageType: cybervector.MessageType_ProcessIntent,
			MessageData: fmt.Sprintf("{'uuid':'%s','message':'Intent Received from EscapePod'}", uuid),
			IntentName:  request.IntentName,
			Parameters:  request.Parameters,
		}

		if err := stream.Send(&response); err != nil {
			log.Println("Closing Event stream (on send):", err)
			return "", nil, err
		}
	}

	return "", nil, nil
}

func (server *ProxyServer) Subscribe(request *cybervector.SubscribeRequest, stream cybervector.CyberVectorProxyService_SubscribeServer) error {

	uuid, err := server.GenerateGuid()
	if err != nil {
		log.Fatal(err)
	}

	server.subscribers[uuid] = stream
	response := cybervector.ProxyMessaage{
		MessageType: cybervector.MessageType_Subscribed,
		MessageData: fmt.Sprintf("{'uuid':'%s','message':'Successfully subscribed to Cyb3rVector EscapePod extension'}", uuid),
	}
	if err := stream.Send(&response); err != nil {
		log.Printf("send error %v", err)
		return err
	} else if err = stream.Context().Err(); err != nil {
		log.Println("Closing Event stream:", err)
		return err
	}

	done := make(chan bool)
	var interval time.Duration = time.Duration(30 * time.Second)
	if request.KeepAlive > 1 {
		interval = time.Duration(request.KeepAlive) * time.Second
	}
	pingTicker := time.NewTicker(interval)
	keepAlive := cybervector.ProxyMessaage{
		MessageType: cybervector.MessageType_KeepAlive,
	}

	for {
		select {
		case <-done:
			return nil
		case time := <-pingTicker.C:
			keepAlive.MessageData = fmt.Sprintf("{'time':'%s'}", time)
			if err := stream.Send(&keepAlive); err != nil {
				log.Println("Closing Event stream (on send):", err)
				return err
			} else if err = stream.Context().Err(); err != nil {
				log.Println("Closing Event stream:", err)
				return err
			}
		}
	}
}

func (server *ProxyServer) UnSubscribe(context context.Context, request *cybervector.UnsubscribeRequest) (*cybervector.ProxyMessaage, error) {

	delete(server.subscribers, "test")
	response := cybervector.ProxyMessaage{
		MessageType: cybervector.MessageType_Unsubscribed,
		MessageData: "{'message':'Successfully unsubscribed from Cyb3rVector EscapePod extension'}",
	}

	return &response, nil
}

func (server *ProxyServer) SelectIntents(context context.Context, request *cybervector.SelectIntentRequest) (*cybervector.SelectIntentResponse, error) {

	intent_data := make(map[string]string)
	response_message := cybervector.ResponseMessage{}

	json_data, err := server.factory.SelectIntents(request.FilterJson)
	if err != nil {
		response_message.Code = cybervector.ResponseMessage_FAILURE
		response_message.Message = err.Error()
	} else {
		response_message.Code = cybervector.ResponseMessage_SUCCESS
		response_message.Message = "OK"
		intent_data = json_data
	}

	response := cybervector.SelectIntentResponse{
		Response:   &response_message,
		FilterJson: request.FilterJson,
		IntentData: intent_data,
	}

	return &response, nil
}

func (server *ProxyServer) InsertIntent(context context.Context, request *cybervector.InsertIntentRequest) (*cybervector.InsertIntentResponse, error) {

	response_message := cybervector.ResponseMessage{}

	err := server.factory.InsertIntent(request.IntentData)
	if err != nil {
		response_message.Code = cybervector.ResponseMessage_FAILURE
		response_message.Message = err.Error()
	} else {
		response_message.Code = cybervector.ResponseMessage_SUCCESS
		response_message.Message = "OK"
	}

	response := cybervector.InsertIntentResponse{
		Response: &response_message,
	}

	return &response, nil
}

func (server *ProxyServer) DeleteIntent(context context.Context, request *cybervector.DeleteIntentRequest) (*cybervector.DeleteIntentResponse, error) {

	response_message := cybervector.ResponseMessage{}

	err := server.factory.DeleteIntent(request.IntentId)
	if err != nil {
		response_message.Code = cybervector.ResponseMessage_FAILURE
		response_message.Message = err.Error()
	} else {
		response_message.Code = cybervector.ResponseMessage_SUCCESS
		response_message.Message = "OK"
	}

	response := cybervector.DeleteIntentResponse{
		Response: &response_message,
	}

	return &response, nil
}
