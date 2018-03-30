import { connect } from 'react-redux';
import PresencePanel from './PresencePanel';

const mapStateToProps = state => ({
  presence: state.presence,
  motions: state.motions,
});

const mapDispatchToProps = () => ({});

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(PresencePanel);
