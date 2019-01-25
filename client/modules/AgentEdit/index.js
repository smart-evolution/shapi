import { connect } from 'react-redux';
import * as actions from 'models/agentConfigs/actions';
import AgentEdit from './AgentEdit';

const mapStateToProps = (state, ownProps) => ({});

const mapDispatchToProps = (dispatch) => ({
  fetchConfig: (agentID) => {
    dispatch(actions.fetchData(agentID))
  },
});

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(AgentEdit);
