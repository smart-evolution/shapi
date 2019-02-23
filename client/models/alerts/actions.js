import * as actionTypes from './actionTypes';

/* eslint-disable import/prefer-default-export */
export const addAlert = (message, alertType) => ({
  type: actionTypes.ADD,
  alertType,
  message,
});
/* eslint-enable import/prefer-default-export */
