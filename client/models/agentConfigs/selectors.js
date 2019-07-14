// @flow
import _ from 'lodash';
import * as types from './types';
import * as queries from './queries';

export const getAgentConfigs = (state: Object): $ReadOnlyArray<types.AgentConfig> => {
  return state.agentConfigs.agentConfigs || [];
};

export const getAgentConfigById = (state: Object, agentID: number): types.AgentConfig => {
  const agentConfigs = getAgentConfigs(state);
  const config = queries.getAgentConfigByAgentId(agentConfigs, agentID);
  return config;
};

