package main

import (
	"fmt"
	"os"

	"goa.design/goa"
	cli "goa.design/goa/examples/error/gen/grpc/cli/divider"
	"google.golang.org/grpc"
)

func doGRPC(scheme, host string, timeout int, debug bool) (goa.Endpoint, interface{}, error) {
	conn, err := grpc.Dial(host, grpc.WithInsecure())
	if err != nil {
		fmt.Fprintln(os.Stderr, fmt.Sprintf("could not connect to gRPC server at %s: %v", host, err))
	}
	return cli.ParseEndpoint(conn)
}