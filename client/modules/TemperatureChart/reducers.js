import * as actionTypes from './actionTypes';

const defaultState = {
    temperature: 0,
};

export default function reducer(state = defaultState, action) {
  switch (action.type) {
      case actionTypes.TEMPERATURE_CHART_FETCH:
          return Object.assign({}, state);
          break;

      case actionTypes.TEMPERATURE_CHART_FETCHED:
          return Object.assign({}, state, {
              temperature: action.temperature,
          });
          break;

    default:
      return state;
  }
}
