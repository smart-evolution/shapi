// @flow
import * as constants from './constants';

export type Status =
  | constants.STATUS_DISCONNECTED
  | constants.STATUS_PENDING
  | constants.STATUS_CONNECTED;
export type Flag = constants.FLAG_CONNECT | constants.FLAG_DISCONNECT | null;
export type Message = {
  left: number,
  top: number,
  flag: Flag,
};
