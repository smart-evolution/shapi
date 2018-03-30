import { connect } from 'react-redux';
import TemperaturePanel from './TemperaturePanel';

const mapStateToProps = state => ({
  temperature: state.temperature,
  timestamp: state.timestamp,
});

const mapDispatchToProps = () => ({});

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(TemperaturePanel);
