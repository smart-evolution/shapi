import { connect } from 'react-redux';
import * as agentActions from 'client/models/agents/actions';
import Type2 from './Type2';

const mapStateToProps = (state, ownProps) => {
  const { agent } = ownProps;
  const { id, name, isOnline, type } = agent;

  return {
    id,
    name,
    isOnline,
    type,
  };
};

const mapDispatchToProps = dispatch => ({
  toggle: agentID => () => {
    dispatch(agentActions.toggleType2(agentID));
  },
});

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(Type2);
