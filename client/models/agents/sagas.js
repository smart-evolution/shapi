import _ from 'lodash';
import { delay } from 'redux-saga';
import { put, call } from 'redux-saga/effects';
import * as actions from './actions';

function getData() {
  return fetch('/api/agents')
    .then(response => {
      if (!response.ok) {
        throw new Error(`Fetching data error: ${response.statusText}`);
      }

      if (response.status === 204) {
        return {
          _embedded: {
            agents: [],
          },
        };
      }

      return response.json();
    })
    .catch(e => e);
}

function callSendAlert() {
  return fetch('/api/sendalert', { method: 'POST' })
    .then(response => response.json())
    .catch(() => 'Send alert failed');
}

export function* fetchData() {
  while (true) {
    const data = yield call(getData);

    if (_.isEmpty(data)) {
      yield put(actions.fetchDataFail('Fetched data empty'));
      return;
    }

    const agents = data._embedded.agents;

    if (_.isArray(agents)) {
      yield put(actions.fetchDataSuccess(agents));
    } else {
      yield put(actions.fetchDataFail('Fetched data is not array of agents'));
    }

    yield delay(5000);
  }
}

export function* sendAlert() {
  yield call(callSendAlert);
}

function callToggleAlerts() {
  return fetch('/api/alerts', { method: 'POST' })
    .then(response => response.json())
    .catch(() => 'Toggling alerts failed');
}

export function* toggleAlerts() {
  const data = yield call(callToggleAlerts);

  if (_.isObject(data)) {
    const isAlerts = data.isAlerts === 'true';

    yield put(actions.setAlerts(isAlerts));
  }
}

function callAlerts() {
  return fetch('/api/alerts')
    .then(response => response.json())
    .catch(() => 'Toggling alerts failed');
}

export function* fetchAlerts() {
  const data = yield call(callAlerts);

  if (_.isObject(data)) {
    const isAlerts = data.isAlerts === 'true';

    yield put(actions.setAlerts(isAlerts));
  } else {
    // @TODO: Figure out how to handle failed actions on the UI side
  }
}

function callToggleType2(agentID) {
  return fetch(`/api/agents/${agentID}`, { method: 'POST' })
    .then(response => response.json())
    .catch(() => 'Toggling Type2 failed');
}

export function* toggleType2({ agentID }) {
  const data = yield call(callToggleType2, agentID);

  if (_.isObject(data)) {
    const isAlerts = data.isAlerts === 'true';

    yield put(actions.setAlerts(isAlerts));
  }
}
