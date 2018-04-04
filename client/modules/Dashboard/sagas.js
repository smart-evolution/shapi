import _ from 'lodash';
import { delay } from 'redux-saga'
import { put, call } from 'redux-saga/effects';
import * as actions from './actions';

function getData() {
  return fetch('/api/home')
    .then(response => response.json())
    .catch(() => "Fetching data failed");
}

function* fetchData() {
  while (true) {
    const data = yield call(getData);

    if(_.isObject(data)) {
      yield put(actions.fetchedData(data.time, data.temperature, data.presence));
    } else {
      yield put(actions.fetchDataFail(data));
    }

    yield delay(1000);
  }
}

export default fetchData;
