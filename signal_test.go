package signal_test

import (
	"context"
	"errors"
	"syscall"
	"testing"
	"time"

	"github.com/vivangkumar/signal"
)

type app struct{}

func (a app) Run(ctx context.Context) error {
	// Wait on context to be cancelled.
	<-ctx.Done()
	return ctx.Err()
}

func TestWrap(t *testing.T) {
	a := app{}
	ctx := context.Background()

	go func() {
		// Kill the current process after a second.
		<-time.After(1 * time.Second)
		syscall.Kill(syscall.Getpid(), syscall.SIGINT)
	}()

	err := signal.Wrap(ctx, a, syscall.SIGINT)
	if !errors.Is(err, context.Canceled) {
		t.Errorf("expected context cancelled, but got %s", err.Error())
		t.Fail()
	}
}
