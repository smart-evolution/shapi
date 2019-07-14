// @flow
import _ from 'lodash';
import React from 'react';
import * as agentsTypes from 'client/models/agents/types';
import * as agentsQueries from 'client/models/agents/queries';
import * as agentConfigTypes from 'client/models/agentConfigs/types';
import * as agentConfigQueries from 'client/models/agentConfigs/queries';
import Jeep from './Agents/Jeep';
import Type1 from './Agents/Type1';
import Type2 from './Agents/Type2';
import Unknown from './Agents/Unknown';

type Props = {
  isLoading: boolean,
  agents: $ReadOnlyArray<agentsTypes.Agent>,
  agentConfigs: $ReadOnlyArray<agentConfigTypes.AgentConfig>,
};

const List = (props: Props) => {
  const { agents, agentConfigs, isLoading } = props;

  const loader = <div className="c-loader" />;
  const content = !_.isEmpty(agents) ? (
    <ul className="agents-list__list">
      {_.map(agents, agent => {
        const noVersionedType = agentsQueries.getNoVersionedType(agent);
        const agentConfig = agentConfigQueries.getAgentConfigByAgentId(
          agentConfigs,
          agent.id
        );

        switch (noVersionedType) {
          case 'type1':
            return (
              <Type1 key={agent.id} agent={agent} agentConfig={agentConfig} />
            );

          case 'type2':
            return (
              <Type2 key={agent.id} agent={agent} agentConfig={agentConfig} />
            );

          case 'jeep':
            return (
              <Jeep key={agent.id} agent={agent} agentConfig={agentConfig} />
            );

          default:
            return (
              <Unknown key={agent.id} agent={agent} agentConfig={agentConfig} />
            );
        }
      })}
    </ul>
  ) : (
    <p>No agents available</p>
  );

  return (
    <div className="agents-list">
      <div className="agents-list__title">Available agents</div>
      {isLoading ? loader : content}
    </div>
  );
};

export default List;
