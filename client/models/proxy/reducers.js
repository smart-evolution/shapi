// @flow
import * as actionTypes from './actionTypes';

type State = {
  wsClient: WebSocket | null,
};

const defaultState = {
  wsClient: null,
};

export default function reducers(state: State = defaultState, action: Object) {
  switch (action.type) {
    case actionTypes.PROXY_ADD_WS_CLIENT:
      return Object.assign({}, state, {
        wsClient: action.client,
      });

    default:
      return state;
  }
}
