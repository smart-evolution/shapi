import { connect } from 'react-redux';
import PresencePanel from './PresencePanel';

const mapStateToProps = state => ({
  presence: state.presence,
});

const mapDispatchToProps = () => ({});

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(PresencePanel);
