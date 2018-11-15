import { connect } from 'react-redux';
import CurrentPanel from './CurrentPanel';
import * as queries from '../../../models/agents/queries';

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
