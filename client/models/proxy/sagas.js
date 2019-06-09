// @flow
import _ from 'lodash';
import { eventChannel } from 'redux-saga';
import { put, call, select, take } from 'redux-saga/effects';
import * as agentTypes from 'client/models/agents/types';
import * as alertsActions from 'client/models/alerts/actions';
import * as alertsConstants from 'client/models/alerts/constants';
import * as selectors from './selectors';
import * as actions from './actions';

/* eslint-disable require-yield */
function* createChannel(client: WebSocket) {
  return eventChannel(emit => {
    client.onmessage = (message) => emit(message.data);
    return () => {
      client.close();
    };
  });
}
/* eslint-enable require-yield */

export function* createWebSocketClient(agent: agentTypes.Agent): Iterable<any> {
  const { host } = window.location;

  const client: WebSocket = new WebSocket(`ws://${host}/sapi`);
  yield put(actions.addWebSocketClient(agent, client));

  const channel = yield call(createChannel, client);

  while (true) {
    const data = yield take(channel);

    if (data instanceof String) {
      const { type, message } = JSON.parse(data.slice(1,-1).replace(/\\"/g, '"'));

      if (type === 'connected') {
        yield put(alertsActions.addAlert(message, alertsConstants.ALERT_TYPE_INFO));
        yield put(actions.setDevStatus(true));
      } else if (type === 'error') {
        yield put(alertsActions.addAlert(message, alertsConstants.ALERT_TYPE_ERROR));
        yield put(actions.setDevStatus(false));
      }
    }
  }
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
        ip: agent.ip,
        left,
        top,
      })
    );
  }
}
