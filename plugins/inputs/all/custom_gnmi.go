//go:build !custom || inputs || inputs.custom_gnmi

package all

import _ "github.com/influxdata/telegraf/plugins/inputs/custom_gnmi" // register plugin
