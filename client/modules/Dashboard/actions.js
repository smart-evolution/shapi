import * as actionTypes from './actionTypes';

export const fetchedData = (temperature, presence) => ({
  type: actionTypes.DATA_FETCHED,
  temperature,
  presence,
});
