import * as actionTypes from './actionTypes';

export const fetchDataSuccess = agents => ({
  type: actionTypes.DATA_FETCH_SUCCESS,
  agents,
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

export const sendAlert = () => ({
  type: actionTypes.SEND_ALERT,
});

export const setAlerts = isAlerts => ({
  type: actionTypes.SET_ALERTS,
  isAlerts,
});

export const toggleType2 = agentID => ({
  type: actionTypes.TOGGLE_TYPE2,
  agentID,
});
