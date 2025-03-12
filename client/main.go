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

// Package main implements a client for Greeter service.
package main

import (
	"context"
	"flag"
	"log"
	pb "remote-build/remote-build"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr            = flag.String("addr", "localhost:50051", "the address to connect to")
	filename        = flag.String("filename", "main.c", "the filename to build")
	compile_command = flag.String("command", "gcc", "the build command")
	content         = flag.String("content", "", "file content to compile")
)

func main() {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewClientServerClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	request := &pb.BuildRequest{
		CompileCommand: *compile_command,
		Filename:       *filename,
		Content:        *content,
	}

	// Send the request to the server
	response, err := c.HelloServer(ctx, request)
	if err != nil {
		log.Fatalf("Build request failed: %v", err)
	}

	log.Printf("Build response received: filename=%s", response.GetFilename())
}
