// @flow
import * as actionTypes from './actionTypes';

type State = {
  isLoaded: boolean,
};

const defaultState = {
  isLoaded: false,
};

export default function reducers(state: State = defaultState, action: Object) {
  switch (action.type) {
    case actionTypes.LOADED:
      return Object.assign({}, state, {
        isLoaded: true,
      });
    default:
      return state;
  }
}
