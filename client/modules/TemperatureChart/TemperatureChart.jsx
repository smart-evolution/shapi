import React from 'react';
import PropTypes from 'prop-types';

const TemperatureChart = (props) => {
  const { temperature } = props;

  return (<div className="temperature-chart">
    Current temp: {temperature} *C
  </div>);
};

TemperatureChart.propTypes = {
  temperature: PropTypes.number,
};

TemperatureChart.defaultProps = {
  temperature: 0,
};

export default TemperatureChart;
