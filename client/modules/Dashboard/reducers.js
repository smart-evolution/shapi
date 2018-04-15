import _ from 'lodash';
import * as actionTypes from './actionTypes';

const defaultState = {
  isAlerts: false,
  temperatures: [],
  motions: {},
  gas: false,
  error: '',
};

export default function reducer(state = defaultState, action) {
  switch (action.type) {
    case actionTypes.DATA_FETCH:
      return Object.assign({}, state);

    case actionTypes.DATA_FETCH_SUCCESS:
      const temperatures = _.map(action.temperatures, (temperature, index) => ({
        time: new Date(action.times[index]),
        value: temperature,
      }));

      const motions = _.chain(action.motions)
        .filter(motion => motion > 0)
        .mapValues((motion, index) => ({
          time: new Date(action.times[index]),
          value: motion,
        }))
        .mapKeys(motion => motion.value)
        .value();

      const gas = _.some(action.gas, a => a == '0');

      return Object.assign({}, state, {
        temperatures,
        motions,
        gas,
        error: '',
      });

    case actionTypes.DATA_FETCH_ERROR:
      return Object.assign({}, state, { error: action.error });

    case actionTypes.SET_ALERTS:
      return Object.assign({}, state, { isAlerts: action.isAlerts });

    default:
      return state;
  }
}
