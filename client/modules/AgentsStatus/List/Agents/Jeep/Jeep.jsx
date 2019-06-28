// @flow
import React from 'react';
import PropTypes from 'prop-types';

type Props = {
  id: PropTypes.string,
  name: PropTypes.string,
  isOnline: boolean,
};

const Jeep = (props: Props) => {
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

export default Jeep;
