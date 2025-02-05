package main

import (
	"os"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/sahirrrr/PARI-Test/internal/app"
	"github.com/sahirrrr/PARI-Test/internal/app/infra"
	"github.com/sirupsen/logrus"
	"github.com/typical-go/typical-go/pkg/typapp"
)

func main() {
	godotenv.Load()

	// Configs
	typapp.Provide("", infra.LoadDatabaseCfg)
	typapp.Provide("", infra.LoadEchoCfg)

	// Infras
	typapp.Provide("", infra.NewDatabases)
	typapp.Provide("", infra.NewEcho)

	// Services

	exitSigns := []os.Signal{
		os.Interrupt,
		os.Kill,
		syscall.SIGINT,
		syscall.SIGKILL,
		syscall.SIGTERM,
		syscall.SIGHUP,
		syscall.SIGQUIT,
	}

	// Start App
	if err := typapp.StartApp(app.Start, app.Shutdown, exitSigns...); err != nil {
		logrus.Fatal(err.Error())
	}
}
