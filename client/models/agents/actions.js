// @flow
import * as actionTypes from './actionTypes';
import * as types from './types';

export const fetchData = () => ({
  type: actionTypes.DATA_FETCH,
});

export const fetchDataSuccess = (agents: $ReadOnlyArray<types.Agent>) => ({
  type: actionTypes.DATA_FETCH_SUCCESS,
  agents,
});

export const fetchDataFail = (error: string) => ({
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

export const setAlerts = (isAlerts: boolean) => ({
  type: actionTypes.SET_ALERTS,
  isAlerts,
});

export const toggleType2 = (agentID: string) => ({
  type: actionTypes.TOGGLE_TYPE2,
  agentID,
});

export const sniffAgents = () => ({
  type: actionTypes.SNIFF_AGENTS,
});
