import * as actionTypes from './actionTypes';

export const fetchData = (agentID) => ({
  type: actionTypes.FETCH_DATA,
  agentID,
});

export const fetchDataSuccess = () => ({
  type: actionTypes.FETCH_AGENT_CONFIG_SUCCESS,
});

export const fetchDataFail = () => ({
  type: actionTypes.FETCH_AGENT_CONFIG_FAILURE,
});
