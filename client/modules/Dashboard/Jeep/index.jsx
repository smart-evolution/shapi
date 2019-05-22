// @flow
import React from 'react';
import agentsTypes from 'models/agents/types';

type Props = {
  agent: agentsTypes.Agent,
};

const Jeep = (props: Props) => {
  const { agent } = props;

  return (
    <div className="dashboard__cell dashboard__cell--full">{agent.id}</div>
  );
};

export default Jeep;
