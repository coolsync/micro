package handler

import (
	"context"

	log "github.com/micro/micro/v3/service/logger"

	bj38 "bj38/proto"
)

type Bj38 struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Bj38) Call(ctx context.Context, req *bj38.Request, rsp *bj38.Response) error {
	log.Info("Received Bj38.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Bj38) Stream(ctx context.Context, req *bj38.StreamingRequest, stream bj38.Bj38_StreamStream) error {
	log.Infof("Received Bj38.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Infof("Responding: %d", i)
		if err := stream.Send(&bj38.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Bj38) PingPong(ctx context.Context, stream bj38.Bj38_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&bj38.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
