package main

import (
	"context"
	"log"
	"net"
	"time"
	b64 "encoding/base64"
	pb "emqx.io/grpc/emqx_schema_registry/protobuf"
	utils "emqx.io/grpc/emqx_schema_registry/utils"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

var cnter *utils.Counter = utils.NewCounter(0, 100)
var last_printed_cnter int64 = 0

// server is used to implement emqx_emqx_schema_registry_v1.s *server
type server struct {
	pb.UnimplementedParserServer
}

// callbacks
func (s *server) Parse(ctx context.Context, request *pb.ParseRequest) (*pb.ParseResponse, error) {
	cnter.Count(1)

	var result []byte
	if request.GetType() == pb.ParseRequest_DECODE {
		sDec, _ := b64.StdEncoding.DecodeString(string(request.GetData()))
		result = []byte(sDec)
	} else if request.GetType() == pb.ParseRequest_ENCODE {
		sEnc := b64.StdEncoding.EncodeToString(request.GetData())
		result = []byte(sEnc)
	}
	//log.Println("parse result: ", string(result))
	return &pb.ParseResponse{
		Code: pb.ParseResponse_SUCCESS,
		Message: "ok",
		Result: result,
	}, nil
}

func (s *server) HealthCheck(ctx context.Context, ping *pb.Ping) (*pb.Pong, error) {
	return &pb.Pong{}, nil
}

func main() {
	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)
	print_counter_periodically(ticker, done)

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterParserServer(s, &server{})
	log.Println("gRPC server is started on port ", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	ticker.Stop()
	done <- true
}

func print_counter_periodically(ticker *time.Ticker, done chan bool) {
	go func() {
		for {
			select {
			case <-done:
				return
			case <-ticker.C:
				maybe_print_counter()
			}
		}
	}()
	return
}

func maybe_print_counter() {
	cnter_now := cnter.GetCount()
	if cnter_now != last_printed_cnter {
		last_printed_cnter = cnter_now
		log.Println("Received: ", cnter_now)
	}
}
