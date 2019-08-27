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
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect %v", err)
	}

	client := imgstream.NewImgStreamServiceClient(conn)
	buffer := readFile()
	stream, err := client.SentFromRasppi(context.Background())
	for {
		err = stream.Send(&imgstream.ImageStream{
			Image: buffer,
		})
		if err != nil {
			log.Fatalf("Error sendin image %v\n", err)
		}
		time.Sleep(time.Millisecond * 100)
	}

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
