package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"

	"github.com/flickyiyo/emergentes/imgstream"
	"google.golang.org/grpc"
)

type server struct{}

type sendToClient func([]byte) error

type streamStatus int

const (
	connected streamStatus = iota
	disconnected
	cancelled
)

type clientCamera struct {
	Callback sendToClient
	Status   streamStatus
}

var connections map[string]*clientCamera

func (*server) AskToRasppi(req *imgstream.ImageRequest, stream imgstream.ImgStreamService_AskToRasppiServer) error {
	var username string = req.GetUsername()
	cc := &clientCamera{
		Callback: func(imgBytes []byte) error {
			err := stream.Send(&imgstream.ImageStream{
				Image: imgBytes,
			})
			if err != nil {
				log.Fatalf("Error streaming to end client %v\n", err)
				return err
			}
			return nil
		},
		Status: connected,
	}
	connections[username] = cc
	for connections[username].Status == connected {
		time.Sleep(time.Second * 2)
	}
	delete(connections, username)
	return nil
}
func (*server) SentFromRasppi(stream imgstream.ImgStreamService_SentFromRasppiServer) error {
	for {
		fmt.Println("Sent From Rasppi invoked")
		imgStream, err := stream.Recv()
		if err == io.EOF {
			log.Printf("Stream finished %v\n", err)
			return err
		}
		if err != nil {
			// log.Fatalf("Error receiving from Rasp %v\n", err)
			return err
		}
		// log.Println(imgStream.GetImage())
		f, err := os.Create("img.jpg")

		if err != nil {
			log.Fatalf("Error while creating new image %v\n", err)
		}
		if imgStream.GetImage() != nil {
			_, err = f.Write(imgStream.GetImage())
		}
	}
}
func (*server) AskFromMobile(stream imgstream.ImgStreamService_AskFromMobileServer) error {
	return nil
}

func (*server) CloseStream(ctx context.Context, req *imgstream.CloseStreamRequest) (*imgstream.CloseStreamResponse, error) {
	username := req.GetUsername()
	connections[username].Status = cancelled
	log.Printf("Canelled connection to %s", username)
	return nil, nil
}

func (*server) SimpleCall(ctx context.Context, req *imgstream.ImageRequest) (*imgstream.ImageRequest, error) {
	fmt.Println("Simple call invoked")
	width := req.GetW()
	height := req.GetH()
	response := &imgstream.ImageRequest{
		W: width, H: height,
	}
	return response, nil
}

func main() {
	fmt.Println("Image streaming server")
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Error on creating listener %v", err)
	}
	s := grpc.NewServer()
	imgstream.RegisterImgStreamServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve %v", err)
	}
}
