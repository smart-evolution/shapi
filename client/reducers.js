import { combineReducers } from 'redux';
import agents from 'client/models/agents/reducers';
import agentConfigs from 'client/models/agentConfigs/reducers';
import alerts from 'client/models/alerts/reducers';
import proxy from 'client/models/proxy/reducers';
import application from 'client/modules/Application/reducers'

export default combineReducers({
  agents,
  agentConfigs,
  alerts,
  proxy,
  application,
});
