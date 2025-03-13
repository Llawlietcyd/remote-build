package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	pb "remote-build/remote-build"
	"sync"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	port        = flag.Int("port", 50051, "The server port")
	workerAddrs = flag.String("worker_addrs", "localhost:50052,localhost:50053,localhost:50054", "Comma-separated list of worker addresses")
)

type server struct {
	pb.UnimplementedClientServerServer
	workerPool *WorkerPool
}

type WorkerPool struct {
	workers []string
	mu      sync.Mutex
	nextIdx int
}

func NewWorkerPool(workerAddrs []string) *WorkerPool {
	return &WorkerPool{
		workers: workerAddrs,
		nextIdx: 0,
	}
}

func (wp *WorkerPool) GetNextWorker() string {
	wp.mu.Lock()
	defer wp.mu.Unlock()

	worker := wp.workers[wp.nextIdx]
	wp.nextIdx = (wp.nextIdx + 1) % len(wp.workers)
	return worker
}

func SendToWorker(workerAddr string, workRequest *pb.WorkRequest) (*pb.WorkResponce, error) {
	conn, err := grpc.Dial(workerAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("Failed to connect to worker %s: %v", workerAddr, err)
		return nil, err
	}
	defer conn.Close()
	client := pb.NewServerWorkerClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	return client.HelloWorker(ctx, workRequest)
}

func (s *server) HelloServer(_ context.Context, in *pb.BuildRequest) (*pb.BuildResponse, error) {
	log.Printf("Received build request for file: %v", in.GetFilename())

	workRequest := &pb.WorkRequest{
		CompileCommand: in.GetCompileCommand(),
		Filename:       in.GetFilename(),
		Content:        in.GetContent(),
	}

	workerAddr := s.workerPool.GetNextWorker()
	log.Printf("Assigning work to worker at: %s", workerAddr)

	resp, err := SendToWorker(workerAddr, workRequest)
	if err != nil {
		log.Printf("Error communicating with worker %s: %v", workerAddr, err)
		return &pb.BuildResponse{Filename: "error_calling_worker"}, err
	}

	log.Printf("Response from worker: %s", resp.GetMessage())
	outputFilename := in.GetFilename() + ".out"
	return &pb.BuildResponse{Filename: outputFilename}, nil
}

func main() {
	flag.Parse()

	workerAddrList := parseWorkerAddrs(*workerAddrs)
	log.Printf("Starting server with workers: %v", workerAddrList)

	workerPool := NewWorkerPool(workerAddrList)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterClientServerServer(s, &server{workerPool: workerPool})

	log.Printf("Server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func parseWorkerAddrs(addrs string) []string {
	workersList := make([]string, 0)

	currentAddr := ""
	for i := 0; i < len(addrs); i++ {
		if addrs[i] == ',' {
			if currentAddr != "" {
				workersList = append(workersList, currentAddr)
				currentAddr = ""
			}
		} else {
			currentAddr += string(addrs[i])
		}
	}

	if currentAddr != "" {
		workersList = append(workersList, currentAddr)
	}

	return workersList
}
