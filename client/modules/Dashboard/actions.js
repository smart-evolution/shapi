import * as actionTypes from './actionTypes';

export const fetchedData = (time, temperature, motions) => ({
  type: actionTypes.DATA_FETCHED,
  time,
  temperature,
  motions,
});

export const fetchDataFail = (error) => ({
  type: actionTypes.DATA_FETCH_ERROR,
  error,
});
