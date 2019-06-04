// @flow
import { type Agent } from './types';
/* eslint-disable import/prefer-default-export */
export const getAgents = (state: Object): $ReadOnlyArray<Agent> =>
  state.agents.agents;
/* eslint-enable import/prefer-default-export */
