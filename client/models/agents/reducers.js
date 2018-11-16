import * as actionTypes from './actionTypes';

const defaultState = {
  error: '',
  agents: [],
};

export default function reducers(state = defaultState, action) {
  const { agents, error } = action;

  switch (action.type) {
    case actionTypes.DATA_FETCH:
      return Object.assign({}, state);

    case actionTypes.DATA_FETCH_SUCCESS:
      return Object.assign({}, state, { agents });

    case actionTypes.DATA_FETCH_ERROR:
      return Object.assign({}, state, { error });

    case actionTypes.SET_ALERTS:
      return Object.assign({}, state, { isAlerts: action.isAlerts });

    default:
      return state;
  }
}
