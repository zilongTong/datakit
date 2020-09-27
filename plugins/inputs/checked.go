package inputs

var (
	AllInputs = map[string]bool{
		"active_directory":       true,
		"activemq":               true,
		"aliyunactiontrail":      false,
		"aliyuncdn":              true,
		"aliyuncms":              true,
		"aliyuncost":             true,
		"aliyunddos":             false,
		"aliyunfc":               true,
		"aliyunlog":              true,
		"aliyunobject":           true,
		"aliyunprice":            true,
		"aliyunrdsslowlog":       true,
		"aliyunsecurity":         true,
		"amqp_consumer":          false,
		"ansible":                true,
		"apache":                 true,
		"aspdotnet":              true,
		"awsbill":                true,
		"awscloudtrail":          true,
		"azure_monitor":          false,
		"baiduIndex":             true,
		"binlog":                 true,
		"bitbucket":              true,
		"cassandra":              true,
		"ceph":                   true,
		"clickhouse":             true,
		"cloudflare":             true,
		"cloudwatch":             false,
		"collectd":               true,
		"confluence":             true,
		"consul":                 true,
		"containerd":             true,
		"coredns":                true,
		"cpu":                    true,
		"csv":                    false,
		"dataclean":              false,
		"disk":                   true,
		"diskio":                 true,
		"dns_query":              true,
		"docker":                 true,
		"docker_containers":      true,
		"dockerlog":              true,
		"dotnetclr":              true,
		"druid":                  true,
		"elasticsearch":          true,
		"envoy":                  true,
		"etcd":                   true,
		"exec":                   true,
		"expressjs":              false, // TODO: impl by http-input
		"external":               true,
		"flink":                  true,
		"fluentd":                true,
		"github":                 true,
		"gitlab":                 true,
		"goruntime":              true,
		"hadoop_hdfs":            true,
		"haproxy":                false,
		"harborMonitor":          true,
		"hostobject":             false,
		"http":                   true,
		"http_response":          true,
		"httpjson":               true,
		"httpstat":               true,
		"huaweiyunces":           true,
		"iis":                    true,
		"influxdb":               true,
		"internal":               true,
		"iptables":               true,
		"jboss":                  true,
		"jenkins":                true,
		"jira":                   true,
		"jolokia2_agent":         true,
		"jvm":                    true,
		"kafka":                  true,
		"kafka_consumer":         true,
		"kapacitor":              true,
		"kernel":                 true,
		"kibana":                 true,
		"kong":                   true,
		"kube_inventory":         true,
		"kubernetes":             true,
		"lighttpd":               true,
		"mem":                    true,
		"memcached":              true,
		"mock":                   false,
		"modbus":                 true,
		"mongodb":                true,
		"mongodb_oplog":          true,
		"mqtt_consumer":          true,
		"msexchange":             false,
		"mysqlMonitor":           true,
		"nats":                   true,
		"neo4j":                  true,
		"net":                    true,
		"net_response":           true,
		"netstat":                true,
		"nfsstat":                true,
		"nginx":                  true,
		"nginx_plus":             true,
		"nginx_plus_api":         true,
		"nginx_upstream_check":   true,
		"nginx_vts":              true,
		"nsq":                    true,
		"nsq_consumer":           true,
		"ntpq":                   true,
		"nvidia_smi":             true,
		"openldap":               true,
		"openntpd":               true,
		"oraclemonitor":          true,
		"phpfpm":                 true,
		"ping":                   true,
		"postgresql":             true,
		"postgresql_replication": true,
		"processes":              false,
		"procstat":               true,
		"prom":                   true,
		"puppetagent":            true,
		"rabbitmq":               true,
		"redis":                  true,
		"scanport":               true,
		"self":                   true,
		"snmp":                   true,
		"socket_listener":        false,
		"solr":                   true,
		"sqlserver":              false,
		"squid":                  true,
		"ssh":                    true,
		"statsd":                 true,
		"swap":                   true,
		"syslog":                 true,
		"system":                 true,
		"systemd":                true,
		"systemd_units":          false,
		"tailf":                  true,
		"tcpdump":                false,
		"tencentcms":             true,
		"tencentcost":            true,
		"tencentobject":          true,
		"tengine":                true,
		"tidb":                   true,
		"timezone":               true,
		"traceJaeger":            false,
		"traceSkywalking":        true,
		"traceZipkin":            false,
		"tracerouter":            true,
		"traefik":                true,
		"ucloud_monitor":         false,
		"uwsgi":                  true,
		"varnish":                true,
		"vsphere":                false,
		"weblogic":               true,
		"wechatminiprogram":      true,
		"win_perf_counters":      true,
		"win_services":           true,
		"wmi":                    true,
		"x509_cert":              true,
		"yarn":                   true,
		"zabbix":                 true,
		"zookeeper":              false,
	}
)
