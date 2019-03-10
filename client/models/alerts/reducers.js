import _ from 'lodash';
import * as actionTypes from './actionTypes';

const defaultState = {
  alerts: [],
};

export default function reducers(state = defaultState, action) {
  const { alerts } = state;
  const alert = {
    message: action.message,
    type: action.alertType,
    timestamp: new Date(),
    isOld: false,
  };

  switch (action.type) {
    case actionTypes.ADD:
      return Object.assign(
        {},
        {
          alerts: _.concat(alerts, [alert]),
        }
      );
    default:
      return state;
  }
}
