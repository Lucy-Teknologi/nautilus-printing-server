package main

import (
	"fmt"
	"nautilus-print-server/log"
	"nautilus-print-server/response"
	"nautilus-print-server/zpl"
	"net/http"
	"os"
	"time"

	jsoniter "github.com/json-iterator/go"
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
		var printable zpl.Printable
		if err := jsoniter.Unmarshal(b, &printable); err != nil {
			log.Default().Printf("Error unmarshalling: %s", err)
			m.Broadcast([]byte("ERROR UNMARSHALLING"))
			return
		}

		if printable.Type == zpl.Cutting {
			zpl_string := zpl.CuttingZPLString(printable)
			log.Default().Println(zpl_string)
			if err := zpl.ExecuteZpl(zpl_string); err != nil {
				log.Default().Printf("Error executing zpl: %s", err)
				m.Broadcast(
					response.Error(err).ToByte(),
				)
				return
			}
			m.Broadcast(
				response.Success(printable).ToByte(),
			)
			return
		}

		m.Broadcast(response.ErrorWithMessage("Command not found").ToByte())
	})

	if err := http.ListenAndServe(":9000", nil); err != nil {
		log.Default().Fatal(err)
	}
}
