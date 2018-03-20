import _ from 'lodash';
import React, { Component } from 'react';
import PropTypes from 'prop-types';

export default class TemperatureChart extends Component {
  constructor(props) {
    super(props);
  }


  render() {
    return (<div className="temperature-chart">
        Temperature chart
    </div>);
  }
}

TemperatureChart.propTypes = {
};

TemperatureChart.defaultProps = {
};
