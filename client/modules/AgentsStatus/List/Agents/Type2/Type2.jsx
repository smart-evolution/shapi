import React from 'react';
import PropTypes from 'prop-types';

const Type2 = (props) => {
  const { id, name } = props;

  return (
    <li className="agent-type2">
      <a
        className="agent-type2__link"
        href={`/agent/${id}`}
      >
        {name}
      </a>
    </li>
  );
};

Type2.propTypes = {
  id: PropTypes.string,
  name: PropTypes.string,
};

export default Type2;
