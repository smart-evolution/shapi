package state

import (
	"errors"
	"github.com/smart-evolution/shapi/mock"
	"github.com/smart-evolution/shapi/models/agent"
	"github.com/smart-evolution/shapi/models/agent/types"
	"github.com/smart-evolution/shapi/models/type1"
	"reflect"
	"testing"
)

func TestAgentByID(t *testing.T) {
	agent1 := agent.New("livingroom", "Living room", "192.168.1.2", types.Type1)
	agent2 := agent.New("bedroom", "Bed room", "192.168.1.3", types.Type2)
	agents := []agent.IAgent{agent1, agent2}

	p := mock.NewPersistanceMock(
		"db-uri",
		"smarthome",
	)

	s := New(p, agents)

	t.Run("Should find agent by matching ID", func(t *testing.T) {
		expectedResult := agent1
		result, _ := s.AgentByID("livingroom")

		if !reflect.DeepEqual(expectedResult, result) {
			t.Errorf("Non agents id match by ID")
		}
	})

	t.Run("Should return error when no agent match by ID", func(t *testing.T) {
		expectedResult := errors.New("no matching agent")
		_, err := s.AgentByID("kidsroom")

		if err.Error() != expectedResult.Error() {
			t.Errorf("no corresponding agent found")
		}
	})
}

func TestRemoveAgent(t *testing.T) {
	agent1 := type1.New("livingroom", "Living room", "192.168.1.2", types.Type1)
	agent2 := agent.New("bedroom", "Bed room", "192.168.1.3", types.Type2)
	agents := []agent.IAgent{agent1, agent2}

	p := mock.NewPersistanceMock(
		"db-uri",
		"smarthome",
	)

	s := New(p, agents)

	t.Run("Should remove type1 agent by ID", func(t *testing.T) {
		err := s.RemoveAgent("livingroom")

		if err != nil {
			t.Errorf("Should have removed agent")
		}

		if len(s.model.Agents) != 1 {
			t.Errorf("Agent array shoube be of lenght 1")
		}
	})

	t.Run("Should throw error when no agent found", func(t *testing.T) {
		expectedResult := errors.New("no corresponding agent found")
		err := s.RemoveAgent("no-existing-id")

		if err.Error() != expectedResult.Error() {
			t.Errorf("Wrong error returned")
		}
	})
}
