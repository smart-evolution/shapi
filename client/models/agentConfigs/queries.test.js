import * as queries from './queries';

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

describe('models/agentConfigs/queries', () => {
  describe('getAgentConfigByAgentId', () => {
    it('should return searched AgentConfig', () => {
      const result = queries.getAgentConfigByAgentId(agentConfigs, '90346453');
      expect(result).toEqual(agentConfigs[1]);
    });

    it('should return empty object when no agentConfig found', () => {
      const result = queries.getAgentConfigByAgentId(agentConfigs, '84750192');
      expect(result).toEqual({});
    });
  });
});
