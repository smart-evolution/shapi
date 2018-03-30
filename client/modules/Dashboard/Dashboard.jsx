import React from 'react';
import PropTypes from 'prop-types';
import TemperaturePanel from './TemperaturePanel';
import PresencePanel from './PresencePanel';

const Dashboard = (props) => {
  const { timestamp, temperature, presence, error } = props;

  return (
    <div className="dashboard">
      { error && (
        <div className="dashboard__error">
          {error}
        </div>
      )}

      <TemperaturePanel
        temperature={temperature}
        timestamp={timestamp}
      />
      <PresencePanel
        presence={presence}
        timestamp={timestamp}
      />
    </div>
  );
}

Dashboard.propTypes = {
  timestamp: PropTypes.number,
  temperature: PropTypes.array,
  presence: PropTypes.string,
  motions: PropTypes.object,
  error: PropTypes.any,
};

Dashboard.defaultProps = {
  temperature: [],
  presence: '-',
};

export default Dashboard;
