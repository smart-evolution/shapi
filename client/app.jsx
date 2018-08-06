import React from 'react';
import { render } from 'react-dom';
import { createStore, applyMiddleware, combineReducers } from 'redux';
import { Provider } from 'react-redux';
import createSagaMiddleware from 'redux-saga';

import Dashboard from './modules/Dashboard';
import dashboardReducer from './modules/Dashboard/reducers';
import dashboardSagas from './modules/Dashboard/sagas';

const agentContainer = document.querySelector('.js-agent');

const sagaMiddleware = createSagaMiddleware();

const store = createStore(
  dashboardReducer,
  applyMiddleware(sagaMiddleware)
);

if (agentContainer) {
  sagaMiddleware.run(dashboardSagas);

  render(
    <Provider store={store}>
      <Dashboard />
    </Provider>,
    agentContainer
  );
}


