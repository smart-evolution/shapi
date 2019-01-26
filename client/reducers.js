import { combineReducers } from 'redux';
import agents from './models/agents/reducers';
import agentConfigs from './models/agentConfigs/reducers';

export default combineReducers({
  agents,
  agentConfigs,
});
