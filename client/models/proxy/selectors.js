// @flow
export const getWsClient = (state: Object): WebSocket => {
  return state.proxy.wsClient;
};

export const getIsDevConnected = (state: Object): boolean => {
  return state.proxy.isDevConnected;
};

