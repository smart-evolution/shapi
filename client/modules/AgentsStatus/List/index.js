import { connect } from 'react-redux';
import List from './List';

const mapStateToProps = state => {
  const {
    agents: { agents, isLoading },
  } = state;

  return {
    agents,
    isLoading,
  };
};

const mapDispatchToProps = () => ({});

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(List);
