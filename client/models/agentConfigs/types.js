// @flow
import * as agentTypes from 'client/models/agents/types';

export type AgentConfig = {
  id: string,
  agentId: agentTypes.AgentID,
  name: string,
  temperature: number,
};

export type AgentConfigs = Array<AgentConfig>;

export type State = {
  agentConfigs: AgentConfigs,
  error: string,
};
