package logs

import (
	"embed"
	"fmt"

	"github.com/dokku/dokku/plugins/common"
)

// MaxSize is the default max retention size for docker logs
const MaxSize = "10m"

var (
	// DefaultProperties is a map of all valid logs properties with corresponding default property values
	DefaultProperties = map[string]string{
		"max-size":    "",
		"vector-sink": "",
	}

	// GlobalProperties is a map of all valid global logs properties
	GlobalProperties = map[string]bool{
		"max-size":     true,
		"vector-image": true,
		"vector-sink":  true,
	}
)

// VectorImage contains the default vector image to run
const VectorImage = "timberio/vector:0.35.X-debian"

// VectorDefaultSink contains the default sink in use for vector log shipping
const VectorDefaultSink = "blackhole://?print_interval_secs=1"

//go:embed templates/*
var templates embed.FS

// GetFailedLogs outputs failed deploy logs for a given app
func GetFailedLogs(appName string) error {
	common.LogInfo2Quiet(fmt.Sprintf("%s failed deploy logs", appName))
	s := common.GetAppScheduler(appName)
	if err := common.PlugnTrigger("scheduler-logs-failed", s, appName); err != nil {
		return err
	}
	return nil
}
