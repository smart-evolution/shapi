import _ from 'lodash';
import * as selectors from './selectors';

export const getTemperature = (agent) => {
  const temperatures = selectors.getTemperatures(agent);
  return _.first(temperatures);
};

export const isMotion = (agent) => {
  const presence = selectors.getMotion(agent);
  return _.reduce(presence, (acc, val) => acc || Number(val), 0);
};

export const isGas = (agent) => {
  const gas = selectors.getGas(agent);
  return _.reduce(gas, (acc, val) => acc || Number(val), 0);
};
