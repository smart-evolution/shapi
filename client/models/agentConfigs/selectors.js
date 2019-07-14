// @flow
import * as agentTypes from 'client/models/agents/types';
import * as types from './types';
import * as queries from './queries';

export const getAgentConfigs = (
  state: Object
): $ReadOnlyArray<types.AgentConfig> => {
  return state.agentConfigs.agentConfigs || [];
};

export const getAgentConfigById = (
  state: Object,
  agentID: agentTypes.AgentID
): types.AgentConfig => {
  const agentConfigs = getAgentConfigs(state);
  const config = queries.getAgentConfigByAgentId(agentConfigs, agentID);
  return config;
};
