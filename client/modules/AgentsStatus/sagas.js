import _ from 'lodash';
import { delay } from 'redux-saga';
import { put, call, fork, takeEvery } from 'redux-saga/effects';
import * as actions from './actions';
import * as actionTypes from './actionTypes';

function getData() {
  return fetch('/api/agents')
    .then((response) => {
      if (!response.ok) {
        throw new Error(`Fetching data error: ${response.statusText}`);
      }

      if (response.status === 204) {
        throw new Error('No data available');
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

function* fetchData() {
  while (true) {
    const agents = yield call(getData);

    if (_.isArray(agents)) {
      yield put(actions.fetchDataSuccess(agents));
    } else {
      yield put(actions.fetchDataFail(agents));
    }

    yield delay(5000);
  }
}

function* sendAlert() {
  yield call(callSendAlert);
}

function callToggleAlerts() {
  return fetch('/api/alerts', { method: 'POST' })
    .then(response => response.json())
    .catch(() => 'Toggling alerts failed');
}

function* toggleAlerts() {
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

function* fetchAlerts() {
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

function* toggleType2({ agentID }) {
  const data = yield call(callToggleType2, agentID);

  if (_.isObject(data)) {
    const isAlerts = data.isAlerts === 'true';

    yield put(actions.setAlerts(isAlerts));
  }
}

function* root() {
  yield [
    fork(fetchData),
    fork(fetchAlerts),
    takeEvery(actionTypes.TOGGLE_ALERTS, toggleAlerts),
    takeEvery(actionTypes.FETCH_ALERTS, fetchAlerts),
    takeEvery(actionTypes.SEND_ALERT, sendAlert),
    takeEvery(actionTypes.TOGGLE_TYPE2, toggleType2),
  ];
}

export default root;
