// @flow
import * as actionTypes from './actionTypes';
import * as types from './types';

const defaultState: {
  isLoading: boolean,
  error: string,
  agents: $ReadOnlyArray<types.Agent>,
} = {
  isLoading: false,
  error: '',
  agents: [],
};

export default function reducers(
  state: Object = defaultState,
  action: Function
) {
  const { agents, error } = action;

  switch (action.type) {
    case actionTypes.DATA_FETCH:
      return Object.assign({}, state, { isLoading: true });

    case actionTypes.DATA_FETCH_SUCCESS:
      return Object.assign({}, state, { agents, isLoading: false });

    case actionTypes.DATA_FETCH_ERROR:
      return Object.assign({}, state, { error, isLoading: false });

    case actionTypes.SET_ALERTS:
      return Object.assign({}, state, { isAlerts: action.isAlerts });

    default:
      return state;
  }
}
