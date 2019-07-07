// @flow
import { connect } from 'react-redux';
import * as proxyActions from 'client/models/proxy/actions';
import * as proxyConstants from 'client/models/proxy/constants';
import * as proxyTypes from 'client/models/proxy/types';
import * as proxySelectors from 'client/models/proxy/selectors';
import * as agentsTypes from 'client/models/agents/types';
import Jeep from './Jeep';

const mapStateToProps = state => ({
  status: proxySelectors.getStatus(state),
});

const mapDispatchToProps = (dispatch: Function, ownProps) => ({
  onToggle: (agent: agentsTypes.Agent, isConnected: boolean) => {
    if (isConnected) {
      dispatch(
        proxyActions.sendMessage(ownProps.agent, {
          left: 25,
          top: 25,
          flag: proxyConstants.FLAG_DISCONNECT,
        })
      );
    } else {
      dispatch(proxyActions.createWebSocketClient(agent));
    }
  },
  onPositionChange: (agent: agentsTypes.Agent, msg: proxyTypes.Message) => {
    dispatch(proxyActions.sendMessage(agent, msg));
  },
});

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(Jeep);
