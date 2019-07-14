// @flow
import { eventChannel } from 'redux-saga';
import { put, call, select, take } from 'redux-saga/effects';
import * as agentTypes from 'client/models/agents/types';
import * as alertsActions from 'client/models/alerts/actions';
import * as alertsConstants from 'client/models/alerts/constants';
import * as selectors from './selectors';
import * as actions from './actions';
import * as constants from './constants';
import * as types from './types';

/* eslint-disable require-yield */
function* createChannel(client: WebSocket) {
  return eventChannel(emit => {
    client.onmessage = message => emit(message.data);
    client.onopen = () => emit('opened');
    return () => {
      client.close();
    };
  });
}
/* eslint-enable require-yield */

export function* onCreateWebSocketClient({
  agent,
}: {
  agent: agentTypes.Agent,
}): Iterable<any> {
  const { host } = window.location;

  const client: WebSocket = new WebSocket(`ws://${host}/sapi`);
  yield put(actions.addWebSocketClient(agent, client));

  const channel = yield call(createChannel, client);

  while (true) {
    const data = yield take(channel);

    if (typeof data === 'string') {
      if (data === 'opened') {
        yield put(
          actions.sendMessage(agent, {
            left: 25,
            top: 25,
            flag: constants.FLAG_CONNECT,
          })
        );
      } else {
        const { type, message } = JSON.parse(
          data.slice(1, -1).replace(/\\"/g, '"')
        );

        if (type === 'connected') {
          yield put(
            alertsActions.addAlert(message, alertsConstants.ALERT_TYPE_INFO)
          );
          yield put(actions.setDevStatus(constants.STATUS_CONNECTED));
        } else if (type === 'disconnect') {
          yield put(actions.removeWebSocketClient());

          yield put(
            alertsActions.addAlert(message, alertsConstants.ALERT_TYPE_INFO)
          );
          yield put(actions.setDevStatus(constants.STATUS_DISCONNECTED));
        } else if (type === 'error') {
          yield put(
            alertsActions.addAlert(message, alertsConstants.ALERT_TYPE_ERROR)
          );
          yield put(actions.setDevStatus(constants.STATUS_DISCONNECTED));
        }
      }
    }
  }
}

export function* onSendMessage({
  agent,
  message,
}: {
  agent: agentTypes.Agent,
  message: types.Message,
}): Iterable<any> {
  const client: ?WebSocket = yield select(selectors.getWsClient);
  const { left, top, flag } = message;
  const isConnected = yield select(selectors.getIsDevConnected);

  if (!isConnected) {
    yield put(actions.setDevStatus(constants.STATUS_PENDING));
  }

  if (client instanceof WebSocket) {
    yield client.send(
      JSON.stringify({
        id: agent.id,
        ip: agent.ip,
        left,
        top,
        flag,
      })
    );
  }
}
