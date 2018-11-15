import _ from 'lodash';
import React from 'react';
import PropTypes from 'prop-types';
import { withRouter } from "react-router";
import TemperaturePanel from './TemperaturePanel';
import SoundPanel from './SoundPanel';
import CurrentPanel from './CurrentPanel';

const Dashboard = (props) => {
  const {
    error,
    agent,
  } = props;

  return (
    <div className="dashboard">
      { error && (
        <div className="dashboard__error">
          {error}
        </div>
      )}
      {!_.isEmpty(agent) &&
        (<div className="dashboard__cell dashboard__cell--full">
          <CurrentPanel
            agent={agent}
          />
        </div>)
      }
      {!_.isEmpty(agent) &&
        (<div className="dashboard__cell dashboard__cell--full">
          <TemperaturePanel
            agent={agent}
          />
        </div>)
      }
      {!_.isEmpty(agent) &&
        (<div className="dashboard__cell dashboard__cell--full">
          <SoundPanel
            agent={agent}
          />
        </div>)
      }
    </div>
  );
};

Dashboard.propTypes = {
  error: PropTypes.error,
};

export default withRouter(Dashboard);
