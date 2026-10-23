package srv

import (
	"log"
	"time"
)

func (s *Srv) Serve() error {

	err := WaitStartupBarrier(s.redis, s.host)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("serve\n")
	s.startNanos = time.Now().UnixNano()
	log.Printf("Start: %v", s.startNanos)
	return s.server.Serve(s.listener)
}
