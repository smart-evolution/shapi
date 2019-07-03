// @flow
import * as d3 from 'd3';
import _ from 'lodash';
import * as types from '../types';
import * as constants from './constants';

const drawChart = (
  chart: HTMLDivElement,
  temperatures: $ReadOnlyArray<types.Temperature>
) => {
  const d3Chart = d3.select(chart);

  d3Chart.select('svg').remove();

  const width = d3Chart.node().clientWidth - constants.CHART_PADDING;
  const containerHeight = width * (1.2 / 3);
  const chartHeight = containerHeight - 50;

  const maxObj = _.maxBy(temperatures, 'value');
  const maxTemp = _.isUndefined(maxObj)
    ? constants.DEFAULT_MAX_TEMP
    : Number(maxObj.value) + constants.TEMP_MARGIN;

  const minObj = _.minBy(temperatures, 'value');
  const minTemp = _.isUndefined(minObj)
    ? constants.DEFAULT_MIN_TEMP
    : Number(minObj.value) - constants.TEMP_MARGIN;

  const earliestData = temperatures[0];
  const latestData = temperatures[temperatures.length - 1];

  const xScale = d3
    .scaleTime()
    .domain([earliestData.time, latestData.time])
    .range([0, width]);

  const yScale = d3
    .scaleLinear()
    .domain([minTemp, maxTemp])
    .range([chartHeight, 0]);

  const line = d3
    .line()
    .x(d => xScale(d.time))
    .y(d => yScale(Number(d.value)))
    .curve(d3.curveMonotoneX);

  const area = d3
    .area()
    .x(d => xScale(d.time))
    .y0(chartHeight)
    .y1(d => yScale(Number(d.value)))
    .curve(d3.curveMonotoneX);

  const svg = d3Chart
    .append('svg')
    .attr('width', width)
    .attr('height', containerHeight)
    .append('g')
    .attr('transform', `translate(${constants.CHART_PADDING}, 0)`);

  svg
    .append('g')
    .attr('class', 'temperature-chart__x-axis')
    .attr('transform', `translate(0, ${chartHeight})`)
    .call(d3.axisBottom(xScale).tickFormat(d3.timeFormat('%Y-%m-%d [%I:%M]')))
    .selectAll('text')
    .attr('transform', 'rotate(-20) translate(0, 15)');

  svg
    .append('g')
    .attr('class', 'temperature-chart__y-axis')
    .call(
      d3
        .axisLeft(yScale)
        .ticks(6)
        .tickFormat(d3.format('d'))
    );

  svg
    .append('path')
    .datum(temperatures)
    .attr('class', 'temperature-chart__area')
    .attr('d', area);

  svg
    .append('path')
    .datum(temperatures)
    .attr('class', 'temperature-chart__line')
    .attr('d', line);

  svg
    .selectAll('.temperature-chart__dot')
    .data(temperatures)
    .enter()
    .append('circle')
    .attr('class', 'temperature-chart__dot')
    .attr('cx', (d, i) => xScale(i))
    .attr('cy', d => yScale(d.value))
    .attr('r', 5);
};

export default drawChart;
