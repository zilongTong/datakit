package datakit

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"

	bstoml "github.com/BurntSushi/toml"

	"gitlab.jiagouyun.com/cloudcare-tools/cliutils"
)

var (
	IntervalDuration = 10 * time.Second

	Cfg = DefaultConfig()
)

func DefaultConfig() *Config {
	return &Config{ //nolint:dupl
		MainCfg: &MainConfig{
			GlobalTags:      map[string]string{},
			flushInterval:   Duration{Duration: time.Second * 10},
			Interval:        "10s",
			MaxPostInterval: "15s", // add 5s plus for network latency
			StrictMode:      false,

			HTTPBind: "0.0.0.0:9529",

			LogLevel:  "info",
			Log:       filepath.Join(InstallDir, "datakit.log"),
			LogRotate: 32,
			LogUpload: false,
			GinLog:    filepath.Join(InstallDir, "gin.log"),

			//DataWay:   &DataWayCfg{},

			RoundInterval: false,
			TelegrafAgentCfg: &TelegrafCfg{
				Interval:                   "10s",
				RoundInterval:              true,
				MetricBatchSize:            1000,
				MetricBufferLimit:          100000,
				CollectionJitter:           "0s",
				FlushInterval:              "10s",
				FlushJitter:                "0s",
				Precision:                  "ns",
				Debug:                      false,
				Quiet:                      false,
				LogTarget:                  "file",
				Logfile:                    filepath.Join(TelegrafDir, "agent.log"),
				LogfileRotationMaxArchives: 5,
				LogfileRotationMaxSize:     "32MB",
				OmitHostname:               true, // do not append host tag
			},
		},
	}
}

//用于支持在datakit.conf中加入telegraf的agent配置
type TelegrafCfg struct {
	Interval                   string `toml:"interval"`
	RoundInterval              bool   `toml:"round_interval"`
	Precision                  string `toml:"precision"`
	CollectionJitter           string `toml:"collection_jitter"`
	FlushInterval              string `toml:"flush_interval"`
	FlushJitter                string `toml:"flush_jitter"`
	MetricBatchSize            int    `toml:"metric_batch_size"`
	MetricBufferLimit          int    `toml:"metric_buffer_limit"`
	FlushBufferWhenFull        bool   `toml:"-"`
	UTC                        bool   `toml:"utc"`
	Debug                      bool   `toml:"debug"`
	Quiet                      bool   `toml:"quiet"`
	LogTarget                  string `toml:"logtarget"`
	Logfile                    string `toml:"logfile"`
	LogfileRotationInterval    string `toml:"logfile_rotation_interval"`
	LogfileRotationMaxSize     string `toml:"logfile_rotation_max_size"`
	LogfileRotationMaxArchives int    `toml:"logfile_rotation_max_archives"`
	OmitHostname               bool   `toml:"omit_hostname"`
}

type Config struct {
	MainCfg      *MainConfig
	InputFilters []string
}

type DataWayCfg struct {
	URL          string `toml:"url"`
	DataCleanURL string `toml:"dataclean_url"`

	DeprecatedHost   string `toml:"host,omitempty"`
	DeprecatedScheme string `toml:"scheme,omitempty"`
	DeprecatedToken  string `toml:"token,omitempty"`

	host   string
	scheme string
	token  string
	path   string

	urlValues url.Values

	Timeout string `toml:"timeout"`
}

func (dc *DataWayCfg) DeprecatedMetricURL() string {
	return fmt.Sprintf("%s://%s%s?%s",
		dc.scheme,
		dc.host,
		"/v1/write/metrics",
		dc.urlValues.Encode())
}

func (dc *DataWayCfg) MetricURL() string {
	return fmt.Sprintf("%s://%s%s?%s",
		dc.scheme,
		dc.host,
		"/v1/write/metric",
		dc.urlValues.Encode())
}

func (dc *DataWayCfg) ObjectURL() string {
	return fmt.Sprintf("%s://%s%s?%s",
		dc.scheme,
		dc.host,
		"/v1/write/object",
		dc.urlValues.Encode())
}

func (dc *DataWayCfg) LoggingURL() string {
	return fmt.Sprintf("%s://%s%s?%s",
		dc.scheme,
		dc.host,
		"/v1/write/logging",
		dc.urlValues.Encode())
}

func (dc *DataWayCfg) KeyEventURL() string {
	return fmt.Sprintf("%s://%s%s?%s",
		dc.scheme,
		dc.host,
		"/v1/write/keyevent",
		dc.urlValues.Encode())
}
func (dc *DataWayCfg) DataCleanMetricURL() string {
	return fmt.Sprintf("%s://%s%s?%s",
		dc.scheme,
		dc.host,
		dc.path,
		"category=/v1/write/metric")
}
func (dc *DataWayCfg) DataCleanObjectURL() string {
	return fmt.Sprintf("%s://%s%s?%s",
		dc.scheme,
		dc.host,
		dc.path,
		"category=/v1/write/object")
}
func (dc *DataWayCfg) DataCleanLoggingURL() string {
	return fmt.Sprintf("%s://%s%s?%s",
		dc.scheme,
		dc.host,
		dc.path,
		"category=/v1/write/logging")
}

func (dc *DataWayCfg) DataCleanKeyEventURL() string {
	return fmt.Sprintf("%s://%s%s?%s",
		dc.scheme,
		dc.host,
		dc.path,
		"category=/v1/write/keyevent")
}

func (dc *DataWayCfg) Test() error {

	tcphost := dc.host

	if _, _, err := net.SplitHostPort(tcphost); err != nil {
		switch dc.scheme {
		case "http", "ws":
			tcphost += ":80"
		case "https", "wss":
			tcphost += ":443"
		}

		if _, _, err := net.SplitHostPort(tcphost); err != nil {
			l.Errorf("net.SplitHostPort(): %s", err)
			return err
		}
	}

	conn, err := net.DialTimeout("tcp", tcphost, time.Second*5)
	if err != nil {
		l.Errorf("TCP dial host `%s' failed: %s", dc.host, err.Error())
		return err
	}

	if err := conn.Close(); err != nil {
		l.Errorf("close failed: %s", err.Error())
		return err
	}
	return nil
}

func (dc *DataWayCfg) addToken(tkn string) {
	if dc.urlValues == nil {
		dc.urlValues = url.Values{}
	}

	if dc.urlValues.Get("token") == "" {
		l.Debugf("use old token %s", dc.DeprecatedToken)
		dc.urlValues.Set("token", dc.DeprecatedToken)
	}
}

func ParseDataway(urlstr string) (*DataWayCfg, error) {
	dwcfg := &DataWayCfg{
		Timeout: "30s",
	}

	if u, err := url.Parse(urlstr); err == nil {
		dwcfg.scheme = u.Scheme
		dwcfg.urlValues = u.Query()
		dwcfg.host = u.Host
		u.Path = "" // clear any path
		dwcfg.URL = u.String()
	} else {
		l.Errorf("parse url %s failed: %s", urlstr, err.Error())
		return nil, err
	}

	return dwcfg, nil
}

func ParseDatawayDataClean(urlstr string) (*DataWayCfg, error) {
	dwcfg := &DataWayCfg{
		Timeout: "30s",
	}

	if u, err := url.Parse(urlstr); err == nil {
		dwcfg.scheme = u.Scheme
		dwcfg.urlValues = u.Query()
		dwcfg.host = u.Host
		dwcfg.path = u.Path
		dwcfg.DataCleanURL = u.String()
	} else {
		l.Errorf("parse dataclean url %s failed: %s", urlstr, err.Error())
		return nil, err
	}

	return dwcfg, nil
}

type MainConfig struct {
	UUID     string      `toml:"uuid"`
	Name     string      `toml:"name"`
	DataWay  *DataWayCfg `toml:"dataway,omitempty"`
	HTTPBind string      `toml:"http_server_addr"`

	// For old datakit verison conf, there may exist these fields,
	// if these tags missing, TOML will parse error
	DeprecatedFtGateway        string `toml:"ftdataway,omitempty"`
	DeprecatedIntervalDuration int64  `toml:"interval_duration,omitempty"`
	DeprecatedConfigDir        string `toml:"config_dir,omitempty"`
	DeprecatedOmitHostname     bool   `toml:"omit_hostname,omitempty"`

	Log       string `toml:"log"`
	LogLevel  string `toml:"log_level"`
	LogRotate int    `toml:"log_rotate,omitempty"`
	LogUpload bool   `toml:"log_upload"`

	GinLog               string            `toml:"gin_log"`
	MaxPostInterval      string            `toml:"max_post_interval"`
	GlobalTags           map[string]string `toml:"global_tags"`
	RoundInterval        bool
	StrictMode           bool   `toml:"strict_mode,omitempty"`
	Interval             string `toml:"interval"`
	flushInterval        Duration
	OutputFile           string       `toml:"output_file"`
	Hostname             string       `toml:"hostname,omitempty"`
	DefaultEnabledInputs []string     `toml:"default_enabled_inputs"`
	InstallDate          time.Time    `toml:"install_date,omitempty"`
	TelegrafAgentCfg     *TelegrafCfg `toml:"agent"`
}

func InitDirs() {
	for _, dir := range []string{TelegrafDir, DataDir, LuaDir, ConfdDir} {
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			l.Fatalf("create %s failed: %s", dir, err)
		}
	}
}

func (c *Config) LoadMainConfig(p string) error {
	cfgdata, err := ioutil.ReadFile(p)
	if err != nil {
		l.Errorf("read main cfg %s failed: %s", p, err.Error())
		return err
	}

	return c.doLoadMainConfig(cfgdata)
}

func (c *Config) InitCfg(p string) error {

	if c.MainCfg.Hostname == "" {
		c.setHostname()
	}

	if mcdata, err := TomlMarshal(c.MainCfg); err != nil {
		l.Errorf("TomlMarshal(): %s", err.Error())
		return err
	} else {
		if err := ioutil.WriteFile(p, mcdata, 0600); err != nil {
			l.Errorf("error creating %s: %s", p, err)
			return err
		}
	}

	return nil
}

func (c *Config) doLoadMainConfig(cfgdata []byte) error {
	_, err := bstoml.Decode(string(cfgdata), c.MainCfg)
	if err != nil {
		l.Errorf("unmarshal main cfg failed %s", err.Error())
		return err
	}

	if c.MainCfg.TelegrafAgentCfg.LogTarget == "file" && c.MainCfg.TelegrafAgentCfg.Logfile == "" {
		c.MainCfg.TelegrafAgentCfg.Logfile = filepath.Join(InstallDir, "embed", "agent.log")
	}

	if c.MainCfg.OutputFile != "" {
		OutputFile = c.MainCfg.OutputFile
	}

	if c.MainCfg.Hostname == "" {
		c.setHostname()
	}

	if c.MainCfg.DataWay.URL == "" && c.MainCfg.DataWay.DataCleanURL == "" {
		l.Fatal("dataway url or dataclean url not set")
	}

	var dw *DataWayCfg
	if c.MainCfg.DataWay.DataCleanURL != "" {
		dw, err = ParseDatawayDataClean(c.MainCfg.DataWay.DataCleanURL)
		if err != nil {
			return err
		}
	} else {
		dw, err = ParseDataway(c.MainCfg.DataWay.URL)
		if err != nil {
			return err
		}
	}
	c.MainCfg.DataWay = dw

	if c.MainCfg.DataWay.DeprecatedToken != "" { // compatible with old dataway config
		c.MainCfg.DataWay.addToken(c.MainCfg.DataWay.DeprecatedToken)
	}

	if c.MainCfg.MaxPostInterval != "" {
		du, err := time.ParseDuration(c.MainCfg.MaxPostInterval)
		if err != nil {
			l.Warnf("parse %s failed: %s, set default to 15s", c.MainCfg.MaxPostInterval)
			du = time.Second * 15
		}
		MaxLifeCheckInterval = du
	}

	if c.MainCfg.Interval != "" {
		du, err := time.ParseDuration(c.MainCfg.Interval)
		if err != nil {
			l.Warnf("parse %s failed: %s, set default to 10s", c.MainCfg.Interval)
			du = time.Second * 10
		}
		IntervalDuration = du
	}

	c.MainCfg.TelegrafAgentCfg.Debug = strings.EqualFold(strings.ToLower(c.MainCfg.LogLevel), "debug")

	// reset global tags
	for k, v := range c.MainCfg.GlobalTags {
		switch strings.ToLower(v) {
		case `$datakit_hostname`:
			if c.MainCfg.Hostname == "" {
				c.setHostname()
			}

			c.MainCfg.GlobalTags[k] = c.MainCfg.Hostname
			l.Debugf("set global tag %s: %s", k, c.MainCfg.Hostname)

		case `$datakit_ip`:
			c.MainCfg.GlobalTags[k] = "unavailable"

			if ipaddr, err := LocalIP(); err != nil {
				l.Errorf("get local ip failed: %s", err.Error())
			} else {
				l.Debugf("set global tag %s: %s", k, ipaddr)
				c.MainCfg.GlobalTags[k] = ipaddr
			}

		case `$datakit_uuid`, `$datakit_id`:
			c.MainCfg.GlobalTags[k] = c.MainCfg.UUID
			l.Debugf("set global tag %s: %s", k, c.MainCfg.UUID)

		default:
			// pass
		}
	}

	return nil
}

func (c *Config) setHostname() {
	hn, err := os.Hostname()
	if err != nil {
		l.Errorf("get hostname failed: %s", err.Error())
	} else {
		c.MainCfg.Hostname = hn
		l.Infof("set hostname to %s", hn)
	}
}

func (c *Config) EnableDefaultsInputs(inputlist string) {
	elems := strings.Split(inputlist, ",")
	if len(elems) == 0 {
		return
	}

	for _, name := range elems {
		c.MainCfg.DefaultEnabledInputs = append(c.MainCfg.DefaultEnabledInputs, name)
	}
}

func (c *Config) LoadEnvs(mcp string) error {
	if !Docker { // only accept configs from ENV within docker
		return nil
	}

	enableInputs := os.Getenv("ENV_ENABLE_INPUTS")
	if enableInputs != "" {
		c.EnableDefaultsInputs(enableInputs)
	}

	globalTags := os.Getenv("ENV_GLOBAL_TAGS")
	if globalTags != "" {
		c.MainCfg.GlobalTags = ParseGlobalTags(globalTags)
	}

	loglvl := os.Getenv("ENV_LOG_LEVEL")
	if loglvl != "" {
		c.MainCfg.LogLevel = loglvl
	}

	dwcfg := os.Getenv("ENV_DATAWAY")
	if dwcfg != "" {
		dw, err := ParseDataway(dwcfg)
		if err != nil {
			return err
		}

		if err := dw.Test(); err != nil {
			return err
		}

		c.MainCfg.DataWay = dw
	}

	dkhost := os.Getenv("ENV_HOSTNAME")
	if dkhost != "" {
		l.Debugf("set hostname to %s from ENV", dkhost)
		c.MainCfg.Hostname = dkhost
	} else {
		c.setHostname()
	}

	c.MainCfg.Name = os.Getenv("ENV_NAME")

	if fi, err := os.Stat(mcp); err != nil || fi.Size() == 0 { // create the main config
		if c.MainCfg.UUID == "" { // datakit.conf not exit: we have to create new datakit with new UUID
			c.MainCfg.UUID = cliutils.XID("dkid_")
		}

		c.MainCfg.InstallDate = time.Now()

		cfgdata, err := TomlMarshal(c.MainCfg)
		if err != nil {
			l.Errorf("failed to build main cfg %s", err)
			return err
		}

		l.Debugf("generating datakit.conf...")
		if err := ioutil.WriteFile(mcp, cfgdata, os.ModePerm); err != nil {
			l.Error(err)
			return err
		}
	}

	return nil
}

const (
	tagsKVPartsLen = 2
)

func ParseGlobalTags(s string) map[string]string {
	tags := map[string]string{}

	parts := strings.Split(s, ",")
	for _, p := range parts {
		arr := strings.Split(p, "=")
		if len(arr) != tagsKVPartsLen {
			l.Warnf("invalid global tag: %s, ignored", p)
			continue
		}

		tags[arr[0]] = arr[1]
	}

	return tags
}
