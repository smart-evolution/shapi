import React from 'react';
import PropTypes from 'prop-types';

const Jeep = props => {
  const { id, name } = props;

  return (
    <li className="agent-jeep">
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
