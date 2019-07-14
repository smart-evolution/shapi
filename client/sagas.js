import { fork, takeEvery } from 'redux-saga/effects';
import * as applicationSagas from './modules/Application/sagas';
import * as applicationActionTypes from './modules/Application/actionTypes';
import * as agentsSagas from './models/agents/sagas';
import * as agentsActionTypes from './models/agents/actionTypes';
import * as agentConfigsSagas from './models/agentConfigs/sagas';
import * as agentConfigsActionTypes from './models/agentConfigs/actionTypes';
import * as proxySagas from './models/proxy/sagas';
import * as proxyActionTypes from './models/proxy/actionTypes';

function* root() {
  yield [
    takeEvery(
      applicationActionTypes.MOUNT,
      applicationSagas.onApplicationMount
    ),
    fork(agentsSagas.subscribeOnFetchAgents),
    fork(agentsSagas.onFetchAlerts),
    takeEvery(agentsActionTypes.SNIFF_AGENTS, agentsSagas.onSniffAgents),
    takeEvery(
      agentConfigsActionTypes.FETCH_AGENT_CONFIGS,
      agentConfigsSagas.onFetchAgentConfigs
    ),
    takeEvery(
      agentConfigsActionTypes.COMMIT_AGENT_CONFIG,
      agentConfigsSagas.onCommitAgentConfig
    ),
    takeEvery(agentsActionTypes.TOGGLE_ALERTS, agentsSagas.onToggleAlerts),
    takeEvery(agentsActionTypes.FETCH_ALERTS, agentsSagas.onFetchAlerts),
    takeEvery(agentsActionTypes.SEND_ALERT, agentsSagas.onSendAlert),
    takeEvery(agentsActionTypes.TOGGLE_TYPE2, agentsSagas.onToggleType2),
    takeEvery(
      proxyActionTypes.PROXY_CREATE_WS_CLIENT,
      proxySagas.onCreateWebSocketClient
    ),
    takeEvery(proxyActionTypes.PROXY_SEND_MESSAGE, proxySagas.onSendMessage),
  ];
}

export default root;
