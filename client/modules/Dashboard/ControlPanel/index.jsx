import { connect } from 'react-redux';
import * as actions from '../actions';
import ControlPanel from './ControlPanel';

const mapStateToProps = state => ({
  isAlerts: state.isAlerts,
});

const mapDispatchToProps = dispatch => ({
  onToggle: () => {
    dispatch(actions.toggleAlerts());
  },
  sendAlert: () => {
    dispatch(actions.sendAlert());
  },
});

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(ControlPanel);
