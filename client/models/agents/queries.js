// @flow
import _ from 'lodash';
import * as types from './types';

export const isOnline = (agent: types.Agent): boolean => agent.isOnline;

export const getTemperatures = (agent: types.Agent): $ReadOnlyArray<string> => {
  const { data } = agent;

  if (_.isEmpty(data) || !_.isArray(data.temperature)) {
    return [];
  }
  return data.temperature;
};

export const getTemperature = (agent: types.Agent): string => {
  const temperatures = getTemperatures(agent);
  return _.first(temperatures);
};

export const getMotion = (agent: types.Agent): $ReadOnlyArray<string> => {
  const { data } = agent;

  if (_.isEmpty(data) || !_.isArray(data.presence)) {
    return [];
  }
  return data.presence;
};

export const getGas = (agent: types.Agent): $ReadOnlyArray<string> => {
  const { data } = agent;

  if (_.isEmpty(data) || !_.isArray(data.gas)) {
    return [];
  }
  return data.gas;
};

export const getTimes = (agent: types.Agent): $ReadOnlyArray<string> => {
  const { data } = agent;

  if (_.isEmpty(data) || !_.isArray(data.time)) {
    return [];
  }
  return data.time;
};

export const getTicks = (times: Array<number>, values: number) =>
  _.map(values, (value, index) => ({
    time: new Date(times[index]),
    value,
  }));

export const isMotion = (agent: types.Agent) => {
  const presence = getMotion(agent);
  return _.reduce(presence, (acc, val) => acc || Number(val), 0);
};

export const isGas = (agent: types.Agent) => {
  const gas = getGas(agent);
  return _.reduce(gas, (acc, val) => acc || Number(val), 0);
};

export const getNoVersionedType = (agent: types.Agent) =>
  _.head(_.split(agent.type, '-'));
