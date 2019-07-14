// @flow
import * as agentTypes from 'client/models/agents/types';

export type AgentConfig = {
  id: agentTypes.AgentID,
  name: string,
  temperature: number,
};

export type AgentConfigs = $ReadOnlyArray<AgentConfig>;

export type State = {
  agentConfigs: AgentConfig,
  error: string,
};
