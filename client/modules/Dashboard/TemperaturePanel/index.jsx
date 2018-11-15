import { connect } from 'react-redux';
import TemperaturePanel from './TemperaturePanel';
import { getTicks } from '../queries';
import * as selectors from '../../../models/agents/selectors';
import * as queries from '../../../models/agents/queries';

const mapStateToProps = (state, ownProps) => {
  const { agent } = ownProps;
  const temperatures = selectors.getTemperatures(agent);
  const times = selectors.getTimes(agent);
  const tempArray = queries.getTicks(times, temperatures);

  return {
    temperatures: tempArray,
  };
};

const mapDispatchToProps = () => ({});

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(TemperaturePanel);
