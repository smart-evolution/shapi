// @flow
import React from 'react';
import * as types from '../types';
import drawChart from './drawChart';

type Props = {
  temperatures: $ReadOnlyArray<types.Temperature>,
};

class TemperatureChart extends React.PureComponent<Props> {
  static defaultProps = {
    temperatures: [],
  };

  componentDidMount() {
    const { temperatures } = this.props;

    if (this.chart !== null) {
      drawChart(this.chart, temperatures);
    }
  }

  componentWillReceiveProps() {
    const { temperatures } = this.props;

    if (this.chart !== null) {
      drawChart(this.chart, temperatures);
    }
  }

  chart: HTMLDivElement | null;

  render() {
    return (
      <div
        className="temperature-chart"
        ref={ref => {
          this.chart = ref;
        }}
      />
    );
  }
}

export default TemperatureChart;
