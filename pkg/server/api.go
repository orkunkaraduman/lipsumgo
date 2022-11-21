package server

import (
	"context"
	"log"

	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"lipsumgo/pkg/lipsum"
	"lipsumgo/pkg/pb"
)

type Api struct {
	pb.UnimplementedApiServer
}

func (a *Api) Ping(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

func (a *Api) Time(context.Context, *emptypb.Empty) (*timestamppb.Timestamp, error) {
	return timestamppb.Now(), nil
}

func (a *Api) GetSentence(ctx context.Context, req *emptypb.Empty) (rep *pb.ApiGetSentenceReply, err error) {
	sentence, index := lipsum.GetSentence()
	log.Printf("sentence: %q %d", sentence, index)
	return &pb.ApiGetSentenceReply{
		Sentence: sentence,
		Index:    int32(index),
	}, nil
}
