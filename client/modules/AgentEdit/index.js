// @flow
import { connect } from 'react-redux';
import * as actions from 'client/models/agentConfigs/actions';
import * as selectors from 'client/models/agentConfigs/selectors';
import * as agentQueries from 'client/models/agents/queries';
import AgentEdit from './AgentEdit';

const mapStateToProps = (state, ownProps) => {
  const {
    match: {
      params,
    },
  } = ownProps;
  const agentID = params.agent;
  const agentConfig = selectors.getAgentConfigById(state, agentID) || {};
  const agent = agentQueries.getAgentById(state, agentID);

  return {
    timestamp: new Date(),
    agent,
    agentConfig
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
