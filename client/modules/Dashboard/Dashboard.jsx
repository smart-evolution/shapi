import React from 'react';
import PropTypes from 'prop-types';
import TemperaturePanel from './TemperaturePanel';
import PresencePanel from './PresencePanel';

const Dashboard = (props) => {
  const { temperature, motions, error } = props;

  return (
    <div className="dashboard">
      { error && (
        <div className="dashboard__error">
          {error}
        </div>
      )}

      <TemperaturePanel
        temperature={temperature}
      />
      <PresencePanel
        motions={motions}
      />
    </div>
  );
}

Dashboard.propTypes = {
  timestamp: PropTypes.number,
  temperature: PropTypes.array,
  motions: PropTypes.object,
  error: PropTypes.any,
};

Dashboard.defaultProps = {
  temperature: [],
  presence: '-',
};

export default Dashboard;
