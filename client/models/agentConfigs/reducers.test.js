import _ from 'lodash';
import reducers from './reducers';
import * as actions from './actions';

const agentConfigs = [
  {
    id: '5d2af74d7b43643c977d23c2',
    agentId: '43366411',
    name: 'Badroom',
    temperature: 0,
  },
  {
    id: '4l2af74d7b43643c977d23e1',
    agentId: '90346453',
    name: 'Livingroom',
    temperature: 0,
  },
];

const state = {
  agentConfigs,
};

describe('models/agentConfigs/reducers', () => {
  describe('UPDATE_PROPERTY', () => {
    it('should add new AgentConfig to state', () => {
      const action = actions.updateProperty('12345678', 'name', 'New Agent');
      const expectedAgentConfigs = _.concat(agentConfigs, [
        {
          agentId: '12345678',
          name: 'New Agent',
        },
      ]);
      const expectedState = {
        agentConfigs: expectedAgentConfigs,
      };

      const result = reducers(state, action);
      expect(result).toEqual(expectedState);
    });

    it('should update existing AgentConfig', () => {
      const action = actions.updateProperty('90346453', 'name', 'New Agent');
      const expectedAgentConfigs = [
        agentConfigs[0],
        {
          id: '4l2af74d7b43643c977d23e1',
          agentId: '90346453',
          name: 'New Agent',
          temperature: 0,
        },
      ];
      const expectedState = {
        agentConfigs: expectedAgentConfigs,
      };
      const result = reducers(state, action);
      expect(result).toEqual(expectedState);
    });
  });
});
