// @flow
import { connect } from 'react-redux';
import * as actions from 'client/models/agentConfigs/actions';
import * as agentConfigSelectors from 'client/models/agentConfigs/selectors';
import * as agentSelectors from 'client/models/agents/selectors';
import AgentEdit from './AgentEdit';

const mapStateToProps = (state, ownProps) => {
  const {
    match: { params },
  } = ownProps;
  const agentID = params.agent;
  const agentConfig =
    agentConfigSelectors.getAgentConfigById(state, agentID) || {};
  const agent = agentSelectors.getAgentById(state, agentID);

  return {
    timestamp: new Date(),
    agent,
    agentConfig,
  };
};

const mapDispatchToProps = dispatch => ({
  updateProperty: (agentID, key, value) => {
    dispatch(actions.updateProperty(agentID, key, value));
  },
  commitConfig: (agentID, config) => {
    dispatch(actions.commitAgentConfig(agentID, config));
  },
});

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(AgentEdit);
