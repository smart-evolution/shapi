import * as actionTypes from './actionTypes';

export const fetchTemperature = () => ({
   type: actionTypes.TEMPERATURE_CHART_FETCH,
})

export const fetchedTemperature = (temperature) => ({
    type: actionTypes.TEMPERATURE_CHART_FETCHED,
    temperature,
});