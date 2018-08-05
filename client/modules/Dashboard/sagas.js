import _ from 'lodash';
import { delay } from 'redux-saga'
import { put, call, fork, takeEvery } from 'redux-saga/effects';
import * as actions from './actions';
import * as actionTypes from './actionTypes';

function callSendAlert() {
  return fetch('/api/sendalert', { method: 'POST' })
    .then(response => response.json())
    .catch(() => "Send alert failed");
}

function* sendAlert() {
  yield call(callSendAlert);
}

function callToggleAlerts() {
  return fetch('/api/alerts', { method: 'POST' })
    .then(response => response.json())
    .catch(() => "Toggling alerts failed");
}

function* toggleAlerts() {
  const data = yield call(callToggleAlerts);

  if(_.isObject(data)) {
    const isAlerts = data.isAlerts == "true" ? true : false;

    yield put(actions.setAlerts(isAlerts));
  } else {

  }
}

function callAlerts() {
  return fetch('/api/alerts')
    .then(response => response.json())
    .catch(() => "Toggling alerts failed");
}

function* fetchAlerts() {
  const data = yield call(callAlerts);

  if(_.isObject(data)) {
    const isAlerts = data.isAlerts == "true" ? true : false;

    yield put(actions.setAlerts(isAlerts));
  } else {
    // @TODO: Figure out how to handle failed actions on the UI side
  }
}

function getData(agentId) {
  return fetch(`/api/home/${agentId}`)
    .then(response => response.json())
    .catch(() => "Fetching data failed");
}

function* fetchData() {
  while (true) {
    const path = window.location.pathname;
    const hasAgentId = _.first(path.match(/\/[a-z0-9]*$/));

    if (!hasAgentId) {
      return put(actions.fetchDataFail('No agentId specified in agent API URL'));
    }

    const agentId = hasAgentId.replace('/', '');
    const data = yield call(getData, agentId);

    if(_.isArray(data)) {
      const { time, temperature, presence, gas, sound } = _.first(data).data;

      yield put(actions.fetchDataSuccess(time, temperature, presence, gas, sound));
    } else {
      yield put(actions.fetchDataFail(data));
    }

    yield delay(3000);
  }
}

function* root() {
  yield [
    fork(fetchData),
    fork(fetchAlerts),
    takeEvery(actionTypes.TOGGLE_ALERTS, toggleAlerts),
    takeEvery(actionTypes.FETCH_ALERTS, fetchAlerts),
    takeEvery(actionTypes.SEND_ALERT, sendAlert),
  ];
}

export default root;
