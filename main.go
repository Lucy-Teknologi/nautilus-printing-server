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

	// get user's home directory
	home_dir := os.Getenv("HOME")
	if home_dir == "" {
		log.Default().Fatal("HOME env variable not set")
	}

	log_file, err := os.OpenFile(fmt.Sprintf("%s/nautilus-print-server.log", home_dir), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Default().Fatal(err)
	}
	defer log_file.Close()
	log.Initialize(log_file)

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		m.HandleRequest(w, r)
	})

	m.HandleMessage(func(s *melody.Session, b []byte) {
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

	if err := http.ListenAndServe(":9000", nil); err != nil {
		log.Default().Fatal(err)
	}
}
