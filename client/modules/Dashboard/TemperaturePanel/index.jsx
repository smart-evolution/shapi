import { connect } from 'react-redux';
import TemperaturePanel from './TemperaturePanel';
import { getTicks } from '../queries';

const mapStateToProps = (state) => {
  const temperatures = getTicks(state.times, state.temperatures);

  return {
    temperatures,
  };
};

const mapDispatchToProps = () => ({});

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(TemperaturePanel);
