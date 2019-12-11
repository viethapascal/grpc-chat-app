package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"grpc-chat-app/proto"
	"log"
	"net"
	"os"
	"sync"
)

var grpcLog grpclog.LoggerV2

type Connection struct {
	stream proto.Broadcast_CreateStreamServer
	id string
	active bool
	err chan error
}

type Server struct {
	Connection []*Connection
}

func init() {
	grpcLog = grpclog.NewLoggerV2(os.Stdout, os.Stdout, os.Stdout)
}

func (s *Server) CreateStream(pconn *proto.Connect, stream proto.Broadcast_CreateStreamServer) error {
	conn := &Connection{
		stream: stream,
		id: pconn.User.Id,
		active: true,
		err : make(chan error),
	}
	log.Println("Established new connection from ", conn.id)
	s.Connection = append(s.Connection, conn)

	return <-conn.err
}

func (s *Server) BroadcastMessage(ctx context.Context, msg *proto.Message) (*proto.Close, error){
	wait := sync.WaitGroup{}
	done := make(chan int)


	for _, conn := range s.Connection {
		wait.Add(1)

		go func(msg *proto.Message, conn *Connection) {
			defer wait.Done()

			if conn.active {
				err := conn.stream.Send(msg)
				grpcLog.Infof("Sending message %v to user %v", msg.Id, conn.id)

				if err != nil {
					grpcLog.Errorf("Error with stream %v. Error: %v",  conn.stream, err)
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
	return &proto.Close{}, nil
}

func main() {
	var connections []*Connection

	server := &Server{connections}

	grpcServer := grpc.NewServer()
	listener, err := net.Listen("tcp", ":17100")

	if err != nil {
		log.Fatal("Error creating server:", err)
	}

	grpcLog.Info("Starting server at port 17100")

	proto.RegisterBroadcastServer(grpcServer, server)
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal(err)
	}
}

