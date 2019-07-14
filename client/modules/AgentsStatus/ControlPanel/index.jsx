// @flow
import { connect } from 'react-redux';
import * as agentsActions from 'client/models/agents/actions';
import * as actions from '../actions';
import ControlPanel from './ControlPanel';

const mapStateToProps = state => ({
  isAlerts: state.isAlerts,
});

const mapDispatchToProps = dispatch => ({
  sniffAgents: () => {
    dispatch(agentsActions.sniffAgents());
  },
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
