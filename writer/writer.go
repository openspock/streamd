package writer

import (
	"crypto/tls"
	"strconv"

	log "github.com/openspock/log"
	"github.com/openspock/streamd/conf"
)

// Start stats a writer
func Start() error {
	log.SysInfo("streamd writer server init")

	cer, err := tls.LoadX509KeyPair("server.crt", "server.key")
	if err != nil {
		log.SysError(err)
		return err
	}

	cfg := &tls.Config{Certificates: []tls.Certificate{cer}}

	listener, err := tls.Listen("tcp", ":"+strconv.FormatInt(conf.Base.Streamd.Port, 10), cfg)
	if err != nil {
		log.SysError(err)
		return err
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.SysError(err)
			return err
		}

	}

	return nil
}
