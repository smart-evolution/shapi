import { connect } from 'react-redux';
import * as queries from 'client/models/agents/queries';
import Type1 from './Type1';

const mapStateToProps = (state, ownProps) => {
  const { agent } = ownProps;
  const { id, name, isOnline, type } = agent;

  const temperature = queries.getTemperature(agent);
  const isMotion = queries.isMotion(agent);
  const isGas = queries.isGas(agent);

  return {
    id,
    name,
    temperature,
    type,
    isMotion,
    isGas,
    isOnline,
  };
};

const mapDispatchToProps = () => ({});

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(Type1);
