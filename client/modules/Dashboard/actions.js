import * as actionTypes from './actionTypes';

export const fetchDataSuccess = (times, temperatures, motions, gas, sounds) => ({
  type: actionTypes.DATA_FETCH_SUCCESS,
  times,
  temperatures,
  motions,
  gas,
  sounds,
});

export const fetchDataFail = error => ({
  type: actionTypes.DATA_FETCH_ERROR,
  error,
});

export const toggleAlerts = () => ({
  type: actionTypes.TOGGLE_ALERTS,
});

export const fetchAlerts = () => ({
  type: actionTypes.FETCH_ALERTS,
});

export const setAlerts = isAlerts => ({
  type: actionTypes.SET_ALERTS,
  isAlerts,
});

export const sendAlert = () => ({
  type: actionTypes.SEND_ALERT,
});
