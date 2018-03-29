import _ from 'lodash';
import * as actionTypes from './actionTypes';

const DATA_RANGE = 30;

const defaultState = {
  temperature: [],
  presence: "0",
};

export default function reducer(state = defaultState, action) {
  switch (action.type) {
    case actionTypes.DATA_FETCH:
      return Object.assign({}, state);

    case actionTypes.DATA_FETCHED:
      const temperatureObj = {
        time: new Date(),
        value: action.temperature,
      }
      const updatedTemps = _.concat(state.temperature, [temperatureObj]);

      const start = updatedTemps.length >= DATA_RANGE ? 1 : 0;

      const temperature = _.slice(updatedTemps, start, start + DATA_RANGE);

      const presence = Number(action.presence)
        ? "Motion detected"
        : "No motion";

      return Object.assign({}, state, {
        temperature,
        presence,
      });

    default:
      return state;
  }
}
