// @flow
import React from 'react';

type Props = {
  id: string,
  name: string,
  type: string,
  isOnline: boolean,
};

const Unknown = (props: Props) => {
  const { id, name, isOnline, type } = props;

  const onlineClass = !isOnline ? 'agent-type1--disabled' : '';

  return (
    <li className={`agent-unknown ${onlineClass}`}>
      Unknown agent [ID: {id} / Name: {name} / Type: {type}]
    </li>
  );
};

export default Unknown;
