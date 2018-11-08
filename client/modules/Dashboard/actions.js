import * as actionTypes from './actionTypes';

export const fetchDataSuccess = (times, temperatures, motions, gas, sounds) => ({
  type: actionTypes.DATA_FETCH_SUCCESS,
  times,
  temperatures,
  motions,
  gas,
  sounds,
});

export const fetchDataFail = error => ({
  type: actionTypes.DATA_FETCH_ERROR,
  error,
});
