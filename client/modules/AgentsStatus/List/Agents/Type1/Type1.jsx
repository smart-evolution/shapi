// @flow
import React from 'react';
import Icon from 'client/components/Icon';

type Props = {
  id: string,
  name: string,
  temperature: string,
  isMotion: number,
  isGas: number,
};

const Type1 = (props: Props) => {
  const { id, name, temperature, isGas, isMotion } = props;

  const motionColor = isMotion ? 'agent-type1__icon--alert' : null;
  const gasColor = isGas ? 'agent-type1__icon--alert' : null;

  return (
    <li className="agent-type1">
      <a className="agent-type1__link" href={`/agent/${id}`}>
        {name}
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
