package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	"path/filepath"
	"strconv"

	pb "tsdb"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type server struct{}

func (s *server) SendMetric(ctx context.Context, in *pb.MetricRequest) (*pb.MetricResponse, error) {
	return &pb.MetricResponse{Message: in.Name + strconv.FormatInt(in.Id, 10)}, nil
}

func main() {
	ScanAllFiles("/Users/nikhil/projects/learn-go/src/tsdb")

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterMetricsServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func readFiles(dirname string) {
	fmt.Println(dirname)
	f, _ := ioutil.ReadDir(dirname)
	for _, file := range f {
		if file.IsDir() {
			readFiles(file.Name())
		} else {
			fmt.Println(file.Name())
		}
	}
}

func ScanAllFiles(dirname string) (err error) {
	numScanned := 0
	var scan = func(path string, fileInfo os.FileInfo, inpErr error) (err error) {
		numScanned++
		fmt.Println(path + fileInfo.Name())
		return nil
	}

	fmt.Println("Scan All")
	err = filepath.Walk(dirname, scan)
	fmt.Println("Total scanned", numScanned)
	return
}
