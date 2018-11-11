import _ from 'lodash';
import { connect } from 'react-redux';
import List from './List';

const mapStateToProps = (state) => {
  const { agents } = state;

  const agentsList = _.map(agents, (agent) => {
    const { id, name, data } = agent;
    const { temperature, presence, gas } = data;

    return {
      id,
      name,
      temperature: _.first(temperature),
      isMotion: _.reduce(presence, (acc, val) => acc || Number(val), 0),
      isGas: _.reduce(gas, (acc, val) => acc || Number(val), 0),
    };
  });

  return {
    agents: agentsList,
  };
};

const mapDispatchToProps = () => ({});

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(AgentsList);
