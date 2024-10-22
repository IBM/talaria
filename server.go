package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"log/slog"
	"net"
	"talaria/api"
	"talaria/protocol"
	"talaria/utils"
)

type Server struct {
	host string
	port string
}

type Client struct {
	conn net.Conn
}

func NewServer() *Server {
	return &Server{
		host: utils.GetEnvVar("BROKER_HOST", "0.0.0.0"),
		port: utils.GetEnvVar("BROKER_PORT", "9092"),
	}
}

func (server *Server) Run() {
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%s", server.host, server.port))
	if err != nil {
		slog.Error("error creating tcp listener", "err", err)
		return
	}
	defer listener.Close()

	slog.Info(fmt.Sprintf("tcp server listening on %s:%s", server.host, server.port))

	for {
		conn, err := listener.Accept()
		if err != nil {
			slog.Error("error accepting tcp connections", "err", err)
		}

		client := &Client{
			conn: conn,
		}

		go client.handleRequest()
	}
}

func (client *Client) handleRequest() {
	defer client.conn.Close()

	for {
		// first 4 bytes contain the message size
		sizeBytes := make([]byte, 4)
		_, err := io.ReadFull(client.conn, sizeBytes[:])
		if err == io.EOF {
			break
		}
		if err != nil {
			slog.Error("tcp read error", "err", err)
			break
		}
		size := binary.BigEndian.Uint32(sizeBytes)

		// read the rest of the message into the buffer.
		messageBytes := make([]byte, size)

		if _, err := io.ReadFull(client.conn, messageBytes[:]); err != nil {
			slog.Error("error decoding message", "err", err)
			break
		}

		// save the message to a file to use for testing later.
		// encoded := hex.EncodeToString(messageBytes)
		// fmt.Println(encoded)

		// Parse the header twice, first time parse only API key and API version, from which we can
		// infer the correct header version and then parse that again in the API code to get the full header.
		header := &protocol.RequestHeader{}
		protocol.VersionedDecode(messageBytes, header, 1)

		slog.Debug(header.String())

		var apiHandler api.API
		switch header.RequestApiKey {
		case (&protocol.ApiVersionsRequest{}).GetKey():
			apiHandler = api.APIVersionsAPI{Request: makeRequest(messageBytes, client.conn, (&protocol.ApiVersionsRequest{Version: header.RequestApiVersion}).GetHeaderVersion())}
		case (&protocol.MetadataRequest{}).GetKey():
			apiHandler = api.MetadataAPI{Request: makeRequest(messageBytes, client.conn, (&protocol.MetadataRequest{Version: header.RequestApiVersion}).GetHeaderVersion())}
			// case protocol.ProduceKey:
			// 	apiHandler = api.ProduceAPI{Request: request}
		}

		err = api.HandleResponse(apiHandler)
		if err != nil {
			slog.Error("error handling response", "err", err)
			break
		}

	}
}

func makeRequest(msg []byte, conn net.Conn, headerVersion int16) api.Request {
	// parse the full header, based on API key and version
	header := &protocol.RequestHeader{}
	protocol.VersionedDecode(msg, header, headerVersion)

	return api.Request{
		Header:  *header,
		Message: msg[header.Len():],
		Conn:    conn,
	}
}
