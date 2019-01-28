import _ from 'lodash';
import * as actionTypes from './actionTypes';

const defaultState = {
  error: '',
  agentConfigs: {},
};

export default function reducers(state = defaultState, action) {
  const { agentID, temperature } = action;

  switch (action.type) {
    case actionTypes.UPDATE_TEMPERATURE:
      return _.merge({}, state, {
        agentConfigs: {
          [agentID]: { temperature },
        },
      });

    default:
      return state;
  }
}
