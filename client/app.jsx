import React from 'react';
import { render } from 'react-dom';
import { createStore, applyMiddleware } from 'redux';
import { Provider } from 'react-redux';
import { Router, Route } from 'react-router';
import { createBrowserHistory } from 'history';
import createSagaMiddleware from 'redux-saga';
import AgentsStatus from './modules/AgentsStatus';
import Dashboard from './modules/Dashboard';
import sagas from './sagas';
import reducers from './reducers';

const appContainer = document.querySelector('.js-app');
const sagaMiddleware = createSagaMiddleware();

const store = createStore(
  reducers,
  applyMiddleware(sagaMiddleware)
);

sagaMiddleware.run(sagas);

if (appContainer) {
  render(
    <Provider store={store}>
      <Router history={createBrowserHistory({})}>
        <div>
          <Route exact path="/" component={AgentsStatus} />
          <Route path="/agent/:agent" component={Dashboard} />
        </div>
      </Router>
    </Provider>,
    appContainer
  );
}
