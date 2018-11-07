import _ from 'lodash';
import React from 'react';
import PropTypes from 'prop-types';
import Icon from '../../components/Icon';

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
          const { id, name, temperature, isMotion, isGas } = agent;

          const motionColor = isMotion ?
            'c-icon--red' :
            '';

          const gasColor = isGas ?
            'c-icon--red' :
            '';

          return (
            <li className="c-list__item">
              <a
                className="agents-list__link"
                href={`/agent/${id}`}
              >
                {name}
              </a> - <span>
                {temperature} <Icon type="thermometer" />
                <Icon
                  className={motionColor}
                  type="motion"
                />
                <Icon
                  className={gasColor}
                  type="fire"
                />
              </span>
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
    isGas: PropTypes.number,
  })),
  error: PropTypes.error,
};

export default AgentsList;
