import _ from 'lodash';
import * as actionTypes from './actionTypes';

const DATA_RANGE = 30;

const defaultState = {
    temperature: [],
};

export default function reducer(state = defaultState, action) {
  switch (action.type) {
    case actionTypes.TEMPERATURE_CHART_FETCH:
      return Object.assign({}, state);

    case actionTypes.TEMPERATURE_CHART_FETCHED:
      const temperatureObj = {
        time: new Date(),
        value: action.temperature,
      }
      const updatedTemps = _.concat(state.temperature, [temperatureObj]);

      const start = updatedTemps.length >= DATA_RANGE ? 1 : 0;

      const temperature = _.slice(updatedTemps, start, start + DATA_RANGE);

      return Object.assign({}, state, {
        temperature,
      });

    default:
      return state;
  }
}
