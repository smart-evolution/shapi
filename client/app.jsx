import React from 'react';
import { render } from 'react-dom';
import { createStore, applyMiddleware } from 'redux';
import { Provider } from 'react-redux';
import createSagaMiddleware from 'redux-saga';

import AgentsList from './modules/AgentsList';
import agentsListReducer from './modules/AgentsList/reducers';
import agentsListSagas from './modules/AgentsList/sagas';

import Dashboard from './modules/Dashboard';
import dashboardReducer from './modules/Dashboard/reducers';
import dashboardSagas from './modules/Dashboard/sagas';

const agentsListContainer = document.querySelector('.js-agents-list');
const agentContainer = document.querySelector('.js-agent');

const sagaMiddleware = createSagaMiddleware();

if (agentsListContainer) {
  const store = createStore(
    agentsListReducer,
    applyMiddleware(sagaMiddleware)
  );

  sagaMiddleware.run(agentsListSagas);

  render(
    <Provider store={store}>
      <AgentsList />
    </Provider>,
    agentsListContainer
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
