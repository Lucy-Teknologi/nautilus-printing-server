package main

import (
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

	log_file, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
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
		log.Default().Println(printable)

		// if printable.Type == zpl.Retouching {
		// 	zpl_string := zpl.Retou

		// 	m.Broadcast(
		// 		response.Success("In Retouching").ToByte(),
		// 	)
		// 	return
		// }
		if printable.Type == zpl.Cutting {
			zpl_string := zpl.CuttingZPLString(printable)
			if err := zpl.ExecuteZpl(zpl_string); err != nil {
				log.Default().Printf("Error executing zpl: %s", err)
				m.Broadcast(
					response.Error(err).ToByte(),
				)
				return
			}
			m.Broadcast(
				response.Success(string(b)).ToByte(),
			)
			return
		}

		m.Broadcast(response.ErrorWithMessage("Command not found").ToByte())
	})

	if err := http.ListenAndServe(":9000", nil); err != nil {
		log.Default().Fatal(err)
	}
}
