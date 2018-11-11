import React from 'react';
import PropTypes from 'prop-types';
import Icon from '../../../../../components/Icon';

const Type2 = (props) => {
  const { id, name, temperature, isMotion, isGas } = props;

  const motionColor = isMotion ?
    'agent-type2__icon--alert' :
    null;

  const gasColor = isGas ?
    'agent-type2__icon--alert' :
    null;

  return (
    <li className="agent-type2">
      <a
        className="agent-type2__link"
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
};

Type2.propTypes = {
  id: PropTypes.string,
  name: PropTypes.string,
  temperature: PropTypes.string,
  isMotion: PropTypes.number,
  isGas: PropTypes.number,
};

export default Type2;
