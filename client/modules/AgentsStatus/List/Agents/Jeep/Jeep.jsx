import React from 'react';
import PropTypes from 'prop-types';

const Jeep = props => {
  const { id, name, isOnline } = props;
  const onlineClass = !isOnline ? 'agent-jeep--disabled' : '';
  return (
    <li className={`agent-jeep ${onlineClass}`}>
      <a className="agent-jeep__link" href={`/agent/${id}`}>
        {name}
      </a>
    </li>
  );
};

Jeep.propTypes = {
  id: PropTypes.string,
  name: PropTypes.string,
};

export default Jeep;
