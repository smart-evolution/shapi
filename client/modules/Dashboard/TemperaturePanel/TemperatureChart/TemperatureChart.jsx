import React from 'react';
import PropTypes from 'prop-types';
import * as d3 from 'd3';

const CHART_PADDING = 20;
const DEFAULT_MAX_TEMP = 30;
const DEFAULT_MIN_TEMP = 10;
const TEMP_MARGIN = 5;

class TemperatureChart extends React.PureComponent {
  componentWillReceiveProps(props) {
    const { temperatures } = props;

    const d3Chart = d3.select(this.chart);

    d3Chart.select('svg').remove();

    const height = 300;
    const width = d3Chart.node().clientWidth - CHART_PADDING;

    const maxObj = _.maxBy(temperatures, 'value');
    const maxTemp = _.isUndefined(maxObj) ? DEFAULT_MAX_TEMP : Number(maxObj.value) + TEMP_MARGIN;

    const minObj = _.minBy(temperatures, 'value');
    const minTemp = _.isUndefined(minObj) ? DEFAULT_MIN_TEMP : Number(minObj.value) - TEMP_MARGIN;

    const earliestData = temperatures[0];
    const latestData = temperatures[temperatures.length -1]

    const xScale = d3.scaleTime()
      .domain([earliestData.time, latestData.time])
      .range([0, width]);

    const line = d3.line()
      .x(d => xScale(d.time))
      .y(d => yScale(Number(d.value)))
      .curve(d3.curveMonotoneX);

    const yScale = d3.scaleLinear()
      .domain([minTemp, maxTemp])
      .range([height, 0]);

    const area = d3.area()
      .x(d => xScale(d.time))
      .y0(height)
      .y1(d => yScale(Number(d.value)))
      .curve(d3.curveMonotoneX);

    const svg = d3Chart.append('svg')
      .attr('width', width)
      .attr('height', height)
      .append('g')
      .attr('transform', `translate(${CHART_PADDING}, -70)`);

    svg.append('g')
      .attr('class', 'temperature-chart__x-axis')
      .attr('transform', `translate(0, ${height})`)
      .call(d3.axisBottom(xScale).tickFormat(d3.timeFormat("%Y-%m-%d [%I:%M]")))
      .selectAll("text")
      .attr("transform", "rotate(-30) translate(0, 40)")
      .attr("dy", ".15em");

    svg.append('g')
      .attr('class', 'temperature-chart__y-axis')
      .call(d3.axisLeft(yScale));

    svg.append("path")
      .datum(temperatures)
      .attr("class", "temperature-chart__area")
      .attr("d", area);

    svg.append('path')
      .datum(temperatures)
      .attr('class', 'temperature-chart__line')
      .attr('d', line);

    svg.selectAll('.temperature-chart__dot')
      .data(temperatures)
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
  temperatures: PropTypes.array,
};

TemperatureChart.defaultProps = {
  temperatures: [],
};

export default TemperatureChart;
