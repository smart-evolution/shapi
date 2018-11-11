import React from 'react';
import PropTypes from 'prop-types';

const Type2 = (props) => {
  const { agent } = props;

  return (
    <li className="agent-type2">
      <a
        className="agent-type2__link"
        href={`/agent/${agent.id}`}
      >
        {agent.name}
      </a>
    </li>
  );
};

Type2.propTypes = {
  agent: PropTypes.shape({
    id: PropTypes.string,
    name: PropTypes.string,
    data: PropTypes.object,
    type: PropTypes.string,
  }),
};

export default Type2;
