package main

import (
	"log"
	"os"
	"strconv"
	pb "tsdb"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
	id          = 1234
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewMetricsClient(conn)

	// Contact the server and print out its response.
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	id, _ := strconv.ParseInt(os.Args[2], 10, 64)

	r, err := c.SendMetric(context.Background(), &pb.MetricRequest{Name: name, Id: id})
	if err != nil {
		log.Fatalf("could not post metric: %v", err)
	}
	log.Printf("Metric: %s", r.Message)
}
