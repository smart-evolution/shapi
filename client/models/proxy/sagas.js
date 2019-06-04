// @flow
import { put, select } from 'redux-saga/effects';
import * as agentTypes from 'models/agents/types';
import * as selectors from './selectors';
import * as actions from './actions';

export function* createWebSockerClient(agent: agentTypes.Agent) {
  const { host } = window.location;

  const client: WebSocket = new WebSocket(`ws://${host}/sapi`);
  yield put(actions.addWebSocketClient(agent, client));
}

export function* sendMessage({ agent, message }) {
  const client: WebSocket = yield select(selectors.getWsClient);
  const { left, top } = message;
  yield client.send(
    JSON.stringify({
      id: agent.id,
      left,
      top,
    })
  );
}
