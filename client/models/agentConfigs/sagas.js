// @flow
import _ from 'lodash';
import { put, call } from 'redux-saga/effects';
import * as alertsActions from 'client/models/alerts/actions';
import * as alertsConstants from 'client/models/alerts/constants';
import * as agentTypes from 'client/models/agents/types';
import * as actions from './actions';
import * as types from './types';
import * as constants from './constants';

function callFetchAgentConfigs(agentID: agentTypes.AgentID) {
  return fetch(`${constants.AGENT_CONFIG_ENDPOINT}/${agentID}`)
    .then(response => {
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

export function* onFetchAgentConfigs({
  agentID,
}: {
  agentID: string,
}): Iterable<any> {
  const data = yield call(callFetchAgentConfigs, agentID);

  if (data !== undefined) {
    const agentConfigs = data._embedded.configs;
    yield put(actions.loadAgentConfigs(agentConfigs));
  } else {
    yield put(
      alertsActions.addAlert(
        'Fetching agent config failed',
        alertsConstants.ALERT_TYPE_ERROR
      )
    );
  }
}

function callCommitAgentConfig(
  agentID: agentTypes.AgentID,
  config: types.AgentConfig
) {
  return fetch(`${constants.AGENT_CONFIG_ENDPOINT}/${agentID}`, {
    method: 'POST',
    body: JSON.stringify(config),
  })
    .then(response => response.json())
    .catch(() => 'Updating agent config failed');
}

export function* onCommitAgentConfig({
  agentID,
  config,
}: {
  agentID: agentTypes.AgentID,
  config: types.AgentConfig,
}): Iterable<any> {
  const response = yield call(callCommitAgentConfig, agentID, config);

  if (!_.isEmpty(response)) {
    yield put(
      alertsActions.addAlert(
        'Updated agent config successfully',
        alertsConstants.ALERT_TYPE_INFO
      )
    );
  } else {
    yield put(
      alertsActions.addAlert(
        'Updating agent config failed',
        alertsConstants.ALERT_TYPE_ERROR
      )
    );
  }
}
