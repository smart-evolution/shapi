import { connect } from 'react-redux';
import Dashboard from './Dashboard';

const mapStateToProps = state => ({
  temperature: state.temperature,
  presence: state.presence,
});

const mapDispatchToProps = () => ({});

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(Dashboard);
