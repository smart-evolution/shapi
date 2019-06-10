// @flow
import * as constants from './constants';

export type AlertType = constants.ALERT_TYPE_INFO | constants.ALERT_TYPE_ERROR;

export type Alert = {
  message: string,
  type: AlertType,
  timestamp: Date,
  isOld: boolean,
};

export type State = {
  alerts: Array<AlertType>,
};
