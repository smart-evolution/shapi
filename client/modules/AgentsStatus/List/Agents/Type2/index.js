import { connect } from 'react-redux';
import * as actions from '../../../actions';
import Type2 from './Type2';

const mapStateToProps = (state, ownProps) => {
  const { agent } = ownProps;
  const { id, name, isOnline } = agent;

  return {
    id,
    name,
    isOnline,
  };
};

const mapDispatchToProps = dispatch => ({
  toggle: agentID => () => {
    dispatch(actions.toggleType2(agentID));
  },
});

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(Type2);
