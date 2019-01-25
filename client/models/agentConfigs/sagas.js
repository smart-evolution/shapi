import _ from 'lodash';
import { put, call } from 'redux-saga/effects';
import * as actions from './actions';

function getData() {
  return fetch('/api/agentsConfig')
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

export function* fetchData() {
    const data = yield call(getData);
    const agents = data._embedded.agents;

    if (_.isArray(agents)) {
      yield put(actions.fetchDataSuccess(agents));
    } else {
      yield put(actions.fetchDataFail(agents));
    }

    yield delay(5000);
}
