package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"

	"lipsumgo/pkg/pb"
)

func main() {
	var address string
	flag.StringVar(&address, "address", "localhost:9000", "grpc client address")
	flag.Parse()

	log.SetOutput(os.Stdout)

	cc, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("dial error: %v", err)
		return
	}
	//goland:noinspection GoUnhandledErrorResult
	defer cc.Close()

	ctx, ctxCancel := signal.NotifyContext(context.Background(), os.Kill, os.Interrupt, syscall.SIGTERM)
	defer ctxCancel()

	cl := pb.NewApiClient(cc)

	for ctx.Err() == nil {
		rep, err := cl.GetSentence(ctx, &emptypb.Empty{})
		if err != nil {
			log.Printf("get sentence error: %v", err)
		} else {
			log.Printf("sentence: %q %d", rep.Sentence, rep.Index)
		}
		select {
		case <-ctx.Done():
		case <-time.After(time.Second):
		}
	}
}
