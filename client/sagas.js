import { fork, takeEvery } from 'redux-saga/effects';
import * as agentsSagas from './models/agents/sagas';
import * as agentsActionTypes from './models/agents/actionTypes';

function* root() {
  yield [
    fork(agentsSagas.fetchData),
    fork(agentsSagas.fetchAlerts),
    takeEvery(agentsActionTypes.TOGGLE_ALERTS, agentsSagas.toggleAlerts),
    takeEvery(agentsActionTypes.FETCH_ALERTS, agentsSagas.fetchAlerts),
    takeEvery(agentsActionTypes.SEND_ALERT, agentsSagas.sendAlert),
    takeEvery(agentsActionTypes.TOGGLE_TYPE2, agentsSagas.toggleType2),
  ];
}

export default root;
