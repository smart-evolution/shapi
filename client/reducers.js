import { combineReducers } from 'redux';
import agents from 'models/agents/reducers';
import agentConfigs from 'models/agentConfigs/reducers';
import alerts from 'models/alerts/reducers';
import proxy from 'models/proxy/reducers';

export default combineReducers({
  agents,
  agentConfigs,
  alerts,
  proxy,
});
