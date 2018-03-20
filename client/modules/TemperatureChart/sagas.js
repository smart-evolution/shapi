import { delay } from 'redux-saga'
import { put, call } from 'redux-saga/effects';
import * as actions from './actions';

function getTmp() {
    return fetch('/api/home')
        .then(function(response) {
            return response.json();
        });
}

function* fetchTemperature() {
    while (true) {
        const data = yield call(getTmp);

        yield put(actions.fetchedTemperature(data.temperature));
        yield delay(1000);
    }
}

export default fetchTemperature;
