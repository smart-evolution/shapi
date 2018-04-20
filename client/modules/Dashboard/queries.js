import _ from 'lodash';

export const getTicks = (times, values) => {
  return _.map(values, (value, index) => ({
    time: new Date(times[index]),
    value,
  }));
};
