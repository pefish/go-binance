package main

import (
	"context"
	"log"
	"os"

	"github.com/davecgh/go-spew/spew"
	"github.com/joho/godotenv"
	ws_service "github.com/pefish/go-binance/announcement/ws-service"
	i_logger "github.com/pefish/go-interface/i-logger"
	t_logger "github.com/pefish/go-interface/t-logger"
	go_logger "github.com/pefish/go-logger"
)

var logger i_logger.ILogger = &i_logger.DefaultLogger

func main() {
	envMap, _ := godotenv.Read("./.env")
	for k, v := range envMap {
		os.Setenv(k, v)
	}

	err := do()
	if err != nil {
		log.Fatal(err)
	}
}

func do() error {
	logger.InfoDump(
		os.Getenv("API_KEY"),
		os.Getenv("API_SECRET"),
	)
	wsService := ws_service.New(
		os.Getenv("API_KEY"),
		os.Getenv("API_SECRET"),
		go_logger.NewLogger(t_logger.Level_DEBUG),
	)
	wsService.WatchAnnouncement(
		context.Background(),
		func(event *ws_service.WsAnnouncementEvent) {
			spew.Dump(event)
		},
	)
	return nil
}
