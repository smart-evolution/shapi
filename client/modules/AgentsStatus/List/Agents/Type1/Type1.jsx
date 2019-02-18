import React from 'react';
import PropTypes from 'prop-types';
import Icon from 'components/Icon';


const Type1 = (props) => {
  const { id, name, temperature, isGas, isMotion } = props;

  const motionColor = isMotion ?
    'agent-type1__icon--alert' :
    null;

  const gasColor = isGas ?
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

Type1.propTypes = {
  id: PropTypes.string,
  name: PropTypes.string,
  temperature: PropTypes.string,
  isMotion: PropTypes.number,
  isGas: PropTypes.number,
};

export default Type1;
