import React from 'react';
import PropTypes from 'prop-types';
import TemperatureChart from './TemperatureChart/TemperatureChart';

const NODATA_SIGN = '-';

const TemperaturePanel = (props) => {
  const { temperature } = props;
  const lastTmp = _.last(temperature);
  const currentTmp = _.isUndefined(lastTmp) ? NODATA_SIGN : lastTmp.value;

   return (<div className="temperature-panel">
      <div className="temperature-panel__title">
       Temperature
     </div>
     <div className="temperature-panel__current">
        {currentTmp} &#8451;
      </div>
      <div className="temperature-panel__chart">
        <TemperatureChart
          temperature={temperature}
        />
      </div>
    </div>);
}

TemperaturePanel.propTypes = {
  time: PropTypes.array,
  temperature: PropTypes.array,
};

TemperaturePanel.defaultProps = {
  time: [],
  temperature: [],
};

export default TemperaturePanel;
