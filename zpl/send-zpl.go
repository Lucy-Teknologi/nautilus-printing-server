package zpl

import (
	"errors"
	"nautilus-print-server/log"
	"os"
	"path/filepath"
	"regexp"
)

var lp_alike = regexp.MustCompile(`^lp\d+`)

func ExecuteZpl(zpl_commands string) error {
	first_lp_device := ""
	filepath.Walk("/dev/usb/", func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		if !lp_alike.MatchString(info.Name()) {
			return nil
		}
		first_lp_device = path
		return nil
	})
	log.Default().Printf("first lp device: %s\n", first_lp_device)
	if first_lp_device == "" {
		log.Default().Panic("no lp device found, restart service and try again")
		return errors.New("no lp device found, restart service and try again")
	}

	file, err := os.OpenFile(first_lp_device, os.O_APPEND|os.O_WRONLY, 0)
	if err != nil {
		return err
	}
	defer file.Close()

	bufferSize := 8

	for len(zpl_commands) > 0 {
		// Determine the size of the next chunk to be written
		chunkSize := len(zpl_commands)
		if chunkSize > bufferSize {
			chunkSize = bufferSize
		}

		// Write the chunk to the printer
		_, err = file.WriteString(zpl_commands[:chunkSize])
		if err != nil {
			return err
		}

		// Remove the written chunk from the remaining ZPL code
		zpl_commands = zpl_commands[chunkSize:]
	}

	return nil

}
