import { fork, takeEvery } from 'redux-saga/effects';
import * as agentsSagas from './models/agents/sagas';
import * as agentsActionTypes from './models/agents/actionTypes';
import * as agentConfigsSagas from './models/agentConfigs/sagas';
import * as agentConfigsActionTypes from './models/agentConfigs/actionTypes';
import * as proxySagas from './models/proxy/sagas';
import * as proxyActionTypes from './models/proxy/actionTypes';

function* root() {
  yield [
    fork(agentsSagas.fetchData),
    fork(agentsSagas.fetchAlerts),
    takeEvery(agentConfigsActionTypes.FETCH_DATA, agentConfigsSagas.fetchData),
    takeEvery(
      agentConfigsActionTypes.POST_AGENT_CONFIG,
      agentConfigsSagas.updateData
    ),
    takeEvery(agentsActionTypes.TOGGLE_ALERTS, agentsSagas.toggleAlerts),
    takeEvery(agentsActionTypes.FETCH_ALERTS, agentsSagas.fetchAlerts),
    takeEvery(agentsActionTypes.SEND_ALERT, agentsSagas.sendAlert),
    takeEvery(agentsActionTypes.TOGGLE_TYPE2, agentsSagas.toggleType2),
    takeEvery(proxyActionTypes.PROXY_CREATE_WS_CLIENT, proxySagas.createWebSockerClient),
    takeEvery(proxyActionTypes.PROXY_SEND_MESSAGE, proxySagas.sendMessage),
  ];
}

export default root;
