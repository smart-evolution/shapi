package agentsniffer

import (
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
	SUB_NETWORKS = 255
	STATIONS     = 255
)

var (
	mutex      = &sync.Mutex{}
	isSniffing = false
)

func scan(wg *sync.WaitGroup, ip string, s state.IState) {
	defer wg.Done()
	utils.Log("sniffing device with IP: " + ip)

	d := net.Dialer{Timeout: time.Duration(1000) * time.Millisecond}
	conn, err := d.Dial("tcp", ip+":81")
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
			s.AddAgent(hardwareID, hardwareID, ip, devType)
		}
	}
}

func SniffAgents(s state.IState) {
	if !isSniffing {
		mutex.Lock()
		isSniffing = true
		mutex.Unlock()

		var wg sync.WaitGroup
		done := make(chan struct{})
		wg.Add(SUB_NETWORKS * STATIONS)

		for i := 1; i <= SUB_NETWORKS; i++ {
			for j := 1; j <= STATIONS; j++ {
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
			mutex.Lock()
			isSniffing = false
			mutex.Unlock()
			return
		case <-time.After(3 * time.Second):
			mutex.Lock()
			isSniffing = false
			mutex.Unlock()
			return
		}
	}
}
