import * as actionTypes from './actionTypes';

export const fetchDataSuccess = (agents) => ({
  type: actionTypes.DATA_FETCH_SUCCESS,
  agents,
});

export const fetchDataFail = error => ({
  type: actionTypes.DATA_FETCH_ERROR,
  error,
});
