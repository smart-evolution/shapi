import { connect } from 'react-redux';
import Dashboard from './Dashboard';

const mapStateToProps = state => ({
  temperature: state.temperature,
  presence: state.presence,
  motions: state.motions,
  error: state.error,
});

const mapDispatchToProps = () => ({});

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(Dashboard);
