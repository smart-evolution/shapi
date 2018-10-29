import _ from 'lodash';
import React from 'react';
import PropTypes from 'prop-types';

const AgentsList = (props) => {
  const { agents } = props;

  return (
    <ul className="c-list">
      {_.map(agents, (agent) => (
        <li class="c-list__item">
          <a href={`/agent/${agent.id}`}>
            {agent.name}
          </a>
          t={agent.temperature}
        </li>
      ))}
    </ul>
  );
};

AgentsList.propTypes = {
  agents: PropTypes.array,
};

export default AgentsList;
