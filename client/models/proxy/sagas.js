// @flow
import { put, select } from 'redux-saga/effects';
import * as agentTypes from 'models/agents/types';
import * as selectors from './selectors';
import * as actions from './actions';

export function* createWebSockerClient(agent: agentTypes.Agent) {
  const client: WebSocket = new WebSocket("ws://localhost:3222/sapi");

  client.onopen = function() {};

  client.onmessage = function(event) {
    const m = JSON.parse(event.data);
    console.debug("Received message", m.message);
  };

  client.onerror = function(event) {
    console.debug(event)
  };

  yield put(actions.addWebSocketClient(agent, client));
}

export function* sendMessage({ agent, message }) {
  const client: WebSocket = yield select(selectors.getWsClient);
  const { left, top } = message;
  yield client.send(JSON.stringify({
    id: agent.id,
    left,
    top,
  }));
}
