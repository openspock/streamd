package main

import (
	"fmt"

	log "github.com/openspock/log"
	"github.com/openspock/streamd/conf"
)

func main() {
	log.SysInfo(fmt.Sprintf("streamd init on port: %d", conf.Base.Streamd.Port))
}
