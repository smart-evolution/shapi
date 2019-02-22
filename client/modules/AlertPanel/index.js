import { connect } from 'react-redux';
import AlertPanel from './AlertPanel';
import * as selectors from 'models/alerts/selectors';

const mapStateToProps = (state) => {
  const alerts = selectors.getAlerts(state);

  return {
    alerts,
  };
}

export default connect(
  mapStateToProps,
  null
)(AlertPanel);
