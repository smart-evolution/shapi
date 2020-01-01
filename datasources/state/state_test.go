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
	agent1 := type1.New("livingroom", "Living room", "192.168.1.2", types.Type1)
	agent2 := type1.New("bedroom", "Bed room", "192.168.1.3", types.Type1)
	agents := []agent.IAgent{agent1, agent2}

	p := mock.NewPersistanceMock(
		"db-uri",
		"smarthome",
		false,
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
	agent2 := agent.New("bedroom", "Bed room", "192.168.1.3", types.Type1)
	agents := []agent.IAgent{agent1, agent2}

	p := mock.NewPersistanceMock(
		"db-uri",
		"smarthome",
		false,
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

func TestAddAgent(t *testing.T) {
	agents := []agent.IAgent{}

	p := mock.NewPersistanceMock(
		"db-uri",
		"smarthome",
		true,
	)

	t.Run("Should add agent", func(t *testing.T) {
		s := New(p, agents)
		s.AddAgent("bedroom", "Bed room", "192.168.1.3", types.Type1)

		if len(s.model.Agents) != 1 {
			t.Errorf("Should have added one agent")
		}

		s.AddAgent("livingroom", "Living room", "192.168.1.2", types.Type1)

		if len(s.model.Agents) != 2 {
			t.Errorf("Should have added second agent")
		}
	})

	t.Run("Should not add duplicated agents", func(t *testing.T) {
		s := New(p, agents)
		s.AddAgent("bedroom", "Bed room 1", "192.168.1.2", types.Type1)
		s.AddAgent("bedroom", "Bed room 2", "192.168.1.3", types.Type1)

		if len(s.model.Agents) != 1 {
			t.Errorf("Should have added one agent")
		}
	})
}
