package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/briheet/kumo/internal/cmd"
)

// Project will be managed by cli for all steps
func main() {

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer cancel()

	ret := cmd.Execute(ctx)
	os.Exit(ret)
}
