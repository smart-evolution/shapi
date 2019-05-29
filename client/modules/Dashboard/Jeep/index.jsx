// @flow
import React from 'react';
import agentsTypes from 'models/agents/types';
import Joystick from 'components/Joystick';

type Props = {
  agent: agentsTypes.Agent,
};

const Jeep = (props: Props) => {
  const { agent } = props;

  return (
    <div className="dashboard__cell dashboard__cell--full">
      <Joystick
        onPositionChange={(left, top) => {}}
      />
    </div>
  );
};

export default Jeep;
