import _ from 'lodash';
import * as actionTypes from './actionTypes';

const defaultState = {
  time: [],
  temperature: [],
  motions: {},
  error: '',
};

export default function reducer(state = defaultState, action) {
  switch (action.type) {
    case actionTypes.DATA_FETCH:
      return Object.assign({}, state);

    case actionTypes.DATA_FETCHED:
      const temperature = _.map(action.temperature, (temperature, index) => ({
        time: new Date(action.time[index]),
        value: temperature,
      }));

      const motions = _.chain(action.motions)
        .filter(motion => motion > 0)
        .mapValues((motion, index) => ({
          time: new Date(action.time[index]),
          value: motion,
        }))
        .mapKeys(motion => motion.value)
        .value();

      return Object.assign({}, state, {
        temperature,
        motions,
        error: '',
      });

    case actionTypes.DATA_FETCH_ERROR:
      return Object.assign({}, state, { error: action.error });

    default:
      return state;
  }
}
