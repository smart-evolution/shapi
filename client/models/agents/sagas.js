import _ from 'lodash';
import { delay } from 'redux-saga';
import { put, call, select } from 'redux-saga/effects';
import * as actions from './actions';
import * as selectors from './selectors';
import * as alertsActions from '../alerts/actions';
import * as alertsConstants from '../alerts/constants';

export function callFetchAgents(period) {
  return fetch(`/api/agents?peroid=${period}`)
    .then(response => {
      if (!response.ok) {
        throw new Error(`Fetching data error: ${response.statusText}`);
      }

      if (response.status === 204) {
        return {
          _embedded: {
            agents: [],
          },
        };
      }

      return response.json();
    })
    .catch(e => e);
}

export function* onFetchAgents() {
  const period = yield select(selectors.getPeriod);
  const data = yield call(callFetchAgents, period);

  if (_.isEmpty(data)) {
    yield put(actions.fetchAgentsError('Fetched data empty'));
    return;
  }

  const agents = data._embedded.agents;

  if (_.isArray(agents)) {
    yield put(actions.loadAgents(agents));
  } else {
    yield put(actions.fetchAgentsError('Fetched data is not array of agents'));
  }
}

export function* subscribeOnFetchAgents() {
  while (true) {
    yield onFetchAgents();
    yield delay(5000);
  }
}

function callSendAlert() {
  return fetch('/api/sendalert', { method: 'POST' })
    .then(response => response.json())
    .catch(() => 'Send alert failed');
}

export function* onSendAlert() {
  yield call(callSendAlert);
}

function callToggleAlerts() {
  return fetch('/api/alerts', { method: 'POST' })
    .then(response => response.json())
    .catch(() => 'Toggling alerts failed');
}

export function* onToggleAlerts() {
  const data = yield call(callToggleAlerts);

  if (_.isObject(data)) {
    const isAlerts = data.isAlerts === 'true';

    yield put(actions.setAlerts(isAlerts));
  }
}

function callAlerts() {
  return fetch('/api/alerts')
    .then(response => response.json())
    .catch(() => 'Toggling alerts failed');
}

export function* onFetchAlerts() {
  const data = yield call(callAlerts);

  if (_.isObject(data)) {
    const isAlerts = data.isAlerts === 'true';

    yield put(actions.setAlerts(isAlerts));
  } else {
    yield put(
      alertsActions.addAlert(
        'Alerts not fetched properly',
        alertsConstants.ALERT_TYPE_ERROR
      )
    );
  }
}

function callToggleType2(agentID) {
  return fetch(`/api/agents/${agentID}`, { method: 'POST' })
    .then(response => response.json())
    .catch(() => 'Toggling Type2 failed');
}

export function* onToggleType2({ agentID }) {
  const data = yield call(callToggleType2, agentID);

  if (_.isObject(data)) {
    const isAlerts = data.isAlerts === 'true';

    yield put(actions.setAlerts(isAlerts));
  }
}

function callSniffAgents() {
  return fetch('/api/sniffagents', { method: 'POST' })
    .then(response => response.json())
    .catch(() => 'Toggling alerts failed');
}

export function* onSniffAgents() {
  yield put(actions.fetchAgents());

  const data = yield call(callSniffAgents);

  if (!_.isEmpty(data)) {
    yield put(
      alertsActions.addAlert(
        'Agents sniffing in progress',
        alertsConstants.ALERT_TYPE_INFO
      )
    );
    return;
  }
  yield put(
    alertsActions.addAlert(
      'Agent sniffing failed',
      alertsConstants.ALERT_TYPE_ERROR
    )
  );
}
