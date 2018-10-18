import _ from 'lodash';

export default (times, values) => _.map(values, (value, index) => ({
  time: new Date(times[index]),
  value,
}));
