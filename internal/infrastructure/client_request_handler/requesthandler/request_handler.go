package requesthandler

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/rpc"
	"time"
	"word-tokenize-in1118/internal/communication"
	"word-tokenize-in1118/internal/constant"
	"word-tokenize-in1118/internal/util"
)

// RequestHandler ...
type RequestHandler struct{}

// TextTokenizeRPCTCP handles a remote procedure call over TCP to text tokenize
func (requestHandler *RequestHandler) TextTokenizeRPCTCP(text string) communication.Response {
	client := util.DialRPCTCPClient()
	log.Println(fmt.Sprintf(constant.SendingRequest, constant.RPC), text)
	response := sendRPC(text, client)
	client.Close()
	log.Println(fmt.Sprintf(constant.ReceivingResponse, constant.RPC), response)
	return response
}

// TextTokenizeTCP handles a TCP request to text tokenize
func (requestHandler *RequestHandler) TextTokenizeTCP(request communication.Request) communication.Response {
	connection := util.DialTCPConnection()
	defer connection.Close()
	log.Println(fmt.Sprintf(constant.SendingRequest, constant.TCP), request)
	response := send(request, connection)
	log.Println(fmt.Sprintf(constant.ReceivingResponse, constant.TCP), response)
	return response
}

// TextTokenizeUDP handles a UDP request to text tokenize
func (requestHandler *RequestHandler) TextTokenizeUDP(request communication.Request) communication.Response {
	connection := util.DialUDPConnection()
	defer connection.Close()
	log.Println(fmt.Sprintf(constant.SendingRequest, constant.UDP), request)
	response := send(request, connection)
	log.Println(fmt.Sprintf(constant.ReceivingResponse, constant.UDP), response)
	return response
}

// Helper: sends the text tokenize RPC request
func sendRPC(text string, client *rpc.Client) communication.Response {
	var response communication.Response
	var tokens []string
	now := time.Now()
	client.Call(constant.NLGTextTokenizeRPC, text, &tokens)
	elapsed := time.Since(now)
	response.Duration = elapsed
	response.Content = tokens
	return response
}

// Helper: sends a request
func send(request communication.Request, connection net.Conn) communication.Response {
	var response communication.Response
	now := time.Now()
	json.NewEncoder(connection).Encode(request)
	response = receive(json.NewDecoder(connection))
	elapsed := time.Since(now)
	response.Duration = elapsed
	return response
}

// Helper: receives a response
func receive(jsonDecoder *json.Decoder) communication.Response {
	var response communication.Response
	jsonDecoder.Decode(&response)
	return response
}
