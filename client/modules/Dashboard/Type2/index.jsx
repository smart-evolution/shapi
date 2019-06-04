// @flow
import React from 'react';
import * as agentsTypes from 'client/models/agents/types';

type Props = {
  agent: agentsTypes.Agent,
};

const Type2 = (props: Props) => {
  const { agent } = props;

  return (
    <div className="dashboard__cell dashboard__cell--full">{agent.id}</div>
  );
};

export default Type2;
