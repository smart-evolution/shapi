import _ from 'lodash';

/* eslint-disable import/prefer-default-export */
export const getAgentConfig = (state, agentID) => {
  const config = state.agentConfigs.agentConfigs[agentID];

  if (_.isEmpty(config)) {
    return {
      temperature: 0,
    }
  }

  return config;
};
/* eslint-enable import/prefer-default-export */
