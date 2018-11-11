import React from 'react';
import PropTypes from 'prop-types';
import Icon from '../../../../../components/Icon';
import * as queries from './queries';

const Type1 = (props) => {
  const { agent } = props;
  const { id, name } = agent;

  const motionColor = queries.isMotion(agent) ?
    'agent-type1__icon--alert' :
    null;

  const gasColor = queries.isGas(agent) ?
    'agent-type1__icon--alert' :
    null;

  return (
    <li className="agent-type1">
      <a
        className="agent-type1__link"
        href={`/agent/${id}`}
      >
        {name}
      </a> - <span>
        {queries.getTemperature(agent)} <Icon type="thermometer" />
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
};

Type1.propTypes = {
  agent: PropTypes.shape({
    id: PropTypes.string,
    name: PropTypes.string,
    data: PropTypes.object,
    type: PropTypes.string,
  }),
};

export default Type1;
