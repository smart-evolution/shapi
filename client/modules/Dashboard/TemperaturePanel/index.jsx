import { connect } from 'react-redux';
import TemperaturePanel from './TemperaturePanel';

const mapStateToProps = state => ({
  temperature: state.temperature,
});

const mapDispatchToProps = () => ({});

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(TemperaturePanel);
