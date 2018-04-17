import _ from 'lodash';
import moment from 'moment';
import React from 'react';
import PropTypes from 'prop-types';

const GasPanel = (props) => {
  const { gas } = props;

  return (<div className="presence-panel">
    <div className="presence-panel__title">
      Combustible gases <div className={gas ? 'c-led-red' : 'c-led-green'}></div>
    </div>
  </div>);
}

GasPanel.propTypes = {
  gas: PropTypes.bool,
};

GasPanel.defaultProps = {
  gas: false,
};

export default GasPanel;
