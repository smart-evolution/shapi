// @flow
import * as types from './types';

export const getAgents = (state: Object): $ReadOnlyArray<types.Agent> =>
  state.agents.agents;

export const isLoading = (state: Object): $ReadOnlyArray<types.Agent> =>
  state.agents.isLoading;
