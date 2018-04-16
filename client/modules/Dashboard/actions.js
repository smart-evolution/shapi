import * as actionTypes from './actionTypes';

export const fetchDataSuccess = (times, temperatures, motions, gas) => ({
  type: actionTypes.DATA_FETCH_SUCCESS,
  times,
  temperatures,
  motions,
  gas,
});

export const fetchDataFail = (error) => ({
  type: actionTypes.DATA_FETCH_ERROR,
  error,
});

export const toggleAlerts = () => ({
  type: actionTypes.TOGGLE_ALERTS,
});

export const fetchAlerts = () => ({
  type: actionTypes.FETCH_ALERTS,
});

export const setAlerts = (isAlerts) => ({
  type: actionTypes.SET_ALERTS,
  isAlerts,
});
