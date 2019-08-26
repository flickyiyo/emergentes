package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/flickyiyo/emergentes/imgstream"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Grpc Client")
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
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
	readFile(client)
}

func readFile(client imgstream.ImgStreamServiceClient) {
	fileName := "./gopher.jpg"
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("Error on open file %v\n", err)
		return
	}
	defer file.Close()
	fileInfo, _ := file.Stat()
	var size int64 = fileInfo.Size()
	bytes := make([]byte, size)

	buffer := bufio.NewReader(file)
	_, err = buffer.Read(bytes)
	// fileType := http.DetectContentType(bytes)
	sendImagesToServer(client, bytes)
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
