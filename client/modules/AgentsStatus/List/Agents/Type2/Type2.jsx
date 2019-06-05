// @flow
import React from 'react';

type Props = {
  id: string,
  name: string,
  toggle: string => void,
};

const Type2 = (props: Props) => {
  const { id, name, toggle } = props;

  return (
    <li className="agent-type2">
      <a className="agent-type2__link" href={`/agent/${id}`}>
        {name}
      </a>
      <button className="agent-type2__toggle" onClick={toggle(id)}>
        Toggle
      </button>
    </li>
  );
};

export default Type2;
