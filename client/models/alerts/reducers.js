import _ from 'lodash';
import * as actionTypes from './actionTypes';

const defaultState = {
  alerts: [],
};

export default function reducers(state = defaultState, action) {
  const { alerts } = state;

  switch (action.type) {
    case actionTypes.ADD:
      const alert = {
        message: action.message,
        type: action.alertType,
        timestamp: new Date(),
        isOld: false,
      };
      return Object.assign({}, {
        alerts: _.concat(alerts, [alert]),
      });
      break;
    default:
      return state;
  }
}
