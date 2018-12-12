package state

import (
    "fmt"
    "errors"
    "testing"
    "reflect"
    "github.com/smart-evolution/smarthome/models/agent"
)

func TestAgentByID(t *testing.T) {
    agent1 := agent.New("livingroom", "Living room", "192.168.1.2", "type1")
    agent2 := agent.New("bedroom", "Bed room", "192.168.1.3", "type2")
    agents := []*agent.Agent{agent1, agent2}

    s := New(agents)

    fmt.Println(s.Agents())

    t.Run("Should find agent by matching ID", func(t *testing.T) {
        expectedResult := agent1
        result, _ := s.AgentByID("livingroom")

        if !reflect.DeepEqual(expectedResult, result) {
            t.Errorf("Non agents id match by ID")
        }
    })

    t.Run("Should return error when no agent match by ID", func(t *testing.T) {
        expectedResult := errors.New("No matching agent")
        _, err := s.AgentByID("kidsroom")

        if err.Error() != expectedResult.Error() {
            t.Errorf("Wrong error returned")
        }
    })
}
