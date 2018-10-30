import _ from 'lodash';
import * as actionTypes from './actionTypes';

const defaultState = {
  error: '',
  agents: [],
};

export default function reducer(state = defaultState, action) {
  const { agents, error } = action;

  switch (action.type) {
    case actionTypes.DATA_FETCH:
      return Object.assign({}, state);

    case actionTypes.DATA_FETCH_SUCCESS:
      return Object.assign({}, state, { agents });

    case actionTypes.DATA_FETCH_ERROR:
      return Object.assign({}, state, { error });

    default:
      return state;
  }
}
