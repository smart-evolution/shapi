// @flow
export type AgentID = string;

export type Agent = {
  id: AgentID,
  name: string,
  data: any,
  type: string,
  ip: string,
  isOnline: boolean,
};

export type State = {
  isLoading: boolean,
  error: string,
  agents: $ReadOnlyArray<Agent>,
  period: number,
};
