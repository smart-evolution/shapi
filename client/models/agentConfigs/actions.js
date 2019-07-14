// @flow
import * as actionTypes from './actionTypes';
import * as types from './types';

export const fetchData = () => ({
  type: actionTypes.FETCH_DATA,
});

export const fetchDataSuccess = () => ({
  type: actionTypes.FETCH_AGENT_CONFIG_SUCCESS,
});

export const fetchDataFail = () => ({
  type: actionTypes.FETCH_AGENT_CONFIG_FAILURE,
});

export const updateDate = (agentID, config) => ({
  type: actionTypes.POST_AGENT_CONFIG,
  agentID,
  config,
});

export const updateDataSuccess = () => ({
  type: actionTypes.POST_AGENT_CONFIG_SUCCESS,
});

export const updateDataFail = () => ({
  type: actionTypes.POST_AGENT_CONFIG_FAILURE,
});

export const updateConfig = (agentID, config) => ({
  type: actionTypes.UPDATE_CONFIG,
  agentID,
  config,
});

export const updateProperty = (agentID, key, value) => ({
  type: actionTypes.UPDATE_PROPERTY,
  agentID,
  key,
  value,
});

export const loadAgentConfigs = (configs: $ReadOnlyArray<types.AgentConfig>) => ({
  type: actionTypes.LOAD_AGENT_CONFIGS,
  configs,
});
