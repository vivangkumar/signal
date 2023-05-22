## signal

Reduce OS signal handling boilerplate for Go programs.

This is mostly written for my own use, really.
Found myself reaching for this bit of code quite often across a few codebases.

## Rationale

This library is intended to be used with long running applications that often rely on OS signals being handled.

It is built on the common Go idiom of propagating context through the function call chain.

## Usage

signal provides a single function that wraps the application with signal handling code, returning a context that is cancelled when a OS signal is sent.

It allows all parts of an application relying on context cancellation to be cancelled together.

It reduces code like this (assuming that a `Run` method is implemented)

```go
ctx, cancel := context.WithCancel(context.Background())
go srv.Run(ctx)

sigint := make(chan os.Signal, 1)
signal.Notify(sigint, os.Interrupt)
<-sigint

cancel()

// and so on..
```

into

```go
ctx := context.Background()
err := signal.Wrap(ctx, srv, os.Interrupt)
if err != nil {
    fmt.Println(err.Error())
}
```

## Linting, formatting & tests

```
make lint
make fmt
make test
```
