import React from 'react';
import { render } from 'react-dom';
import { createStore, applyMiddleware } from 'redux';
import { Provider } from 'react-redux';
import createSagaMiddleware from 'redux-saga';

import AgentsStatus from './modules/AgentsStatus';
import agentsStatusReducer from './modules/AgentsStatus/reducers';
import agentsStatusSagas from './modules/AgentsStatus/sagas';

import Dashboard from './modules/Dashboard';
import dashboardReducer from './modules/Dashboard/reducers';
import dashboardSagas from './modules/Dashboard/sagas';

const agentsStatusContainer = document.querySelector('.js-agents-status');
const agentContainer = document.querySelector('.js-agent');

const sagaMiddleware = createSagaMiddleware();

if (agentsStatusContainer) {
  const store = createStore(
    agentsStatusReducer,
    applyMiddleware(sagaMiddleware)
  );

  sagaMiddleware.run(agentsStatusSagas);

  render(
    <Provider store={store}>
      <AgentsStatus />
    </Provider>,
    agentsStatusContainer
  );
}

if (agentContainer) {
  const store = createStore(
    dashboardReducer,
    applyMiddleware(sagaMiddleware)
  );

  sagaMiddleware.run(dashboardSagas);

  render(
    <Provider store={store}>
      <Dashboard />
    </Provider>,
    agentContainer
  );
}
