// @flow
import { connect } from 'react-redux';
import AgentsStatus from './AgentsStatus';

const mapStateToProps = state => {
  const { error } = state;

  return {
    error,
  };
};

const mapDispatchToProps = () => ({});

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(AgentsStatus);
