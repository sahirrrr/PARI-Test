package main

import (
	"os"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/sahirrrr/PARI-Test/internal/app"
	"github.com/sahirrrr/PARI-Test/internal/app/infra"
	"github.com/sahirrrr/PARI-Test/internal/app/service/create_new_item"
	"github.com/sahirrrr/PARI-Test/internal/app/service/delete_item"
	"github.com/sahirrrr/PARI-Test/internal/app/service/get_item_by_id"
	"github.com/sahirrrr/PARI-Test/internal/app/service/get_list_items"
	"github.com/sahirrrr/PARI-Test/internal/app/service/update_item"
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
	typapp.Provide("", get_list_items.New)
	typapp.Provide("", get_item_by_id.New)
	typapp.Provide("", create_new_item.New)
	typapp.Provide("", delete_item.New)
	typapp.Provide("", update_item.New)

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
