export const getTemperatures = (agent) => {
  const { data } = agent;
  return data.temperature;
};

export const getGas = (agent) => {
  const { data } = agent;
  return data.gas;
};

export const getMotion = (agent) => {
  const { data } = agent;
  return data.presence;
};
