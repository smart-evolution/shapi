import * as actionTypes from './actionTypes';

export const fetchData = agentID => ({
  type: actionTypes.FETCH_DATA,
  agentID,
});

export const fetchDataSuccess = () => ({
  type: actionTypes.FETCH_AGENT_CONFIG_SUCCESS,
});

export const fetchDataFail = () => ({
  type: actionTypes.FETCH_AGENT_CONFIG_FAILURE,
});

export const updateDate = (agentID, data) => ({
  type: actionTypes.POST_AGENT_CONFIG,
  agentID,
  data,
});

export const updateDataSuccess = () => ({
  type: actionTypes.POST_AGENT_CONFIG_SUCCESS,
});

export const updateDataFail = () => ({
  type: actionTypes.POST_AGENT_CONFIG_FAILURE,
});

export const updateTemperature = (agentID, temperature) => ({
  type: actionTypes.UPDATE_TEMPERATURE,
  agentID,
  temperature,
});
