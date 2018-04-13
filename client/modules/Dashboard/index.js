import { connect } from 'react-redux';
import Dashboard from './Dashboard';

const mapStateToProps = state => ({
  isAlerts: state.isAlerts,
  temperatures: state.temperatures,
  motions: state.motions,
  error: state.error,
});

const mapDispatchToProps = () => ({});

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(Dashboard);
