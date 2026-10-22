package util

import (
	"flag"
	"fmt"
)

func PrintParams() {
	flag.VisitAll(func(f *flag.Flag) {
		fmt.Printf("-%s Value: %v\n", f.Name, f.Value.String())
	})
}
