import React from 'react';
import PropTypes from 'prop-types';
import ControlPanel from './ControlPanel'
import TemperaturePanel from './TemperaturePanel';
import PresencePanel from './PresencePanel';
import GasPanel from './GasPanel';

const Dashboard = (props) => {
  const { isAlerts, temperatures, motions, gas, error } = props;

  return (
    <div className="dashboard">
      { error && (
        <div className="dashboard__error">
          {error}
        </div>
      )}

      <div className="dashboard__cell dashboard__cell--full">
        <ControlPanel
          isAlerts={isAlerts}
        />
      </div>
      <div className="dashboard__cell">
        <TemperaturePanel
          temperatures={temperatures}
        />
      </div>
      <div className="dashboard__cell">
        <PresencePanel
          motions={motions}
        />
      </div>
      <div className="dashboard__cell dashboard__cell--full">
        <GasPanel
          gas={gas}
        />
      </div>
    </div>
  );
}

Dashboard.propTypes = {
  isAlerts: PropTypes.bool,
  temperatures: PropTypes.array,
  motions: PropTypes.object,
  gas: PropTypes.bool,
  error: PropTypes.any,
};

Dashboard.defaultProps = {
  temperatures: [],
  motions: {},
  gas: false,
};

export default Dashboard;
