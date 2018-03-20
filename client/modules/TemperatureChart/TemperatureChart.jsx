import _ from 'lodash';
import React, { Component } from 'react';
import PropTypes from 'prop-types';

export default class TemperatureChart extends Component {
  constructor(props) {
    super(props);
  }


  render() {
    const { temperature } = this.props;

    return (<div className="temperature-chart">
        Current temp: {temperature} *C
    </div>);
  }
}

TemperatureChart.propTypes = {
};

TemperatureChart.defaultProps = {
};
