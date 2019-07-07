// @flow
import * as agentTypes from 'client/models/agents/types';
import * as actionTypes from './actionTypes';
import * as types from './types';

export const createWebSocketClient = (agent: agentTypes.Agent) => ({
  type: actionTypes.PROXY_CREATE_WS_CLIENT,
  agent,
});

export const addWebSocketClient = (
  agent: agentTypes.Agent,
  client: WebSocket
) => ({
  type: actionTypes.PROXY_ADD_WS_CLIENT,
  agent,
  client,
});

export const removeWebSocketClient = () => ({
  type: actionTypes.PROXY_REMOVE_WS_CLIENT,
});

export const sendMessage = (
  agent: agentTypes.Agent,
  message: types.Message
) => ({
  type: actionTypes.PROXY_SEND_MESSAGE,
  agent,
  message,
});

export const setDevStatus = (status: types.Status) => ({
  type: actionTypes.PROXY_SET_DEV_STATUS,
  status,
});
