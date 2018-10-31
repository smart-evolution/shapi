import _ from 'lodash';
import { delay } from 'redux-saga';
import { put, call, fork } from 'redux-saga/effects';
import * as actions from './actions';

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

function* root() {
  yield [
    fork(fetchData),
  ];
}

export default root;
