// @flow
import _ from 'lodash';
import React from 'react';
import { withRouter } from 'react-router';
import * as agentTypes from 'client/models/agents/types';
import * as agentQueries from 'client/models/agents/queries';
import Type1 from './Type1';
import Type2 from './Type2';
import Jeep from './Jeep';

type Props = {
  pathname: string,
  error: string,
  agent: agentTypes.Agent,
};

const Dashboard = (props: Props) => {
  const { error, agent, pathname } = props;

  if (_.isEmpty(agent)) {
    return <div>no agent passed</div>;
  }

  let content;

  const noVersionedAgentType = agentQueries.getNoVersionedType(agent);
  switch (noVersionedAgentType) {
    case 'type1':
      content = <Type1 agent={agent} pathname={pathname} />;
      break;
    case 'type2':
      content = <Type2 agent={agent} />;
      break;
    case 'jeep':
      content = <Jeep agent={agent} />;
      break;
    default:
      content = <div>Unknown agent type</div>;
      break;
  }

  return (
    <div className="dashboard">
      {error && <div className="dashboard__error">{error}</div>}
      {content}
    </div>
  );
};

Dashboard.defaultProps = {
  error: '',
};

export default withRouter(Dashboard);
