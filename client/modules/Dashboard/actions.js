import * as actionTypes from './actionTypes';

export const fetchDataSuccess = (times, temperatures, motions) => ({
  type: actionTypes.DATA_FETCH_SUCCESS,
  times,
  temperatures,
  motions,
});

export const fetchDataFail = (error) => ({
  type: actionTypes.DATA_FETCH_ERROR,
  error,
});
