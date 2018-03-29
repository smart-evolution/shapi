import { delay } from 'redux-saga'
import { put, call } from 'redux-saga/effects';
import * as actions from './actions';

function detData() {
  return fetch('/api/home').then(response => response.json());
}

function* fetchData() {
  while (true) {
    const data = yield call(detData);

    yield put(actions.fetchedData(data.temperature, data.presence));
    yield delay(1000);
  }
}

export default fetchData;
