// @flow
export type Agent = {
  id: string,
  name: string,
  data: any,
  agentType: string,
};

export type State = {
  isLoading: boolean,
  error: string,
  agents: $ReadOnlyArray<Agent>,
};
