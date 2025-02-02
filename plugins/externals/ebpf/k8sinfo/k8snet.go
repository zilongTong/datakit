package k8sinfo

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type K8sNetInfo struct {
	sync.RWMutex
	cli         *K8sClient
	svcNetInfo  map[string]*K8sServicesNet
	poNetInfoIP map[string]*K8sPodNet
	// 使用主机网络的 pod 		需要判断端口
	poNetInfoPort map[string]map[Port]*K8sPodNet

	autoUpdate bool
}

func (kinfo *K8sNetInfo) Update() error {
	k8sSvcMap := map[string]*K8sServicesNet{}
	k8sPodMap := map[string]*K8sPodNet{}
	k8sPodNetPortMap := map[string]map[Port]*K8sPodNet{} // pod (include host_network)
	k8sPodTmpNetMap := map[string][]*K8sPodNet{}

	ns, err := kinfo.cli.GetNamespaces()
	if err != nil {
		return err
	}

	for _, ns := range ns {
		svcNet, err := kinfo.cli.GetServiceNet(ns)
		if err != nil {
			return err
		}

		endpointNet, err := kinfo.cli.GetEndpointNet(ns)
		if err != nil {
			return err
		}

		podNet, err := kinfo.cli.GetPodNet(ns)
		if err != nil {
			return err
		}

		for ip, list := range podNet {
			for _, v := range list {
				if !v.HostNetwork {
					k8sPodMap[ip] = v
				}
				if _, ok := k8sPodTmpNetMap[ip]; !ok {
					k8sPodTmpNetMap[ip] = make([]*K8sPodNet, 0)
				}
				k8sPodTmpNetMap[ip] = append(k8sPodTmpNetMap[ip], v)
			}
		}

		// for range services
		for name, svc := range svcNet {
			ep, ok := endpointNet[name]
			if !ok {
				continue
			}

			for _, v := range svc.ClusterIPs {
				k8sSvcMap[v] = svc
			}

			// svc‘ endpoint’ ip port
			// 取 endpoint 的 ip 和端口通过
			// label selector 匹配出 pod
			for ip, ports := range ep.IPPort {
				pods, ok := k8sPodTmpNetMap[ip]
				if ok {
					otherPods := []*K8sPodNet{}
					// for range pods
					for _, pod := range pods {
						flag := true
						for k, v := range svc.Selector {
							if tv, ok := pod.Labels[k]; ok && tv == v {
							} else {
								flag = false
								otherPods = append(otherPods, pod)
								break
							}
						}
						if flag {
							if _, ok := k8sPodNetPortMap[ip]; !ok {
								k8sPodNetPortMap[ip] = map[Port]*K8sPodNet{}
							}
							for _, port := range ports {
								pod.ServiceName = svc.Name
								k8sPodNetPortMap[ip][port] = pod
							}
						}
					}
					k8sPodTmpNetMap[ip] = otherPods
				}
				pod, ok := k8sPodMap[ip]
				if ok {
					pod.ServiceName = svc.Name
				}
			}
		}
	}

	kinfo.Lock()
	defer kinfo.Unlock()

	kinfo.svcNetInfo = k8sSvcMap
	kinfo.poNetInfoIP = k8sPodMap
	kinfo.poNetInfoPort = k8sPodNetPortMap

	return nil
}

func (kinfo *K8sNetInfo) AutoUpdate(ctx context.Context) {
	if kinfo.autoUpdate {
		return
	} else {
		kinfo.autoUpdate = true
	}

	go func() {
		ticker := time.NewTicker(time.Second * 9)
		for {
			select {
			case <-ticker.C:
				// TODO: log error
				_ = kinfo.Update()
			case <-ctx.Done():
				return
			}
		}
	}()
}

// QueryPodSvcName returns (pod name, svc name, namespace, port, err).
func (kinfo *K8sNetInfo) QueryPodSvcName(ip string, port uint32, protocol string) (string, string, string, uint32, error) {
	kinfo.RLock()
	defer kinfo.RUnlock()

	if port == 0 || port > 65535 {
		return "", "", "", 0, fmt.Errorf("no match pod")
	}
	pP := Port{
		Port: port,
	}
	switch protocol {
	case "tcp":
		pP.Protocol = "TCP"
	case "udp":
		pP.Protocol = "UDP"
	}
	if p, ok := kinfo.poNetInfoPort[ip]; ok {
		// 可能是 HostNetwork ip pod, 需要 port 辅助判定
		if v, ok := p[pP]; ok {
			return v.Name, v.ServiceName, v.Namespace, port, nil
		}
	}

	// 作为 client 发送请求的 pod， 不含(host network ip)
	if v, ok := kinfo.poNetInfoIP[ip]; ok {
		return v.Name, v.ServiceName, v.Namespace, 0, nil
	}

	return "", "", "", 0, fmt.Errorf("no match pod")
}

// QuerySvcName returns (svc name, namespace, error).
func (kinfo *K8sNetInfo) QuerySvcName(ip string) (string, string, error) {
	kinfo.RLock()
	defer kinfo.RUnlock()
	if v, ok := kinfo.svcNetInfo[ip]; ok {
		return v.Name, v.Namespace, nil
	}
	return "", "", fmt.Errorf("no match svc")
}

func NewK8sNetInfo(cli *K8sClient) (*K8sNetInfo, error) {
	kinfo := &K8sNetInfo{
		cli:        cli,
		autoUpdate: false,
	}

	if err := kinfo.Update(); err != nil {
		return nil, err
	}

	return kinfo, nil
}
