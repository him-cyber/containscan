package main

import (
	"fmt"

	"github.com/himaneesh/containscan/logging"
)

func main() {
	logging.Info("ContainScan started successfully!")
	logging.Warning("Potential issue detected")
	logging.Error(nil)
	logging.Error(fmt.Errorf("critical failure occurred"))
}
