// @flow
import { all, put, take } from 'redux-saga/effects';
import * as agentActions from 'client/models/agents/actions';
import * as agentActionTypes from 'client/models/agents/actionTypes';
import * as agentConfigActions from 'client/models/agentConfigs/actions';
import * as agentConfigActionTypes from 'client/models/agentConfigs/actionTypes';
import * as actions from './actions';

/* eslint-disable import/prefer-default-export */
export function* onApplicationMount(): Iterable<any> {
  yield all([put(agentActions.fetchAgents())]);
  yield take([agentActionTypes.LOAD_AGENTS]);

  yield put(agentConfigActions.fetchAgentConfig());
  yield take([agentConfigActionTypes.LOAD_AGENT_CONFIGS]);

  yield put(actions.loaded());
}
/* eslint-enable import/prefer-default-export */
