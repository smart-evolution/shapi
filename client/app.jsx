import React from 'react';
import { render } from 'react-dom';
import { createStore, applyMiddleware } from 'redux';
import { Provider } from 'react-redux';
import createSagaMiddleware from 'redux-saga';
import TemperaturePanel from './modules/TemperaturePanel';
import temperaturePanelReducer from './modules/TemperaturePanel/reducers';
import temperaturePanelSagas from './modules/TemperaturePanel/sagas';

const temperatureContainer = document.querySelector('.js-temperature');

const sagaMiddleware = createSagaMiddleware();

if (temperatureContainer) {
  const store = createStore(
    temperaturePanelReducer,
    applyMiddleware(sagaMiddleware)
  );

  sagaMiddleware.run(temperaturePanelSagas);

  render(
    <Provider store={store}>
      <TemperaturePanel />
    </Provider>,
    temperatureContainer
  );
}
