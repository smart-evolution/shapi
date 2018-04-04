import { connect } from 'react-redux';
import TemperaturePanel from './TemperaturePanel';

const mapStateToProps = state => ({
  temperatures: state.temperatures,
});

const mapDispatchToProps = () => ({});

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(TemperaturePanel);
