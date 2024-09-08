package main

import (
	"context"
	"fmt"
	"log"
	"logger-service/data"
	"logger-service/logs"
	"net"

	"google.golang.org/grpc"
)

type LogServer struct {
	logs.UnimplementedLogServiceServer
	Models data.Models
}

func (l *LogServer) WriteLog(ctx context.Context, req *logs.LogRequest) (*logs.LogResponse, error) {
	input := req.GetLogEntry()

	logEntry := data.LogEntry{
		Name: input.Name,
		Data: input.Data,
	}

	err := l.Models.LogEntry.Insert(logEntry)
	if err != nil {
		resp := &logs.LogResponse{
			Result: "failed",
		}

		return resp, err
	}

	resp := &logs.LogResponse{
		Result: "logged!",
	}

	return resp, nil
}

func (app *Config) grpcListen() {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", gRpcPort))
	if err != nil {
		log.Fatalf("Failed to create listener for gRPC: %v", err)
	}

	server := grpc.NewServer()

	logs.RegisterLogServiceServer(server, &LogServer{Models: app.Models})

	log.Printf("gRPC server started on port %s", gRpcPort)

	if err := server.Serve(listener); err != nil {
		log.Fatalf("Failed to listen for gRPC: %v", err)
	}
}
