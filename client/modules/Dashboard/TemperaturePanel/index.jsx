import { connect } from 'react-redux';
import TemperaturePanel from './TemperaturePanel';
import * as queries from 'models/agents/queries';

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
