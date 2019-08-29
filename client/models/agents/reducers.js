// @flow
import * as actionTypes from './actionTypes';
import * as types from './types';

const defaultState = {
  isLoading: true,
  error: '',
  agents: [],
  period: 30,
};

export default function reducers(
  state: types.State = defaultState,
  action: Object
) {
  const { agents, error } = action;

  switch (action.type) {
    case actionTypes.FETCH_AGENTS:
      return Object.assign({}, state, { isLoading: true });

    case actionTypes.LOAD_AGENTS:
      return Object.assign({}, state, { agents, isLoading: false });

    case actionTypes.FETCH_AGENTS_ERROR:
      return Object.assign({}, state, { error, isLoading: false });

    case actionTypes.SET_ALERTS:
      return Object.assign({}, state, { isAlerts: action.isAlerts });

    case actionTypes.CHANGE_PERIOD:
      return Object.assign({}, state, { period: action.period });

    default:
      return state;
  }
}
