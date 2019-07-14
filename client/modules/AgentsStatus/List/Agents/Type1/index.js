// @flow
import { connect } from 'react-redux';
import Type1 from './Type1';

const mapStateToProps = (state, ownProps) => {
  const { agent, agentConfig } = ownProps;

  return {
    agent,
    agentConfig,
  };
};

const mapDispatchToProps = () => ({});

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(Type1);
