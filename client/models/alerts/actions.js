// @flow
import * as actionTypes from './actionTypes';
import * as types from './types';

/* eslint-disable import/prefer-default-export */
export const addAlert = (message: string, alertType: types.AlertType) => ({
  type: actionTypes.ADD,
  alertType,
  message,
});
/* eslint-enable import/prefer-default-export */
