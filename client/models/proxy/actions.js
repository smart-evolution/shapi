// @flow
import * as agentTypes from 'models/agents/types';
import * as actionTypes from './actionTypes';

export const createWebSocketClient = () => ({
  type: actionTypes.PROXY_CREATE_WS_CLIENT,
});

export const addWebSocketClient = (agent: agentTypes.Agent, client: WebSocket) => ({
  type: actionTypes.PROXY_ADD_WS_CLIENT,
  agent,
  client,
});

export const sendMessage = (agent: agentTypes.Agent, message: string) => ({
  type: actionTypes.PROXY_SEND_MESSAGE,
  agent,
  message,
});
