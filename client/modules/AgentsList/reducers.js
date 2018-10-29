import _ from 'lodash';
import * as actionTypes from './actionTypes';

const defaultState = {
  agents: [],
};

export default function reducer(state = defaultState, action) {
  const { agents } = action;

  switch (action.type) {
    case actionTypes.DATA_FETCH:
      return Object.assign({}, state);

    case actionTypes.DATA_FETCH_SUCCESS:
      const list = _.map(agents, (agent) => ({
        id: agent.id,
        name: agent.name,
        temperature: _.first(agent.data.temperature),
      }))

      return Object.assign({}, state, { list });

    case actionTypes.DATA_FETCH_ERROR:
      return Object.assign({}, state, { error: action.error });

    default:
      return state;
  }
}
