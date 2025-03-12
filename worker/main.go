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

package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"
	"path/filepath"

	pb "remote-build/remote-build"

	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50052, "The worker port")
)

type worker struct {
	pb.UnimplementedServerWorkerServer
}

func (w *worker) HelloWorker(ctx context.Context, in *pb.WorkRequest) (*pb.WorkResponce, error) {
	log.Printf("Worker received compile request for file: %v", in.GetFilename())
	log.Printf("Compile command: %v", in.GetCompileCommand())

	tempDir, err := os.MkdirTemp("", "remote-build-worker")
	if err != nil {
		log.Printf("Failed to create temp directory: %v", err)
		return &pb.WorkResponce{Message: "Error: Failed to create compilation environment"}, err
	}
	defer os.RemoveAll(tempDir)

	sourcePath := filepath.Join(tempDir, in.GetFilename())
	err = os.WriteFile(sourcePath, []byte(in.GetContent()), 0644)
	if err != nil {
		log.Printf("Failed to write source file: %v", err)
		return &pb.WorkResponce{Message: "Error: Failed to write source file"}, err
	}

	cmd := exec.Command("sh", "-c", in.GetCompileCommand())
	cmd.Dir = tempDir
	output, err := cmd.CombinedOutput()

	log.Printf("Compilation output: %s", string(output))

	if err != nil {
		log.Printf("Compilation failed: %v", err)
		return &pb.WorkResponce{Message: "Compilation failed: " + string(output)}, nil
	}

	return &pb.WorkResponce{Message: "Compilation successful: " + string(output)}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("Worker failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterServerWorkerServer(s, &worker{})
	log.Printf("Worker listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Worker failed to serve: %v", err)
	}
}
