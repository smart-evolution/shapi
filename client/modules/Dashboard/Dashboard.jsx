import React from 'react';
import PropTypes from 'prop-types';
import ControlPanel from './ControlPanel'
import TemperaturePanel from './TemperaturePanel';
import SoundPanel from './SoundPanel';
import PresencePanel from './PresencePanel';
import CurrentPanel from './CurrentPanel';

const Dashboard = (props) => {
  const { error } = props;

  return (
    <div className="dashboard">
      { error && (
        <div className="dashboard__error">
          {error}
        </div>
      )}

      <div className="dashboard__cell dashboard__cell--full">
        <ControlPanel/>
      </div>
      <div className="dashboard__cell dashboard__cell--full">
        <CurrentPanel/>
      </div>
      <div className="dashboard__cell dashboard__cell--full">
        <TemperaturePanel/>
      </div>
      <div className="dashboard__cell dashboard__cell--full">
        <SoundPanel/>
      </div>
    </div>
  );
}

Dashboard.propTypes = {
  isAlerts: PropTypes.bool,
  error: PropTypes.string,
};

Dashboard.defaultProps = {
  isAlerts: [],
};

export default Dashboard;
