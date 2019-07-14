// @flow
import _ from 'lodash';
import React from 'react';
import Icon from 'client/components/Icon';
import * as agentTypes from 'client/models/agents/types';
import * as agentQueries from 'client/models/agents/queries';
import * as agentConfigsTypes from 'client/models/agentConfigs/types';

type Props = {
  agent: agentTypes.Agent,
  agentConfig: agentConfigsTypes.AgentConfig,
};

const Type1 = (props: Props) => {
  const { agent, agentConfig } = props;

  const temperature = agentQueries.getTemperature(agent);
  const isMotion = agentQueries.isMotion(agent);
  const isGas = agentQueries.isGas(agent);

  const onlineClass = !agent.isOnline ? 'agent-type1--disabled' : '';
  const motionColor = isMotion ? 'agent-type1__icon--alert' : '';
  const gasColor = isGas ? 'agent-type1__icon--alert' : '';

  const humanName = _.isEmpty(agentConfig.name) ? '' : `${agentConfig.name} - `;

  return (
    <li className={`agent-type1 ${onlineClass}`}>
      <a className="agent-type1__link" href={`/agent/${agent.id}`}>
        {humanName}
        {agent.name} [{agent.type}]
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
