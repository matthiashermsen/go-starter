package main

import (
	"fmt"

	"github.com/matthiashermsen/go-starter/app"
	"github.com/matthiashermsen/go-starter/logging"
)

func main() {
	logging.Info(fmt.Sprintf("Running on version %s", app.Version))
}
