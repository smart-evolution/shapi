// @flow
import _ from 'lodash';
import { all, put, take, select } from 'redux-saga/effects';
import * as agentSelectors from 'client/models/agents/selectors'
import * as agentActions from 'client/models/agents/actions';
import * as agentConfigActions from 'client/models/agentConfigs/actions';
import * as agentActionTypes from 'client/models/agents/actionTypes';
import * as actions from './actions';

/* eslint-disable import/prefer-default-export */
export function* mount(): Iterable<any> {
  yield all([
    put(agentActions.fetchData()),
  ]);
  yield take([agentActionTypes.DATA_FETCH_SUCCESS]);

  const agents = yield select(agentSelectors.getAgents);
  yield all(
    _.map(agents, agent => put(agentConfigActions.fetchData(agent.id)))
  );

  yield put(actions.loaded());
}
/* eslint-enable import/prefer-default-export */
