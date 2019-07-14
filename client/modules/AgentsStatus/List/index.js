// @flow
import { connect } from 'react-redux';
import * as agentSelectors from 'client/models/agents/selectors';
import * as agentConfigSelectors from 'client/models/agentConfigs/selectors';
import List from './List';

const mapStateToProps = state => {
  const agents = agentSelectors.getAgents(state);
  const isLoading = agentSelectors.isLoading(state);
  const agentConfigs = agentConfigSelectors.getAgentConfigs(state);

  return {
    agents,
    agentConfigs,
    isLoading,
  };
};

const mapDispatchToProps = () => ({});

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(List);
