import { connect } from 'react-redux';
import * as queries from 'client/models/agents/queries';
import TemperaturePanel from './TemperaturePanel';

const mapStateToProps = (state, ownProps) => {
  const { agent } = ownProps;
  const tmpArr = queries.getTemperatures(agent);
  const timeArr = queries.getTimes(agent);
  const temperatures = queries.getTicks(timeArr, tmpArr);

  return {
    temperatures,
  };
};

const mapDispatchToProps = () => ({});

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(TemperaturePanel);
