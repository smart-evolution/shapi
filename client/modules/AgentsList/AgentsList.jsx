import _ from 'lodash';
import React from 'react';
import PropTypes from 'prop-types';

const AgentsList = (props) => {
  const { agents, error } = props;

  return (
    <div className="agents-list">
      { error && (
        <div className="agents-list__error">
          {error}
        </div>
      )}
      <ul className="c-list">
        {_.map(agents, (agent) => {
          const { id, name, temperature, isMotion } = agent;
          return (
            <li className="c-list__item">
              <a href={`/agent/${id}`}>{name}</a> t[{temperature}] m[{isMotion}]
            </li>
          );
        })}
      </ul>
    </div>
  );
};

AgentsList.propTypes = {
  agents: PropTypes.arrayOf(PropTypes.shape({
    id: PropTypes.string,
    name: PropTypes.string,
    temperature: PropTypes.string,
    isMotion: PropTypes.number,
  })),
  error: PropTypes.error,
};

export default AgentsList;
