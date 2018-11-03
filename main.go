package main

import (
	"fmt"

	cfg "github.com/openspock/conf"
	log "github.com/openspock/log"
)

func main() {
	log.SysInfo(fmt.Sprintf("streamd init on port: %d", cfg.Base.Streamd.Port))
}
