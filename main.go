package main

import (
	"fmt"
	"nautilus-print-server/log"
	"nautilus-print-server/response"
	"nautilus-print-server/zpl"
	"net/http"
	"os"
	"time"

	"github.com/olahol/melody"
)

func main() {
	m := melody.New()
	m.Config.PingPeriod = 1 * time.Second
	m.Config.MaxMessageSize = 1024 * 100 // ~100kb

	log_file, err := os.OpenFile(fmt.Sprintf("%s/nautilus-print-server.log", os.TempDir()), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Printf("Couldn't create/open log file: %s", err)
		panic(err)
	}
	defer log_file.Close()
	log.Initialize(log_file)

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		m.HandleRequest(w, r)
	})

	m.HandleMessage(func(s *melody.Session, b []byte) {
		defer func() {
			if r := recover(); r != nil {
				log.Default().Printf("Recovered in f: %s", r)
			}
		}()
		if err := zpl.ExecuteZpl(string(b)); err != nil {
			log.Default().Printf("Error executing zpl: %s because %s", b, err)
			m.Broadcast(
				response.Error(err).ToByte(),
			)
			return
		}

		m.Broadcast(
			response.Success(string(b)).ToByte(),
		)
	})

	m.HandleError(func(s *melody.Session, err error) {
		log.Default().Printf("Error: %s", err)
	})

	if err := http.ListenAndServe(":9000", nil); err != nil {
		log.Default().Fatal(err)
	}
}
