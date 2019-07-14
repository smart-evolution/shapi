// @flow
import { all, put, take } from 'redux-saga/effects';
import * as agentActions from 'client/models/agents/actions';
import * as agentActionTypes from 'client/models/agents/actionTypes';
import * as agentConfigActions from 'client/models/agentConfigs/actions';
import * as agentConfigActionTypes from 'client/models/agentConfigs/actionTypes';
import * as actions from './actions';

/* eslint-disable import/prefer-default-export */
export function* mount(): Iterable<any> {
  yield all([put(agentActions.fetchData())]);
  yield take([agentActionTypes.DATA_FETCH_SUCCESS]);

  yield put(agentConfigActions.fetchData());
  yield take([agentConfigActionTypes.LOAD_AGENT_CONFIGS]);

  yield put(actions.loaded());
}
/* eslint-enable import/prefer-default-export */
