// @flow
import _ from 'lodash';
import * as types from './types';

export const getAgents = (state: Object): $ReadOnlyArray<types.Agent> =>
  state.agents.agents;

export const isLoading = (state: Object): $ReadOnlyArray<types.Agent> =>
  state.agents.isLoading;

export const isAlerts = (state: Object): $ReadOnlyArray<types.Agent> =>
  state.agents.isAlerts;

export const getAgentById = (state: Object, id: string) => {
  const agents = getAgents(state);
  return _.find(agents, { id });
};

export const getPeriod = (state: Object): number => state.agents.period;
