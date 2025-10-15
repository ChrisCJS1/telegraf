package custom_gnmi

import (
	"context"
	_ "embed"
	"sync"

	"github.com/influxdata/telegraf"

	// gNMI specific imports

	"google.golang.org/grpc"
)

//go:embed sample.conf
var sampleConfig string

// CustomGNMI is the Telegraf input plugin struct
type CustomGNMI struct {
	Addresses          []string          `toml:"addresses"`
	Username           string            `toml:"username"`
	Password           string            `toml:"password"`
	TLSCA              string            `toml:"tls_ca"`
	TLSCert            string            `toml:"tls_cert"`
	TLSKey             string            `toml:"tls_key"`
	InsecureSkipVerify bool              `toml:"insecure_skip_verify"`
	Subscriptions      []Subscription    `toml:"subscriptions"`
	SampleInterval     telegraf.Duration `toml:"sample_interval"` // For stream mode

	// Internal fields for gNMI client
	client    gnmi.gNMIClient
	conn      *grpc.ClientConn
	cancelCtx context.CancelFunc
	wg        sync.WaitGroup
	Log       telegraf.Logger
}

// Subscription defines a single gNMI subscription path
type Subscription struct {
	Path   string `toml:"path"`
	Mode   string `toml:"mode"` // "stream", "once", "poll"
	Origin string `toml:"origin"`
}

func (c *CustomGNMI) Description() string {
	return "Input plugin for gNMI telemetry with custom path parsing."
}

func (c *CustomGNMI) SampleConfig() string {
	return sampleConfig
}
