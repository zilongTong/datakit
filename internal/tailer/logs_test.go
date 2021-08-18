package tailer

import (
	"strconv"
	"testing"
	"time"

	tu "gitlab.jiagouyun.com/cloudcare-tools/cliutils/testutil"
)

var (
	ts, _ = time.Parse(time.RFC3339, "2021-06-30T16:25:27Z")

	timeStr = func() string {
		return strconv.Itoa(int(ts.UnixNano()))
	}()
)

func TestLogsAll(t *testing.T) {
	const source = "testing"

	var cases = []struct {
		text string
		res  string
	}{
		{
			"2021-06-30T16:25:27.394+0800    DEBUG   config  config/inputcfg.go:27   ignore dir /usr/local/datakit/conf.d/jvm",
			`testing message="2021-06-30T16:25:27.394+0800    DEBUG   config  config/inputcfg.go:27   ignore dir /usr/local/datakit/conf.d/jvm",status="info" ` + timeStr,
		},
	}

	for _, tc := range cases {
		logs := NewLogs(tc.text)
		logs.ts = ts

		output := logs.Pipeline(nil).
			CheckFieldsLength().
			AddStatus(false).
			IgnoreStatus(nil).
			Point(source, nil).
			Output()

		tu.Assert(t, output == tc.res,
			"\nexpect: %s\n   got: %s",
			tc.res, output)
	}
}
