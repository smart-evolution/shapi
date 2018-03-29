import React from 'react';
import PropTypes from 'prop-types';
import TemperaturePanel from './TemperaturePanel';
import PresencePanel from './PresencePanel';

const Dashboard = (props) => {
  const { temperature, presence } = props;

  return (
    <div className="dashboard">
      <TemperaturePanel temperature={temperature} />
      <PresencePanel presence={presence} />
    </div>
  );
}

Dashboard.propTypes = {
  temperature: PropTypes.array,
  presence: PropTypes.string,
};

Dashboard.defaultProps = {
  temperature: [],
  presence: "0",
};

export default Dashboard;
