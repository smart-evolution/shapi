// @flow
import _ from 'lodash';
import { connect } from 'react-redux';
import * as agentsActions from 'client/models/agents/actions';
import Type1 from './Type1';

const mapDispatchToProps = dispatch => ({
  onScroll: period => {
    dispatch(agentsActions.changePeriod(_.round(period)));
  },
});

export default connect(
  null,
  mapDispatchToProps
)(Type1);
