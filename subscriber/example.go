package subscriber

import (
	"context"
	"github.com/hxlb/user/proto/example"
	"github.com/micro/go-log"
)

type Example struct{}

func (e *Example) Handle(ctx context.Context, msg *com_hxlb_srv_user.Message) error {
	log.Log("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *com_hxlb_srv_user.Message) error {
	log.Log("Function Received message: ", msg.Say)
	return nil
}
