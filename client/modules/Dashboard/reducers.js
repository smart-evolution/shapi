import _ from 'lodash';
import * as actionTypes from './actionTypes';

const DATA_RANGE = 30;

const defaultState = {
  temperature: [],
  motions: {},
  presence: '-',
  error: '',
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

      const isMotion = Number(action.presence);

      const presence = isMotion
        ? 'Motion detected'
        : 'No motion';

      const motions = isMotion
        ? _.merge(state.motions, {[action.presence]: Date.now()})
        : state.motions;

      return Object.assign({}, state, {
        temperature,
        presence,
        motions,
        error: '',
      });

    case actionTypes.DATA_FETCH_ERROR:
      return Object.assign({}, state, { error: action.error });

    default:
      return state;
  }
}
