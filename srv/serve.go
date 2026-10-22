package srv

import (
	"log"
	"time"
)

func (s *Srv) Serve() error {
	log.Printf("serve")
	s.startNanos = time.Now().UnixNano()
	return s.server.Serve(s.listener)
}
