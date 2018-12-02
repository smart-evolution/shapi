package homebot

import (
    "errors"
    "testing"
    "reflect"
    "github.com/smart-evolution/smarthome/models/agent"
)

func TestFindAgentByID(t *testing.T) {
    var agent1 = agent.Agent{
        "livingroom",
        "Living room",
        "192.168.1.2",
        "type1",
    }
    var agent2 = agent.Agent{
        "bedroom",
        "Bed room",
        "192.168.1.3",
        "type2",
    }

    hb := HomeBot{
        Agents: []agent.Agent{agent1, agent2},
    }

    t.Run("Should find agent by matching ID", func(t *testing.T) {
        expectedResult := agent1
        result, _ := hb.FindAgentByID("livingroom")

        if !reflect.DeepEqual(expectedResult, result) {
            t.Errorf("Non agents id match by ID")
        }
    })

    t.Run("Should return error when no agent match by ID", func(t *testing.T) {
        expectedResult := errors.New("No matching agent")
        _, err := hb.FindAgentByID("kidsroom")

        if err.Error() != expectedResult.Error() {
            t.Errorf("Non agents id match by ID")
        }
    })
}
