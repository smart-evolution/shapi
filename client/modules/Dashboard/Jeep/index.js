// @flow
import { connect } from 'react-redux';
import * as proxyActions from 'client/models/proxy/actions';
import * as proxySelectors from 'client/models/proxy/selectors';
import * as agentsTypes from 'client/models/agents/types';
import Jeep from './Jeep';

const mapStateToProps = state => ({
  isDevConnected: proxySelectors.getIsDevConnected(state),
});

const mapDispatchToProps = (dispatch: Function, ownProps) => ({
  setup: (agent: agentsTypes.Agent) => {
    dispatch(proxyActions.createWebSocketClient(agent));
  },
  onToggle: () => {
    dispatch(
      proxyActions.sendMessage(ownProps.agent, {
        left: 25,
        top: 25,
      })
    );
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
