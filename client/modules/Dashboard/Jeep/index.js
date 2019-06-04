// @flow
import { connect } from 'react-redux';
import * as proxyActions from 'client/models/proxy/actions';
import * as agentsTypes from 'client/models/agents/types';
import Jeep from './Jeep';

const mapStateToProps = () => ({});

const mapDispatchToProps = (dispatch: Function) => ({
  setup: (agent: agentsTypes.Agent) => {
    dispatch(proxyActions.createWebSocketClient(agent));
  },
  onPositionChange: (
    agent: agentsTypes.Agent,
    msg: { left: number, top: number }
  ) => {
    dispatch(proxyActions.sendMessage(agent, msg));
  },
});

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(Jeep);
