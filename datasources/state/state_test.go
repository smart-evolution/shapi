package state

import (
    "errors"
    "testing"
    "reflect"
    "github.com/smart-evolution/smarthome/models/agent"
)

func TestFindAgentByID(t *testing.T) {
    s := New()
    s.AddAgent("livingroom", "Living room", "192.168.1.2", "type1")
    s.AddAgent("bedroom", "Bed room", "192.168.1.3", "type2")

    agent1 := agent.New("livingroom", "Living room", "192.168.1.2", "type1")

    t.Run("Should find agent by matching ID", func(t *testing.T) {
        expectedResult := agent1
        result, _ := s.FindAgentByID("livingroom")

        if !reflect.DeepEqual(expectedResult, result) {
            t.Errorf("Non agents id match by ID")
        }
    })

    t.Run("Should return error when no agent match by ID", func(t *testing.T) {
        expectedResult := errors.New("No matching agent")
        _, err := s.FindAgentByID("kidsroom")

        if err.Error() != expectedResult.Error() {
            t.Errorf("Non agents id match by ID")
        }
    })
}
