package agentsniffer

import (
	"github.com/smart-evolution/smarthome/constants"
	"github.com/smart-evolution/smarthome/datasources/state"
	"github.com/smart-evolution/smarthome/utils"
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
	stations      = 254
)

var (
	mutex      = &sync.Mutex{}
	isSniffing = false
)

func scan(wg *sync.WaitGroup, ip string, s state.IState) {
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
		defer resp.Body.Close()

		contents, err := ioutil.ReadAll(resp.Body)
		head := strings.Split(string(contents), "\n")[1]
		hardwareVal := strings.Split(head, "=")[1]
		hardwareID := hardwareVal[1 : len(hardwareVal)-1]

		if err != nil {
			utils.Log("failed to fetch config of agent with IP:" + ip)
		} else {
			mutex.Lock()
			s.AddAgent(hardwareID, hardwareID, ip, devType)
			mutex.Unlock()
		}
	}
}

// SniffAgents - function looking for agents by sending CMDWHO
func SniffAgents(s state.IState) {
	if !isSniffing {
		isSniffing = true

		var wg sync.WaitGroup
		done := make(chan struct{})
		wg.Add(subNetworks * stations)

		utils.Log("sniffing devices within range " + strconv.Itoa(subNetworks) + "." + strconv.Itoa(stations))
		for i := 1; i <= subNetworks; i++ {
			for j := 1; j <= stations; j++ {
				ip := "192.168." + strconv.Itoa(i) + "." + strconv.Itoa(j)
				go scan(&wg, ip, s)
			}
		}

		go func() {
			wg.Wait()
			close(done)
		}()

		select {
		case <-done:
			isSniffing = false
			utils.Log("sniffing devices completed")
			return
		case <-time.After(sniffTimeout * time.Millisecond):
			isSniffing = false
			utils.Log("sniffing devices time out")
			return
		}
	}
}
