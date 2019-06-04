import { connect } from 'react-redux';
import * as queries from 'client/models/agents/queries';
import CurrentPanel from './CurrentPanel';

const mapStateToProps = (state, ownProps) => {
  const { agent } = ownProps;
  const isMotion = queries.isMotion(agent);
  const isGas = queries.isGas(agent);

  return {
    isMotion,
    isGas,
  };
};

const mapDispatchToProps = () => ({});

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(CurrentPanel);
