import * as actionTypes from './actionTypes';

const defaultState = {
  isAlerts: false,
  times: [],
  temperatures: [],
  motions: {},
  gas: false,
  sounds: [],
  error: '',
};

export default function reducer(state = defaultState, action) {
  const { times, temperatures, motions, gas, sounds } = action;

  switch (action.type) {
    case actionTypes.DATA_FETCH:
      return Object.assign({}, state);

    case actionTypes.DATA_FETCH_SUCCESS:
      return Object.assign({}, state, {
        times,
        temperatures,
        motions,
        gas,
        sounds,
        error: '',
      });

    case actionTypes.DATA_FETCH_ERROR:
      return Object.assign({}, state, { error: action.error });

    default:
      return state;
  }
}
