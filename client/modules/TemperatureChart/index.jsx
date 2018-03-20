import { connect } from 'react-redux';
import TemperatureChart from './TemperatureChart';

const mapStateToProps = state => ({
  temperature: state.temperature,
});

const mapDispatchToProps = () => ({});

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(TemperatureChart);
