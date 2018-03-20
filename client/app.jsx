import React from 'react';
import { render } from 'react-dom';
import { createStore, applyMiddleware } from 'redux';
import { Provider } from 'react-redux';
import createSagaMiddleware from 'redux-saga';
import TemperatureChart from './modules/TemperatureChart';
import temperatureChartReducer from './modules/TemperatureChart/reducers';
import temperatureChartSagas from './modules/TemperatureChart/sagas';

const temperatureContainer = document.querySelector('.js-temperature');

const sagaMiddleware = createSagaMiddleware();

if (temperatureContainer) {
    const store = createStore(
        temperatureChartReducer,
        applyMiddleware(sagaMiddleware)
    );

    sagaMiddleware.run(temperatureChartSagas);

    render(
        <Provider store={store}>
            <TemperatureChart />
        </Provider>,
        temperatureContainer
    );
}

