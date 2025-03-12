/*
 *
 * Copyright 2015 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// Package main implements a server for Greeter service.
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	pb "remote-build/remote-build"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const defaultName = "world"

var (
	port       = flag.Int("port", 50051, "The server port")
	workerAddr = flag.String("worker_addr", "localhost:50052", "The worker address")
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedClientServerServer
}

func (s *server) HelloServer(_ context.Context, in *pb.BuildRequest) (*pb.BuildResponse, error) {
	log.Printf("Received build request for file: %v", in.GetFilename())
	log.Printf("Compile command: %v", in.GetCompileCommand())
	//server to worker
	conn, err := grpc.Dial(*workerAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Server did not connect to worker: %v", err)
	}
	defer conn.Close()
	client := pb.NewServerWorkerClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	workRequest := &pb.WorkRequest{
		CompileCommand: in.GetCompileCommand(),
		Filename:       in.GetFilename(),
		Content:        in.GetContent(),
	}
	r, err := client.HelloWorker(ctx, workRequest)
	if err != nil {
		log.Printf("Server call to HelloWorker failed: %v", err)
		return &pb.BuildResponse{Filename: "error_calling_worker"}, err
	}
	log.Printf("Response from worker: %s", r.GetMessage())

	outputFilename := in.GetFilename() + ".out"
	return &pb.BuildResponse{Filename: outputFilename}, nil
}

func main() {
	//client to server
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterClientServerServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
