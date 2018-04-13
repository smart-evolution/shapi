import _ from 'lodash';
import { delay } from 'redux-saga'
import { put, call, fork, takeEvery } from 'redux-saga/effects';
import * as actions from './actions';
import * as actionTypes from './actionTypes';

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

  }
}

function getData() {
  return fetch('/api/home')
    .then(response => response.json())
    .catch(() => "Fetching data failed");
}

function* fetchData() {
  while (true) {
    const data = yield call(getData);

    if(_.isObject(data)) {
      const { time, temperature, presence } = data;

      yield put(actions.fetchDataSuccess(time, temperature, presence));
    } else {
      yield put(actions.fetchDataFail(data));
    }

    yield delay(1000);
  }
}

function* root() {
  yield [
    fork(fetchData),
    fork(fetchAlerts),
    takeEvery(actionTypes.TOGGLE_ALERTS, toggleAlerts),
    takeEvery(actionTypes.FETCH_ALERTS, fetchAlerts),
  ];
}

export default root;
