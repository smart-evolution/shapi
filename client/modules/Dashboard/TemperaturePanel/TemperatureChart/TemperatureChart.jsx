import React from 'react';
import PropTypes from 'prop-types';
import * as d3 from 'd3';

const DEFAULT_MAX_TEMP = 30;
const DEFAULT_MIN_TEMP = 10;
const TEMP_MARGIN = 5;

class TemperatureChart extends React.PureComponent {
  componentWillReceiveProps(props) {
    const { temperature } = props;

    const d3Chart = d3.select(this.chart);

    d3Chart.select('svg').remove();

    const height = 300;
    const width = 800;

    const maxObj = _.maxBy(temperature, 'value');
    const maxTemp = _.isUndefined(maxObj) ? DEFAULT_MAX_TEMP : Number(maxObj.value) + TEMP_MARGIN;

    const minObj = _.minBy(temperature, 'value');
    const minTemp = _.isUndefined(minObj) ? DEFAULT_MIN_TEMP : Number(minObj.value) - TEMP_MARGIN;

    const earliestData = temperature[0];
    const latestData = temperature[temperature.length -1]

    const xScale = d3.scaleTime()
      .domain([earliestData.time, latestData.time])
      .range([0, width]);

    const yScale = d3.scaleLinear()
      .domain([minTemp, maxTemp])
      .range([height, 0]);

    const line = d3.line()
      .x(d => xScale(d.time))
      .y(d => yScale(Number(d.value)))
      .curve(d3.curveMonotoneX)

    const svg = d3Chart.append('svg')
      .attr('width', width)
      .attr('height', height)
      .append('g')
      .attr('transform', 'translate(50, -50)');

    svg.append('g')
      .attr('class', 'temperature-chart__x-axis')
      .attr('transform', `translate(0, ${height})`)
      .call(d3.axisBottom(xScale));

    svg.append('g')
      .attr('class', 'temperature-chart__y-axis')
      .call(d3.axisLeft(yScale));

    svg.append('path')
      .datum(temperature)
      .attr('class', 'temperature-chart__line')
      .attr('d', line);

    svg.selectAll('.temperature-chart__dot')
      .data(temperature)
      .enter().append('circle')
      .attr('class', 'temperature-chart__dot')
      .attr('cx', (d, i) =>  xScale(i))
      .attr('cy', (d) => yScale(d.value))
      .attr('r', 5);
  }

  render() {
    return (<div
      className="temperature-chart"
      ref={(ref) => { this.chart = ref; }}
    />);
  }
}

TemperatureChart.propTypes = {
  temperature: PropTypes.array,
};

TemperatureChart.defaultProps = {
  temperature: [],
};

export default TemperatureChart;
