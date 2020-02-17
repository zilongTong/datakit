package telegrafwrap

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
	"time"

	"gitlab.jiagouyun.com/cloudcare-tools/datakit/internal"

	"gitlab.jiagouyun.com/cloudcare-tools/datakit/config"
)

type (
	TelegrafSvr struct {
		MainCfg *config.MainConfig
	}
)

var Svr = &TelegrafSvr{}

func (s *TelegrafSvr) Start(ctx context.Context) error {

	if s.MainCfg.CustomAgentConfigFile == "" {
		telcfg, err := s.GenerateTelegrafConfig()
		if err == config.ErrNoTelegrafConf {
			log.Printf("no sub service configuration found")
			return nil
		}

		if err != nil {
			return fmt.Errorf("fail to generate sub service config, %s", err)
		}

		if err = ioutil.WriteFile(s.agentConfPath(false), []byte(telcfg), 0664); err != nil {
			return fmt.Errorf("fail to create file, %s", err.Error())
		}
	}

	log.Printf("starting sub service...")

	if err := s.startAgent(ctx); err != nil {
		return err
	}

	//检查telegraf是否被kill掉
	go func(ctx context.Context) {

		for {

			select {
			case <-ctx.Done():
				s.StopAgent()
				return
			default:
			}

			piddata, err := ioutil.ReadFile(agentPidPath())
			if err != nil {
				log.Printf("W! read pid file failed, %s", err)
			}

			npid, err := strconv.Atoi(string(piddata))
			if err != nil || npid <= 2 {
				log.Printf("W! invalid pid, %s", err)
			}

			if CheckPid(npid) != nil {
				log.Printf("W! sub service(%v) was killed", npid)
			} else {
				_, err := os.FindProcess(npid)
				if err != nil {
					log.Printf("W! sub service(%v) not found: %s", npid, err.Error())
				}
			}

			internal.SleepContext(ctx, 3*time.Second)
		}
	}(ctx)

	return nil
}

func (s *TelegrafSvr) StopAgent() error {

	piddata, err := ioutil.ReadFile(agentPidPath())
	if err != nil {
		return err
	}
	if string(piddata) == "" {
		return nil
	}
	npid, err := strconv.Atoi(string(piddata))
	if err != nil {
		return err
	}

	if npid > 0 {
		return s.killProcessByPID(npid)
	}
	return nil
}

func (s *TelegrafSvr) startAgent(ctx context.Context) error {

	s.StopAgent()

	env := os.Environ()
	if runtime.GOOS == "windows" {
		env = append(env, fmt.Sprintf(`TELEGRAF_CONFIG_PATH=%s`, s.agentConfPath(false)))
	}
	procAttr := &os.ProcAttr{
		Env: env,
		Files: []*os.File{
			os.Stdin,
			os.Stdout,
			os.Stderr,
		},
	}

	var p *os.Process
	var err error

	if runtime.GOOS == "windows" {

		// pg := filepath.Join(config.ExecutableDir, "agent_debug.log")

		// f, _ := os.Create(pg)
		// procAttr.Files = []*os.File{
		// 	nil,
		// 	f,
		// 	f,
		// }

		cmd := exec.Command(agentPath(true), "-console")
		//cmd := exec.CommandContext(ctx, agentPath(true), "-console")
		cmd.Env = env
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		if err := cmd.Start(); err != nil {
			return err
		}
		p = cmd.Process

	} else {
		p, err = os.StartProcess(agentPath(false), []string{"agent", "-config", s.agentConfPath(false)}, procAttr)
		if err != nil {
			return fmt.Errorf("start agent failed, %s", err)
		}
	}

	if p != nil {
		ioutil.WriteFile(agentPidPath(), []byte(fmt.Sprintf("%d", p.Pid)), 0664)
	}

	time.Sleep(time.Millisecond * 100)

	return nil
}

func (s *TelegrafSvr) killProcessByPID(npid int) error {

	if CheckPid(npid) == nil {
		prs, err := os.FindProcess(npid)
		if err == nil && prs != nil {

			if err = prs.Kill(); err != nil {
				return err
			}
			time.Sleep(time.Millisecond * 100)
		}
	}

	return nil
}

func (s *TelegrafSvr) agentConfPath(quote bool) string {
	path := ""
	if s.MainCfg.CustomAgentConfigFile != "" {
		path = s.MainCfg.CustomAgentConfigFile
	} else {
		path = filepath.Join(config.ExecutableDir, "agent.conf")
	}

	if quote {
		return fmt.Sprintf(`"%s"`, path)
	}
	return path
}

func agentPidPath() string {
	return filepath.Join(config.ExecutableDir, "agent.pid")
}

func agentPath(win bool) string {
	if win {
		return filepath.Join(config.ExecutableDir, "agent.exe")
	}
	return filepath.Join(config.ExecutableDir, "agent")
}
