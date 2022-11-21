package server

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	"lipsumgo/pkg/lipsum"
	"lipsumgo/pkg/pb"
)

type Api struct {
	pb.UnimplementedApiServer
}

func (a *Api) GetSentence(ctx context.Context, req *emptypb.Empty) (rep *pb.ApiGetSentenceReply, err error) {
	sentence, index := lipsum.GetSentence()
	return &pb.ApiGetSentenceReply{
		Sentence: sentence,
		Index:    int32(index),
	}, nil
}
