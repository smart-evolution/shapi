import _ from 'lodash';
import * as constants from './constants';

export const getAlerts = state => state.alerts.alerts;

export const getLimitedAlerts = state =>
  _.takeRight(getAlerts(state), constants.ALERT_LIMIT);
