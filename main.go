package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/grpclog"
	gen_pb "grpc-chat-app/go-pb"
	"log"
	"net"
	"net/http"
	"os"
	"sync"
)

var grpcLog grpclog.LoggerV2

type Connection struct {
	stream gen_pb.Broadcast_CreateStreamServer
	id     string
	active bool
	err    chan error
}

type Server struct {
	gen_pb.UnimplementedBroadcastServer
	Connection []*Connection
}

func init() {
	grpcLog = grpclog.NewLoggerV2(os.Stdout, os.Stdout, os.Stdout)
}

func (s *Server) CreateStream(pconn *gen_pb.Connect, stream gen_pb.Broadcast_CreateStreamServer) error {
	conn := &Connection{
		stream: stream,
		id:     pconn.User.Id,
		active: true,
		err:    make(chan error),
	}
	log.Println("Established new connection from ", conn.id)
	s.Connection = append(s.Connection, conn)

	return <-conn.err
}

func (s *Server) BroadcastMessage(ctx context.Context, msg *gen_pb.Message) (*gen_pb.Close, error) {
	wait := sync.WaitGroup{}
	done := make(chan int)

	for _, conn := range s.Connection {
		wait.Add(1)

		go func(msg *gen_pb.Message, conn *Connection) {
			defer wait.Done()

			if conn.active {
				err := conn.stream.Send(msg)
				grpcLog.Infof("Sending message %v to user %v", msg.Id, conn.id)

				if err != nil {
					grpcLog.Errorf("Error with stream %v. Error: %v", conn.stream, err)
					conn.active = false
					conn.err <- err
				}
			}
		}(msg, conn)
	}

	go func() {
		wait.Wait()
		close(done)
	}()

	<-done
	return &gen_pb.Close{}, nil
}
func (s *Server) Echo(ctx context.Context, in *gen_pb.EchoRequest) (*gen_pb.EchoReply, error) {
	return &gen_pb.EchoReply{Message: "Hello" + in.Name}, nil
}

func allowedOrigin(origin string) bool {
	return true
}

func cors(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if allowedOrigin(r.Header.Get("Origin")) {
			w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization, ResponseType")
		}
		if r.Method == "OPTIONS" {
			return
		}
		h.ServeHTTP(w, r)
	})
}
func main() {
	var connections []*Connection

	server := &Server{Connection: connections}
	var port int
	flag.IntVar(&port, "port", 8080, "")
	flag.Parse()
	log.Println(port)

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	grpcServer := grpc.NewServer()
	gen_pb.RegisterBroadcastServer(grpcServer, server)
	go func() {
		log.Fatalln(grpcServer.Serve(listener))
	}()
	log.Println("Serving gRPC on 0.0.0.0:8080")

	// Create a client connection to the gRPC server we just started
	// This is where the gRPC-Gateway proxies the requests
	conn, err := grpc.DialContext(
		context.Background(),
		"0.0.0.0:8080",
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}

	gwmux := runtime.NewServeMux()
	// Register Greeter
	err = gen_pb.RegisterBroadcastHandler(context.Background(), gwmux, conn)
	if err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}

	gwServer := &http.Server{
		Addr:    ":8090",
		Handler: cors(gwmux),
	}

	log.Println("Serving gRPC-Gateway on http://0.0.0.0:8090")
	log.Fatalln(gwServer.ListenAndServe())
}
