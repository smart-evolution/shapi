// @flow
import { connect } from 'react-redux';
import * as agentActions from 'client/models/agents/actions';
import * as agentSelectors from 'client/models/agents/selectors';
import ControlPanel from './ControlPanel';

const mapStateToProps = state => ({
  isAlerts: agentSelectors.isAlerts(state),
});

const mapDispatchToProps = dispatch => ({
  sniffAgents: () => {
    dispatch(agentActions.sniffAgents());
  },
  onToggle: () => {
    dispatch(agentActions.toggleAlerts());
  },
  sendAlert: () => {
    dispatch(agentActions.sendAlert());
  },
});

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(ControlPanel);
