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
	takePicture(client)
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

func sendImagesToServer(client imgstream.ImgStreamServiceClient, buffer []byte) {
	fmt.Println("Doing client to server")
	stream, err := client.SentFromRasppi(context.Background())
	if err != nil {
		log.Fatalf("Error while sending images %v\n", err)
	}
	// msg, err := res.Recv()
	for {
		stream.Send(&imgstream.ImageStream{
			Image: buffer,
		})
		time.Sleep(time.Second * 2)
	}
}

func takePicture(client imgstream.ImgStreamServiceClient) {
	for {

		f, err := os.Create("current.jpg")
		if err != nil {
			log.Fatalf("Error taking picture %v\n", err)
		}
		s := raspicam.NewStill()
		s.BaseStill.Timeout = time.Millisecond * 200
		errCh := make(chan error)
		log.Println("Capturing image...")
		raspicam.Capture(s, f, errCh)
		bytes, _ := readFile()
		fmt.Println(bytes)
		sendImagesToServer(client, bytes)
	}
}
