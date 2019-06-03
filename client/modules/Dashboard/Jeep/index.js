// @flow
import { connect } from 'react-redux';
import * as proxyActions from 'models/proxy/actions';
import agentsTypes from 'models/agents/types';
import Jeep from './Jeep';

const mapStateToProps = () => ({});

const mapDispatchToProps = (dispatch: Dispatch) => ({
  setup: (agent: agentsTypes.Agent) => {
    dispatch(proxyActions.createWebSocketClient(agent));
  },
  onPositionChange: (agent: agentsTypes.Agent, msg: string) => {
    dispatch(proxyActions.sendMessage(agent, msg));
  }
});

export default connect(
  mapStateToProps,
  mapDispatchToProps,
)(Jeep);
