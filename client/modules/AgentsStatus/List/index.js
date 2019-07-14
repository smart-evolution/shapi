import { connect } from 'react-redux';
import List from './List';

const mapStateToProps = state => {
  const {
    agents: { agents, isLoading },
    agentConfigs,
  } = state;

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
