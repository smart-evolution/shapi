import _ from 'lodash';
import React from 'react';
import PropTypes from 'prop-types';
import AgentType1 from './Agents/Type1';

const List = (props) => {
  const { agents } = props;

  return (
    <div className="agents-list">
      <div className="agents-list__title">
        Available agents
      </div>
      <ul className="agents-list__list">
        {_.map(agents, (agent) => {
          const { id, name, temperature, isMotion, isGas } = agent;

          return (<AgentType1
            id={id}
            name={name}
            temperature={temperature}
            isMotion={isMotion}
            isGas={isGas}
          />);
        })}
      </ul>
    </div>
  );
};

List.propTypes = {
  agents: PropTypes.arrayOf(PropTypes.shape({
    id: PropTypes.string,
    name: PropTypes.string,
    temperature: PropTypes.string,
    isMotion: PropTypes.number,
    isGas: PropTypes.number,
  })),
};

export default List;
