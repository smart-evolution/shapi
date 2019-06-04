// @flow
import * as actionTypes from './actionTypes';

const defaultState: {
  wsClient: WebSocket | null,
} = {
  wsClient: null,
};

export default function reducers(state: State = defaultState, action: Action) {
  switch (action.type) {
    case actionTypes.PROXY_ADD_WS_CLIENT:
      return Object.assign({}, state, {
        wsClient: action.client,
      });

    default:
      return state;
  }
}
