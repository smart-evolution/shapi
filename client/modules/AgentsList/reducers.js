import * as actionTypes from './actionTypes';

const defaultState = {
  agents: [],
};

export default function reducer(state = defaultState, action) {
  const { agents } = action;

  switch (action.type) {
    case actionTypes.DATA_FETCH:
      return Object.assign({}, state);

    case actionTypes.DATA_FETCH_SUCCESS:
      return Object.assign({}, state, { agents });

    case actionTypes.DATA_FETCH_ERROR:
      return Object.assign({}, state, { error: action.error });

    default:
      return state;
  }
}
