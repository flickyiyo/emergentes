package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/dhowden/raspicam"
	"github.com/flickyiyo/emergentes/imgstream"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Grpc Client")
	conn, err := grpc.Dial("192.168.1.66:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect %v", err)
	}
	defer conn.Close()
	client := imgstream.NewImgStreamServiceClient(conn)
	fmt.Printf("Created Client %f\n", client)
	request := &imgstream.ImageRequest{
		W: 213,
		H: 12312,
	}
	res, err := client.SimpleCall(context.Background(), request)
	if err != nil {
		log.Fatalf("Error invoking simple call %v", err)
	}
	fmt.Println(res.GetW())
	sendImagesToServer(client)
}

func readFile() ([]byte, int64) {
	fileName := "./current.jpg"
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("Error on open file %v\n", err)
	}
	defer file.Close()
	fileInfo, _ := file.Stat()
	size := fileInfo.Size()
	fmt.Println(size)
	bytes := make([]byte, size)

	buffer := bufio.NewReader(file)
	_, err = buffer.Read(bytes)
	return bytes, size
	// fileType := http.DetectContentType(bytes)

}

func sendImagesToServer(client imgstream.ImgStreamServiceClient) {
	fmt.Println("Doing client to server")
	stream, err := client.SentFromRasppi(context.Background())
	if err != nil {
		log.Fatalf("Error while sending images %v\n", err)
	}
	// msg, err := res.Recv()
	for {
		takePicture()
		bytes, _ := readFile()
		stream.Send(&imgstream.ImageStream{
			Image: bytes,
		})
	}
}

func takePicture() {
	f, err := os.Create("current.jpg")
	if err != nil {
		log.Fatalf("Error taking picture %v\n", err)
	}
	s := raspicam.NewStill()
	s.BaseStill.Height = 500
	s.BaseStill.Width = 500
	s.BaseStill.Timeout = time.Millisecond * 50
	errCh := make(chan error)
	log.Println("Capturing image...")
	raspicam.Capture(s, f, errCh)
	f.Close()
}
