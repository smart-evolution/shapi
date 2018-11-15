export const getAgents = (state) => state.agents.agents;

export const getTemperatures = (agent) => {
  const { data } = agent;
  return data.temperature;
};

export const getMotion = (agent) => {
  const { data } = agent;
  return data.presence;
};

export const getGas = (agent) => {
  const { data } = agent;
  return data.gas;
};

export const getTimes = (agent) => {
  const { data } = agent;
  return data.time;
};
