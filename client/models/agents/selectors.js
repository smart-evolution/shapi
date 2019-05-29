// @flow
import * as types from './types';
/* eslint-disable import/prefer-default-export */
export const getAgents = (state: State): $ReadOnlyArray<types.Agent> =>
  state.agents.agents;
/* eslint-enable import/prefer-default-export */
