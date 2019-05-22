import _ from 'lodash';
import React from 'react';
import PropTypes from 'prop-types';
import { withRouter } from 'react-router';
import Type1 from './Type1';
import Type2 from './Type2';
import Jeep from './Jeep/';

const Dashboard = props => {
  const { error, agent, pathname } = props;

  let content;

  switch (agent.type) {
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

Dashboard.propTypes = {
  pathname: PropTypes.string,
  error: PropTypes.string,
  agent: PropTypes.shape(),
};

export default withRouter(Dashboard);
