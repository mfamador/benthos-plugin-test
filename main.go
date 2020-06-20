// main package for test application
package main

import (
	"github.com/Jeffail/benthos/v3/lib/log"
	"github.com/Jeffail/benthos/v3/lib/metrics"
	"github.com/Jeffail/benthos/v3/lib/processor"
	"github.com/Jeffail/benthos/v3/lib/service"
	"github.com/Jeffail/benthos/v3/lib/types"
	"github.com/mfamador/benthos-plugins/lib/sarcasm"
)

func main() {
	processor.RegisterPlugin(
		"how_sarcastic",
		func() interface{} {
			s := sarcasm.SarcasmProc{}
			return &s
		},
		func(
			iconf interface{},
			mgr types.Manager,
			logger log.Modular,
			stats metrics.Type,
		) (types.Processor, error) {
			return iconf.(*sarcasm.SarcasmProc), nil
		},
	)
	service.Run()
}
