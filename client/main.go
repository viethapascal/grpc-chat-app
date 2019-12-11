package main

import (
	"bufio"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"grpc-chat-app/proto"
	"log"
	"os"
	"sync"
	"time"
)

var client proto.BroadcastClient
var wait *sync.WaitGroup

func init() {
	wait = &sync.WaitGroup{}
}

func connect(user *proto.User) (error) {
	var streamError error
	fmt.Println(*user)
	stream, err := client.CreateStream(context.Background(), &proto.Connect{
		User: user,
		Active: true,
	})

	if err != nil {
		return fmt.Errorf("Connect failed: %v", err)
	}

	wait.Add(1)

	go func(str proto.Broadcast_CreateStreamClient) {
		defer wait.Done()

		for  {
			msg, err := str.Recv()

			if err != nil {
				streamError = fmt.Errorf("Error reading message: %v", err)
				break
			}

			fmt.Printf("%v : %s\n", msg.User.DisplayName, msg.Message)
		}
	}(stream)

	return streamError
}

func main() {
	ts := time.Now()

	done := make(chan int)

	name := flag.String("N", "Anonymous", "")
	flag.Parse()

	id := sha256.Sum256([]byte(ts.String() + *name))
	user := &proto.User{
		Id : hex.EncodeToString(id[:]),
		DisplayName: *name,
	}

	conn, err := grpc.Dial("localhost:17100", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect to server %v", err)
	}

	client = proto.NewBroadcastClient(conn)

	connect(user)
	if err != nil {
		log.Fatal(err)
	}

	wait.Add(1)
	go func() {
		defer wait.Done()
		scanner := bufio.NewScanner(os.Stdin)
		ts := time.Now()
		msgID := sha256.Sum256([]byte(ts.String() + *name))
		for scanner.Scan() {
			msg := &proto.Message{
				Id: hex.EncodeToString(msgID[:]),
				User: user,
				Message: scanner.Text(),
				Timestamp: ts.String(),
			}

			_, err := client.BroadcastMessage(context.Background(), msg)
			if err != nil {
				fmt.Printf("Error sending message: %v", err)
				break
			}
		}
	}()

	go func() {
		wait.Wait()
		close(done)
	}()

	<-done
}

