// @flow
import * as proxyConstants from 'client/models/proxy/constants';

export const getWsClient = (state: Object): WebSocket => {
  return state.proxy.wsClient;
};

export const getStatus = (state: Object): boolean => {
  return state.proxy.status;
};

export const getIsDevConnected = (state: Object): boolean => {
  return state.proxy.status === proxyConstants.STATUS_CONNECTED;
};
