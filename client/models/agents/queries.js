import _ from 'lodash';
import * as selectors from './selectors';

export const getTemperatures = agent => {
  const { data } = agent;
  return data.temperature;
};

export const getTemperature = agent => {
  const temperatures = getTemperatures(agent);
  return _.first(temperatures);
};

export const getMotion = agent => {
  const { data } = agent;
  return data.presence;
};

export const getGas = agent => {
  const { data } = agent;
  return data.gas;
};

export const getTimes = agent => {
  const { data } = agent;
  return data.time;
};

export const getTicks = (times, values) =>
  _.map(values, (value, index) => ({
    time: new Date(times[index]),
    value,
  }));

export const getAgentById = (state, id) => {
  const agents = selectors.getAgents(state);
  return _.find(agents, { id });
};

export const isMotion = agent => {
  const presence = getMotion(agent);
  return _.reduce(presence, (acc, val) => acc || Number(val), 0);
};

export const isGas = agent => {
  const gas = getGas(agent);
  return _.reduce(gas, (acc, val) => acc || Number(val), 0);
};
