package main

import (
	"errors"
	"github.com/ehsandavari/go-context-plus"
	"github.com/ehsandavari/go-tracer"
	"github.com/google/uuid"
	"time"
)

func main() {

	ctx := contextplus.Background()
	ctx.SetValue("test", "test_value")
	ctx.SetRequestId(uuid.New().String())
	ctx.SetTraceId(uuid.New().String())
	ctx.User.SetId(uuid.New())
	ctx.User.SetPhoneNumber("989215580690")

	newTracer := tracer.NewTracer(
		true,
		true,
		false,
		"localhost",
		"4317",
		12,
		"test",
		"sConfig.Service.Namespace",
		"sConfig.Service.InstanceId",
		"sConfig.Service.Version",
		"sConfig.Service.Mode",
		"sConfig.Service.CommitId",
	)

	span := newTracer.Tracer("asdasdsd")
	ctx = span.Start(ctx, "create")
	span.SetString("1", "1")
	span.SetString("2", "2")
	span.SetStatus(tracer.Error, "asdfasdfasdfadsfasdfasdfasdfadsf")
	span.RecordError(errors.New("asdfasdfasdfadsf"))
	span.End()

	time.Sleep(10 * time.Second)
}
