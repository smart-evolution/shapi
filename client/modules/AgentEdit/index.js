import { connect } from 'react-redux';
import * as actions from 'models/agentConfigs/actions';
import * as selectors from 'models/agentConfigs/selectors';
import AgentEdit from './AgentEdit';

const mapStateToProps = (state, ownProps) => {
  const {
    match: {
      params: {
        agent,
      },
    },
  } = ownProps;
  const agentID = agent;
  const config = selectors.getAgentConfig(state, agentID);

  return {
    agentID,
    config,
  };
};

const mapDispatchToProps = dispatch => ({
  fetchConfig: (agentID) => {
    dispatch(actions.fetchData(agentID));
  },
  updateTemperature: (agentID, temperature) => {
    dispatch(actions.updateTemperature(agentID, temperature));
  },
  updateConfig: (agentID, config) => {
    dispatch(actions.updateDate(agentID, config));
  },
});

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(AgentEdit);
