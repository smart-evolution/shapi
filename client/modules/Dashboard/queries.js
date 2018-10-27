import _ from 'lodash';

/* eslint-disable import/prefer-default-export */
export const getTicks = (times, values) => _.map(values, (value, index) => ({
  time: new Date(times[index]),
  value,
}));
/* eslint-enable import/prefer-default-export */
