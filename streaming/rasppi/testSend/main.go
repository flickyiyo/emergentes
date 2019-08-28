package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/flickyiyo/emergentes/imgstream"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("35.203.126.106:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect %v", err)
	}

	client := imgstream.NewImgStreamServiceClient(conn)
	buffer := readFile()
	stream, err := client.SentFromRasppi(context.Background())
	go func() {
		for {
			err = stream.Send(&imgstream.ImageStream{
				Image: buffer,
			})
			if err != nil {
				log.Fatalf("Error sendin image %v\n", err)
			}
			time.Sleep(time.Millisecond * 100)
		}
	}()
	waitc := make(chan string)
	stream2, err := client.AskFromMobile(context.Background())
	go func() {
		for {
			msg, err := stream2.Recv()
			if err == io.EOF {
				waitc <- "EOF"
			}
			if err != nil {
				waitc <- "Error"
			}
			fmt.Println(len(msg.GetImage()))
			time.Sleep(time.Second * 1)
		}

	}()
	fmt.Println(<-waitc)
}

func readFile() []byte {
	fileName := "./img.jpg"
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
	return bytes
	// fileType := http.DetectContentType(bytes)

}
