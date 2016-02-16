package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/STNS/libnss_stns/cmd"
	"github.com/STNS/libnss_stns/config"
	"github.com/STNS/libnss_stns/request"
)

var ConfigFileName = "/etc/stns/libnss_stns.conf"

func main() {
	config, err := cmd.LoadConfig()
	if err == nil {
		if raw := Fetch(config, flag.Arg(0)); raw != "" {
			fmt.Println(raw)
		}
	}
}

func Fetch(config *config.Config, path string) string {

	r, err := request.NewRequest(config, path)

	if err != nil {
		log.Print(err)
		return ""
	}

	result, err := r.GetRaw()
	if err != nil {
		log.Print(err)
		return ""
	}

	return string(result)
}