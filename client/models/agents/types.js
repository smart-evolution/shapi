// @flow
export type Agent = {
  id: string,
  name: string,
  data: any,
  type: string,
  ip: string,
};

export type State = {
  isLoading: boolean,
  error: string,
  agents: $ReadOnlyArray<Agent>,
};
