// @flow
import { connect } from 'react-redux';
import * as agentSelectors from 'client/models/agents/selectors';
import Dashboard from './Dashboard';

const mapStateToProps = (state, ownProps) => {
  const {
    match: {
      params: { agent },
    },
    location: { pathname },
  } = ownProps;
  const agentId = agent;

  return {
    error: state.error,
    agent: agentSelectors.getAgentById(state, agentId),
    pathname,
  };
};

const mapDispatchToProps = () => ({});

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(Dashboard);
