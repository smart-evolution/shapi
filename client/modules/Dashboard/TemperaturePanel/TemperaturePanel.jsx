import React from 'react';
import PropTypes from 'prop-types';
import TemperatureChart from './TemperatureChart/TemperatureChart';

const NODATA_SIGN = '-';

const TemperaturePanel = (props) => {
  const { temperatures } = props;
  const nowTmp = _.head(temperatures);
  const value = _.isUndefined(nowTmp) ? NODATA_SIGN : nowTmp.value;

   return (<div className="temperature-panel">
      <div className="temperature-panel__title">
       Temperature
     </div>
     <div className="temperature-panel__current">
        {value} &#8451;
      </div>
      <div className="temperature-panel__chart">
        <TemperatureChart
          temperatures={temperatures}
        />
      </div>
    </div>);
}

TemperaturePanel.propTypes = {
  temperatures: PropTypes.array,
};

TemperaturePanel.defaultProps = {
  temperatures: [],
};

export default TemperaturePanel;
