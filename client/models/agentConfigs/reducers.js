// @flow
import _ from 'lodash';
import * as actionTypes from './actionTypes';
import * as types from './types';
import * as queries from './queries';

const defaultState = {
  error: '',
  agentConfigs: [],
};

function updateProperty(state, action) {
  const { key, value, agentID } = action;
  const { agentConfigs } = state;

  const agentConfig = queries.getAgentConfigByAgentId(agentConfigs, agentID) || {};
  agentConfig[key] = value;

  const newAgentConfig = _.defaults(agentConfigs, {
    [action.key]: action.value,
  });

  agentConfigs[agentID] = newAgentConfig;

  const newState = Object.assign({}, state, {
    agentConfigs,
  });

  return newState;
}

export default function reducers(state: types.State = defaultState, action) {
  const { agentID } = action;

  switch (action.type) {
    case actionTypes.LOAD_AGENT_CONFIGS:
      const newState = Object.assign({}, state, { agentConfigs: action.configs });
      return newState;

    case actionTypes.UPDATE_CONFIG:
      return _.merge({}, state, {
        agentConfigs: {
          [agentID]: action.config,
        },
      });

    case actionTypes.UPDATE_PROPERTY:
      return updateProperty(state, action);

    default:
      return state;
  }
}
