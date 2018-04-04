import { connect } from 'react-redux';
import PresencePanel from './PresencePanel';

const mapStateToProps = state => ({
  motions: state.motions,
});

const mapDispatchToProps = () => ({});

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(PresencePanel);
