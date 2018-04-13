import { connect } from 'react-redux';
import * as actions from '../actions';
import ControlPanel from './ControlPanel';

const mapStateToProps = state => ({
  isAlerts: state.isAlerts,
});

const mapDispatchToProps = (dispatch) => ({
  onToggle: () => {
    dispatch(actions.toggleAlerts());
  },
});

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(ControlPanel);
