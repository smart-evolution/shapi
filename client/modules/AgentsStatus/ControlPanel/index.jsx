// @flow
import { connect } from 'react-redux';
import * as agentsActions from 'client/models/agents/actions';
import ControlPanel from './ControlPanel';

const mapStateToProps = state => ({
  isAlerts: state.isAlerts,
});

const mapDispatchToProps = dispatch => ({
  sniffAgents: () => {
    dispatch(agentsActions.sniffAgents());
  },
  onToggle: () => {
    dispatch(agentsActions.toggleAlerts());
  },
  sendAlert: () => {
    dispatch(agentsActions.sendAlert());
  },
});

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(ControlPanel);
