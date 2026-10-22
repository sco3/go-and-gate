package srv

import (
	"log"
	"time"
)

func (s *Srv) Serve() error {
	log.Printf("serve")
	s.startNanos = time.Now().UnixNano()
	log.Printf("Start: %v", s.startNanos)
	return s.server.Serve(s.listener)
}
