// @flow
import _ from 'lodash';
import * as actionTypes from './actionTypes';
import * as types from './types';

const defaultState: types.State = {
  alerts: [],
};

export default function reducers(
  state: types.State = defaultState,
  action: Function
) {
  const { alerts } = state;
  const alert: types.Alert = {
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
