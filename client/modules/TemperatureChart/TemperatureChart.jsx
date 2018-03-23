import React from 'react';
import PropTypes from 'prop-types';

const TemperatureChart = (props) => {
  const { temperature } = props;

  return (<div className="temperature-chart">
    Current: {temperature} &#8451;
  </div>);
};

TemperatureChart.propTypes = {
  temperature: PropTypes.number,
};

TemperatureChart.defaultProps = {
  temperature: 0,
};

export default TemperatureChart;
