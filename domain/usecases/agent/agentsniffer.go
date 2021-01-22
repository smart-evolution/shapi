package agent

import (
	"github.com/coda-it/goutils/logger"
	"github.com/smart-evolution/shapi/constants"
	"io/ioutil"
	"net"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"
)

const (
	sniffTimeout = 5000
	subNetworks  = 2
	stations     = 254
)

var (
	mutex      = &sync.Mutex{}
	isSniffing = false
)

func scan(wg *sync.WaitGroup, ip string, s IStateRepository) {
	defer wg.Done()

	d := net.Dialer{Timeout: time.Duration(sniffTimeout) * time.Millisecond}
	conn, err := d.Dial("tcp", ip+":"+constants.AgentTCPPort)
	if err != nil {
		return
	}

	_, err = conn.Write([]byte("CMDWHO"))

	if err != nil {
		return
	}

	buff := make([]byte, 512)
	n, err := conn.Read(buff)

	if err != nil {
		return
	}

	devType := string(buff[:n])

	_, err = s.AgentByIP(ip)
	if err != nil {
		resp, err := http.Get("http://" + ip + "/config")

		if err != nil {
			logger.Log("agent doesn't provide /config endpoint")
			return
		}

		defer resp.Body.Close()

		contents, err := ioutil.ReadAll(resp.Body)
		head := strings.Split(string(contents), "\n")[1]
		hardwareVal := strings.Split(head, "=")[1]
		hardwareID := hardwareVal[1 : len(hardwareVal)-1]

		if err != nil {
			logger.Log("failed to fetch config of agent with IP:" + ip)
		} else {
			mutex.Lock()
			s.AddAgent(hardwareID, hardwareID, ip, devType)
			mutex.Unlock()
		}
	}
}

// SniffAgents - function looking for agents by sending CMDWHO
func (u *Usecase) SniffAgents() {
	if !isSniffing {
		isSniffing = true

		var wg sync.WaitGroup
		done := make(chan struct{})
		wg.Add(subNetworks * stations)

		logger.Log("sniffing devices within range " + strconv.Itoa(subNetworks) + "." + strconv.Itoa(stations))
		for i := 1; i <= subNetworks; i++ {
			for j := 1; j <= stations; j++ {
				ip := "192.168." + strconv.Itoa(i) + "." + strconv.Itoa(j)
				go scan(&wg, ip, u.stateRepository)
			}
		}

		go func() {
			wg.Wait()
			close(done)
		}()

		select {
		case <-done:
			isSniffing = false
			logger.Log("sniffing devices completed")
			return
		case <-time.After(sniffTimeout * time.Millisecond):
			isSniffing = false
			logger.Log("sniffing devices time out")
			return
		}
	}
}
