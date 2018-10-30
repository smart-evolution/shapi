import _ from 'lodash';
import { connect } from 'react-redux';
import AgentsList from './AgentsList';

const mapStateToProps = (state) => {
  const { agents, error } = state;

  const agentsList = _.map(agents, (agent) => {
    const { id, name, data } = agent;
    const { temperature, presence } = data;

    return {
      id,
      name,
      temperature: _.first(temperature),
      isMotion: _.reduce(presence, (acc, val) => acc || Number(val), 0),
    };
  });

  return {
    agents: agentsList,
    error,
  };
};

const mapDispatchToProps = () => ({});

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(AgentsList);
