// @flow
import * as actionTypes from './actionTypes';

type State = {
  wsClient: WebSocket | null,
  isDevConnected: boolean,
};

const defaultState = {
  wsClient: null,
  isDevConnected: false,
};

export default function reducers(state: State = defaultState, action: Object) {
  switch (action.type) {
    case actionTypes.PROXY_ADD_WS_CLIENT:
      return Object.assign({}, state, {
        wsClient: action.client,
      });

    case actionTypes.PROXY_SEND_MESSAGE:
      return Object.assign({}, state, {
        isDevConnected: action.isDevConnected,
      });

    default:
      return state;
  }
}
