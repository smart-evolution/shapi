import { connect } from 'react-redux';
import * as selectors from 'models/alerts/selectors';
import AlertPanel from './AlertPanel';

const mapStateToProps = state => {
  const alerts = selectors.getLimitedAlerts(state);

  return {
    alerts,
  };
};

export default connect(
  mapStateToProps,
  null
)(AlertPanel);
