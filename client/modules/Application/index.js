// @flow
import { connect } from 'react-redux';
import Application from './Application';
import * as actions from './actions';
import * as selectors from './selectors';

const mapStateToProps = state => {
  return {
    isLoaded: selectors.getIsLoaded(state),
  };
};

const mapDispatchToProps = dispatch => ({
  mount: () => {
    dispatch(actions.mount());
  },
});

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(Application);
