import * as actionTypes from './actionTypes';

export const fetchedData = (timestamp, temperature, presence) => ({
  type: actionTypes.DATA_FETCHED,
  timestamp,
  temperature,
  presence,
});

export const fetchDataFail = (error) => ({
  type: actionTypes.DATA_FETCH_ERROR,
  error,
});
