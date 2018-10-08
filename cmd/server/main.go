package main

import (
	"fmt"
	"log"
	"word-tokenize-in1118/internal"
	"word-tokenize-in1118/internal/constant"
	"word-tokenize-in1118/internal/layers/infrastructure/server"
)

func main() {
	protocol, rpc := internal.GetServerArgs()
	requestHandler := new(server.RequestHandler)

	if rpc {
		switch protocol {
		case constant.TCP:
			requestHandler.BringUpRPCTCPServer()
		default:
			// RPC over UDP is not supported by Golang
			// https://astaxie.gitbooks.io/build-web-application-with-golang/en/08.4.html
			// https://ipfs.io/ipfs/QmfYeDhGH9bZzihBUDEQbCbTc5k5FZKURMUoUvfmc27BwL/rpc/go_rpc.html
			log.Fatal(fmt.Errorf(constant.RPCUDPPatternError, constant.TCP))
		}
	} else {
		switch protocol {
		case constant.TCP:
			requestHandler.BringUpTCPServer()
		case constant.UDP:
			requestHandler.BringUpUDPServer()
		default:
			log.Fatal(fmt.Errorf(constant.ProtocolPatternError, constant.TCP, constant.UDP))
		}
	}
}
