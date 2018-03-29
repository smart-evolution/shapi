import React from 'react';
import { render } from 'react-dom';
import { createStore, applyMiddleware, combineReducers } from 'redux';
import { Provider } from 'react-redux';
import createSagaMiddleware from 'redux-saga';

import Dashboard from './modules/Dashboard';
import dashboardReducer from './modules/Dashboard/reducers';
import dashboardSagas from './modules/Dashboard/sagas';

const dashboardContainer = document.querySelector('.js-dashboard');

const sagaMiddleware = createSagaMiddleware();

const store = createStore(
  dashboardReducer,
  applyMiddleware(sagaMiddleware)
);

sagaMiddleware.run(dashboardSagas);

if (dashboardContainer) {
  render(
    <Provider store={store}>
      <Dashboard />
    </Provider>,
    dashboardContainer
  );
}


