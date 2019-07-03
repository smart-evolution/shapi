// @flow
import React from 'react';
import Icon from 'client/components/Icon';

type Props = {
  id: string,
  name: string,
  type: string,
  temperature: string,
  isMotion: number,
  isGas: number,
  isOnline: boolean,
};

const Type1 = (props: Props) => {
  const { id, name, temperature, type, isGas, isMotion, isOnline } = props;

  const onlineClass = !isOnline ? 'agent-type1--disabled' : '';
  const motionColor = isMotion ? 'agent-type1__icon--alert' : '';
  const gasColor = isGas ? 'agent-type1__icon--alert' : '';

  return (
    <li className={`agent-type1 ${onlineClass}`}>
      <a className="agent-type1__link" href={`/agent/${id}`}>
        {name} [{type}]
      </a>{' '}
      -{' '}
      <span>
        {temperature} <Icon type="thermometer" />
        <Icon className={motionColor} type="motion" />
        <Icon className={gasColor} type="fire" />
      </span>
    </li>
  );
};

export default Type1;
