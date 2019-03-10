import _ from 'lodash';
import React from 'react';
import PropTypes from 'prop-types';
import * as d3 from 'd3';

const CHART_PADDING = 20;
const DEFAULT_MAX_SOUND = 30;
const DEFAULT_MIN_SOUND = 10;
const SOUND_MARGIN = 2;

class SoundChart extends React.PureComponent {
  componentWillReceiveProps(props) {
    const { sounds } = props;

    const d3Chart = d3.select(this.chart);

    d3Chart.select('svg').remove();

    const width = d3Chart.node().clientWidth - CHART_PADDING;
    const containerHeight = width * (1.2 / 3);
    const chartHeight = containerHeight - 50;

    const maxObj = _.maxBy(sounds, 'value');
    const maxSound = _.isUndefined(maxObj)
      ? DEFAULT_MAX_SOUND
      : Number(maxObj.value) + SOUND_MARGIN;

    const minObj = _.minBy(sounds, 'value');
    const minSound = _.isUndefined(minObj)
      ? DEFAULT_MIN_SOUND
      : Number(minObj.value) - SOUND_MARGIN;

    const earliestData = sounds[0];
    const latestData = sounds[sounds.length - 1];

    const xScale = d3
      .scaleTime()
      .domain([earliestData.time, latestData.time])
      .range([0, width]);

    const yScale = d3
      .scaleLinear()
      .domain([minSound, maxSound])
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
      .attr('transform', `translate(${CHART_PADDING}, 0)`);

    svg
      .append('g')
      .attr('class', 'sound-chart__x-axis')
      .attr('transform', `translate(0, ${chartHeight})`)
      .call(d3.axisBottom(xScale).tickFormat(d3.timeFormat('%Y-%m-%d [%I:%M]')))
      .selectAll('text')
      .attr('transform', 'rotate(-20) translate(0, 15)');

    svg
      .append('g')
      .attr('class', 'sound-chart__y-axis')
      .call(
        d3
          .axisLeft(yScale)
          .ticks(6)
          .tickFormat(d3.format('d'))
      );

    svg
      .append('path')
      .datum(sounds)
      .attr('class', 'sound-chart__area')
      .attr('d', area);

    svg
      .append('path')
      .datum(sounds)
      .attr('class', 'sound-chart__line')
      .attr('d', line);

    svg
      .selectAll('.sound-chart__dot')
      .data(sounds)
      .enter()
      .append('circle')
      .attr('class', 'sound-chart__dot')
      .attr('cx', (d, i) => xScale(i))
      .attr('cy', d => yScale(d.value))
      .attr('r', 5);
  }

  render() {
    return (
      <div
        className="sound-chart"
        ref={ref => {
          this.chart = ref;
        }}
      />
    );
  }
}

SoundChart.propTypes = {
  sounds: PropTypes.arrayOf(PropTypes.string),
};

SoundChart.defaultProps = {
  sounds: [],
};

export default SoundChart;
