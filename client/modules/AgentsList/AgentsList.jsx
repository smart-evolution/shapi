import _ from 'lodash';
import React from 'react';
import PropTypes from 'prop-types';

const AgentsList = (props) => {
  const { agents } = props;

  return (
    <div className="agents-list">
      {_.map(agents, (agent) => (
        <div>
          {agent.name} t={agent.temperature}
        </div>
      ))}
    </div>
  );
};

AgentsList.propTypes = {
  agents: PropTypes.array,
};

export default AgentsList;
