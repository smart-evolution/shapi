import React from 'react';
import { render } from 'react-dom';
import { createStore, applyMiddleware } from 'redux';
import { Provider } from 'react-redux';
import { Router, Route } from 'react-router';
import { createBrowserHistory } from 'history'
import createSagaMiddleware from 'redux-saga';
import AgentsStatus from './modules/AgentsStatus';
import agentsStatusReducer from './modules/AgentsStatus/reducers';
import agentsStatusSagas from './modules/AgentsStatus/sagas';
import Dashboard from './modules/Dashboard';
import dashboardReducer from './modules/Dashboard/reducers';
import dashboardSagas from './modules/Dashboard/sagas';

const appContainer = document.querySelector('.js-app');
const agentsStatusContainer = document.querySelector('.js-agents-status');
const agentContainer = document.querySelector('.js-agent');
const sagaMiddleware = createSagaMiddleware();

let store;

if (agentsStatusContainer) {
  store = createStore(
    agentsStatusReducer,
    applyMiddleware(sagaMiddleware)
  );

  sagaMiddleware.run(agentsStatusSagas);
}

if (agentContainer) {
  store = createStore(
    dashboardReducer,
    applyMiddleware(sagaMiddleware)
  );

  sagaMiddleware.run(dashboardSagas);
}

if (appContainer) {
  render(
    <Provider store={store}>
      <Router history={createBrowserHistory({})}>
        <div>
          <Route path='/' component={AgentsStatus} />
          <Route path='/agent/:agent' component={Dashboard} />
        </div>
      </Router>
    </Provider>,
    appContainer
  );
}


