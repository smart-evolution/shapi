package state

import (
	"github.com/smart-evolution/shapi/models/agent"
)

type AgentList []agent.IAgent

type State struct {
	IsAlerts  bool `bson:"isAlerts"`
	SendAlert bool `bson:"sendAlert"`
	Agents    AgentList
}

//func (al *AgentList) SetBSON(raw bson.Raw) error {
//	fmt.Printf("raw:%s\n", raw)
//	fmt.Println(raw.Data)
//	*al = AgentList{}
//	*al = append(*al, agent.New("das", "dsa", "dsa", "das"))
//	return nil
//}
//
//func (al *AgentList) GetBSON() (interface{}, error) {
//	fmt.Println("------ In GetBSON")
//	l := AgentList{}
//	l = append(l, agent.New("xx", "dsdas", "192.168.1.1", "bson"), nil)
//	return l, nil
//}
