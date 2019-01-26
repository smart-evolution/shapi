import _ from 'lodash';
import * as actionTypes from './actionTypes';

const defaultState = {
  error: '',
  agentConfigs: {},
};

export default function reducers(state = defaultState, action) {
  switch (action.type) {
    case actionTypes.UPDATE_TEMPERATURE:
      const { agentID, temperature } = action;

      return _.merge({}, state, {
        agentConfig: {
          [agentID]: { temperature },
        },
      });

    default:
      return state;
  }
}
