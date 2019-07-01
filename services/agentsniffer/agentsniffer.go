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

func scan(wg *sync.WaitGroup, ip string, s state.IState) {
	defer wg.Done()

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

	_, err = s.AgentByID(ip)
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
	var (
		wg sync.WaitGroup
	)

	for i := 1; i <= 255; i++ {
		ip := "192.168.2." + strconv.Itoa(i)
		wg.Add(1)
		go scan(&wg, ip, s)
	}
	wg.Wait()
}
