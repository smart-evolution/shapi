import { connect } from 'react-redux';
import * as queries from 'models/agents/queries';
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
    agent: queries.getAgentById(state, agentId),
    pathname,
  };
};

const mapDispatchToProps = () => ({});

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(Dashboard);
