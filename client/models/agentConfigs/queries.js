// @flow
import _ from 'lodash';

export const getAgentConfigByAgentId = (agentConfigs, agentID) => _.find(agentConfigs, { agentId: agentID });
