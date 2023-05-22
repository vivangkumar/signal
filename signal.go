// Package signal provides context cancellation by handling OS signals.
//
// It provides a minimal interface to handle commonly occurring repeated
// go code when building services.
//
// It uses the same package name as the go std lib in an effort to ensure
// that callers don't have to rely on it to implement OS signal handling.
package signal

import (
	"context"
	"fmt"
	"os"
	"os/signal"
)

// Runner represents a runnable entity.
type Runner interface {
	// Run runs the entity with the provided context.
	Run(ctx context.Context) error
}

// Wrap provides context cancellation by gracefully handling OS signals.
//
// It accepts any type that implements the Runner interface.
//
// Typically, this should be used in conjunction with a long running
// process that relies on contexts for cancellation, ensuring that the
// chain of function calls that propagate the same context are also
// cancelled.
func Wrap(ctx context.Context, r Runner, sig ...os.Signal) error {
	err := r.Run(signalCtx(ctx, sig...))
	if err != nil {
		return fmt.Errorf("signal: %w", err)
	}

	return nil
}

// signalCtx returns a context that is cancelled when encountering an OS signal.
func signalCtx(ctx context.Context, sig ...os.Signal) context.Context {
	ctx, cancel := context.WithCancel(ctx)

	go func() {
		c := make(chan os.Signal, len(sig))

		signal.Notify(c, sig...)
		defer signal.Stop(c)

		select {
		case <-ctx.Done():
		case <-c:
			cancel()
		}
	}()

	return ctx
}
