package srv

import "log"

func (s *Srv) Serve() error {
	log.Printf("serve")
	return s.server.Serve(s.listener)
}
