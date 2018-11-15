import React from 'react';
import PropTypes from 'prop-types';

const Type2 = (props) => {
  const { id, name, toggle } = props;

  return (
    <li className="agent-type2">
      <a
        className="agent-type2__link"
        href={`/agent/${id}`}
      >
        {name}
      </a>
      <button onClick={toggle(id)}>
        Toggle
      </button>
    </li>
  );
};

Type2.propTypes = {
  id: PropTypes.string,
  name: PropTypes.string,
  toggle: PropTypes.func,
};

export default Type2;
