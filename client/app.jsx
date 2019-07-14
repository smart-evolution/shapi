// @flow
import React from 'react';
import { render } from 'react-dom';
import { createStore, applyMiddleware } from 'redux';
import { Provider } from 'react-redux';
import { Router, Route } from 'react-router';
import { createBrowserHistory } from 'history';
import createSagaMiddleware from 'redux-saga';
import Application from './modules/Application';
import AgentsStatus from './modules/AgentsStatus';
import Dashboard from './modules/Dashboard';
import AgentEdit from './modules/AgentEdit';
import sagas from './sagas';
import reducers from './reducers';

const appContainer = document.querySelector('.js-app');
const sagaMiddleware = createSagaMiddleware();

const store = createStore(reducers, applyMiddleware(sagaMiddleware));

sagaMiddleware.run(sagas);

if (appContainer) {
  render(
    <Provider store={store}>
      <Router history={createBrowserHistory({})}>
        <Application>
          <Route exact path="/" component={AgentsStatus} />
          <Route exact path="/agent/:agent" component={Dashboard} />
          <Route path="/agent/:agent/edit" component={AgentEdit} />
        </Application>
      </Router>
    </Provider>,
    appContainer
  );
}
