package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/flickyiyo/emergentes/imgstream"
	"google.golang.org/grpc"
)

type server struct {
	Bytes []byte
	Users map[string]chan string
}

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

func (s *server) AskToRasppi(req *imgstream.ImageRequest, stream imgstream.ImgStreamService_AskToRasppiServer) error {
	fmt.Println("Receiving ask to rasppi")
	waitChan := make(chan string)
	go func() {
		for {
			fmt.Println("iterating")
			if s.Bytes == nil {
				waitChan <- "Close"
				break
			}
			err := stream.Send(&imgstream.ImageStream{
				Image: s.Bytes,
			})
			if err == io.EOF {
				waitChan <- "Close"
				return
			}
			if err != nil {
				waitChan <- "Close"
				return
				// log.Fatalf("Error sending image to client %v\n", err)
			}
			time.Sleep(time.Millisecond * 100)
		}
	}()
	<-waitChan
	return nil
}
func (s *server) SentFromRasppi(stream imgstream.ImgStreamService_SentFromRasppiServer) error {
	fmt.Println("Sent From Rasppi invoked")
	for {
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
		// f, err := os.Create("img.jpg")
		if err != nil {
			log.Fatalf("Error while creating new image %v\n", err)
		}
		if imgStream.GetImage() != nil {
			// _, err = f.Write(imgStream.GetImage())
			s.Bytes = imgStream.GetImage()
			createFile(imgStream.GetImage())
		}
	}
}

func createFile(buffer []byte) {
	element, err := os.Create("static/img.png")
	if err != nil {
		log.Fatalf("Error creating serve file %v\n", err)
	}
	defer element.Close()
	element.Write(buffer)
	element.Close()
}
func (s *server) AskFromMobile(stream imgstream.ImgStreamService_AskFromMobileServer) error {
	fmt.Println("Mobile phone asking")
	waitChan := make(chan string)
	// _, err := stream.Recv()
	// if err == io.EOF || err != nil {
	// 	log.Fatalf("Receivied emty or end of ")
	// }
	go func() {
		for {
			fmt.Println("iterating")
			if s.Bytes == nil {
				waitChan <- "Close"
				break
			}
			err := stream.Send(&imgstream.ImageStream{
				Image: s.Bytes,
			})
			if err == io.EOF {
				waitChan <- "Close"
				return
			}
			if err != nil {
				waitChan <- "Close"
				log.Fatalf("Error sending image to client %v\n", err)
			}
			time.Sleep(time.Millisecond * 100)
		}
	}()
	<-waitChan
	return nil
}

func (s *server) CloseStream(ctx context.Context, req *imgstream.CloseStreamRequest) (*imgstream.CloseStreamResponse, error) {
	username := req.GetUsername()
	s.Users[username] <- "Close"
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
	srv := &server{}
	imgstream.RegisterImgStreamServiceServer(s, srv)
	go func() {
		fs := http.FileServer(http.Dir("static"))
		http.Handle("/", fs)

		http.ListenAndServe(":8080", nil)
	}()
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve %v", err)
	}
}
