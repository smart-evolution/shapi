import * as actionTypes from './actionTypes';

export const addAlert = (message, alertType) => ({
  type: actionTypes.ADD,
  alertType,
  message,
});
