package main

import (
	"io/ioutil"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	log.Info().Msg("Main function starting")
	Print()
}

func load() []byte {
	data, err := ioutil.ReadFile("/tmp/myhotel/info.txt")
	if err != nil {
		log.Error().Err(err).Msg("Read file error")
	}
	return data
}

func Print() {
	log.Info().Msg(string(load()))
}
