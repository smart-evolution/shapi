import React from 'react';
import PropTypes from 'prop-types';
import TemperaturePanel from './TemperaturePanel';
import PresencePanel from './PresencePanel';

const Dashboard = (props) => {
  const { temperatures, motions, error } = props;

  return (
    <div className="dashboard">
      { error && (
        <div className="dashboard__error">
          {error}
        </div>
      )}

      <TemperaturePanel
        temperatures={temperatures}
      />
      <PresencePanel
        motions={motions}
      />
    </div>
  );
}

Dashboard.propTypes = {
  temperatures: PropTypes.array,
  motions: PropTypes.object,
  error: PropTypes.any,
};

Dashboard.defaultProps = {
  temperatures: [],
  motions: {},
};

export default Dashboard;
