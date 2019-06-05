// @flow
import _ from 'lodash';
import React from 'react';
import TemperatureChart from './TemperatureChart/TemperatureChart';
import * as types from './types';

type Props = {
  temperatures: $ReadOnlyArray<types.Temperature>,
};

const NODATA_SIGN = '-';

const TemperaturePanel = (props: Props) => {
  const { temperatures } = props;
  const nowTmp = _.head(temperatures);
  const value = _.isUndefined(nowTmp) ? NODATA_SIGN : nowTmp.value;

  return (
    <div className="temperature-panel">
      <div className="temperature-panel__title">Temperature</div>
      <div className="temperature-panel__current">{value} &#8451;</div>
      <div className="temperature-panel__chart">
        {temperatures.length > 0 ? (
          <TemperatureChart temperatures={temperatures} />
        ) : (
          'No data available'
        )}
      </div>
    </div>
  );
};

TemperaturePanel.defaultProps = {
  temperatures: [],
};

export default TemperaturePanel;
