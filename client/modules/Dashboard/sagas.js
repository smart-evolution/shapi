import _ from 'lodash';
import { delay } from 'redux-saga';
import { put, call, fork, takeEvery } from 'redux-saga/effects';
import * as actions from './actions';

function getData(agentId) {
  return fetch(`/api/agents/${agentId}`)
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
    const path = window.location.pathname;
    const hasAgentId = _.first(path.match(/\/[a-z0-9]*$/));

    if (!hasAgentId) {
      return put(actions.fetchDataFail('No agentId specified in agent API URL'));
    }

    const agentId = hasAgentId.replace('/', '');
    const data = yield call(getData, agentId);

    if (_.isArray(data)) {
      const { time, temperature, presence, gas, sound } = _.first(data).data;

      yield put(actions.fetchDataSuccess(time, temperature, presence, gas, sound));
    } else {
      yield put(actions.fetchDataFail(data));
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
