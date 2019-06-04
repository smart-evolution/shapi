// @flow
import { put, select } from 'redux-saga/effects';
import * as agentTypes from 'client/models/agents/types';
import * as selectors from './selectors';
import * as actions from './actions';

export function* createWebSockerClient(agent: agentTypes.Agent): Iterable<any> {
  const { host } = window.location;

  const client: WebSocket = new WebSocket(`ws://${host}/sapi`);
  yield put(actions.addWebSocketClient(agent, client));
}

export function* sendMessage({
  agent,
  message,
}: {
  agent: agentTypes.Agent,
  message: { left: number, top: number },
}): Iterable<any> {
  const client: ?WebSocket = yield select(selectors.getWsClient);
  const { left, top } = message;

  if (client instanceof WebSocket) {
    yield client.send(
      JSON.stringify({
        id: agent.id,
        url: agent.url,
        left,
        top,
      })
    );
  }
}
