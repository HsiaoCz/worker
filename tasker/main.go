package main

import (
	"context"
	"time"
)

type response[T any] struct {
	value T
	err   error
}

type Task[T any] struct {
	respch chan response[T]
	ctx    context.Context
}

func (t *Task[T]) Await() (T, error) {
	for {
		select {
		case <-t.ctx.Done():
			var val T
			return val, t.ctx.Err()
		case resp := <-t.respch:
			return resp.value, resp.err
		}
	}
}

type TaskFunc[T any] func(context.Context) (T, error)

func Spawn[T any](t TaskFunc[T]) *Task[T] {
	ctx := context.Background()
	return spawn[T](ctx, t)
}

func spawn[T any](ctx context.Context, t TaskFunc[T]) *Task[T] {
	respch := make(chan response[T], 32)
	go func() {
		val, err := t(ctx)
		respch <- response[T]{
			value: val,
			err:   err,
		}
	}()
	return &Task[T]{
		respch: respch,
		ctx:    ctx,
	}
}

func main() {

}

func FetchData(ctx context.Context) (string, error) {
	time.Sleep(time.Millisecond * 500)
	return "FOO", nil
}
