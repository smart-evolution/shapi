// @flow
import * as actionTypes from './actionTypes';
import * as types from './types';
import * as constants from './constants';

type State = {
  wsClient: WebSocket | null,
  status: types.Status,
};

const defaultState = {
  wsClient: null,
  status: constants.STATUS_DISCONNECTED,
};

export default function reducers(state: State = defaultState, action: Object) {
  switch (action.type) {
    case actionTypes.PROXY_ADD_WS_CLIENT:
      return Object.assign({}, state, {
        wsClient: action.client,
      });

    case actionTypes.PROXY_REMOVE_WS_CLIENT:
      return Object.assign({}, state, {
        wsClient: null,
      });

    case actionTypes.PROXY_SET_DEV_STATUS:
      return Object.assign({}, state, {
        status: action.status,
      });

    default:
      return state;
  }
}
