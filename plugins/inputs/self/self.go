package self

import (
	"os"
	"runtime"
	"time"

	"gitlab.jiagouyun.com/cloudcare-tools/datakit"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit/io"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs"
)

var (
	inputName = "self"
)

type SelfInfo struct {
	stat *ClientStat
}

func (_ *SelfInfo) Catalog() string {
	return "self"
}

func (_ *SelfInfo) SampleConfig() string {
	return ``
}

// func (_ *SelfInfo) Description() string {
// 	return ""
// }

func (s *SelfInfo) Run() {

	tick := time.NewTicker(time.Second * 10)
	defer tick.Stop()

	for {

		select {
		case <-datakit.Exit.Wait():
			return
		case <-tick.C:
			s.stat.Update()
			statMetric := s.stat.ToMetric()

			io.NamedFeed([]byte(statMetric.String()), io.Metric, inputName)
		}
	}
}

func init() {
	StartTime = time.Now()
	inputs.Add(inputName, func() inputs.Input {
		return &SelfInfo{
			stat: &ClientStat{
				OS:   runtime.GOOS,
				Arch: runtime.GOARCH,
				PID:  os.Getpid(),
			},
		}
	})
}
