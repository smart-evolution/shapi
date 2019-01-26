import _ from 'lodash';
import { put, call } from 'redux-saga/effects';
import * as actions from './actions';

function getData(agentID) {
  return fetch(`/api/agentsConfig/${agentID}`)
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

export function* fetchData({ agentID }) {
    yield call(getData, agentID);
}

function callUpdateData(agentID, data) {
  return fetch(`/api/agentsConfig/${agentID}`, {
    method: 'POST',
    body: JSON.stringify(data),
  })
  .then(response => response.json())
  .catch(() => 'Send alert failed');
}

export function* updateData({ agentID, data }) {
  yield call(callUpdateData, agentID, data);
}
