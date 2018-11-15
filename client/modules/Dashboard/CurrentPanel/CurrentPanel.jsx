import React from 'react';
import PropTypes from 'prop-types';

const CurrentPanel = (props) => {
  const {
    isMotion,
    isGas,
  } = props;

  return (<div className="current-panel">
    <div className="current-panel__title">
      Current state <span className="current-panel__annotate">(~3 mins back)</span>
    </div>
    <div className="current-panel__content">
      <div className="state">
        <div className="state__item">
          <div className="state__status">
            <div className={isMotion ? 'state__alert' : 'state__ok'} />
          </div>
          <div className="state__name">
            Motion
          </div>
        </div>
        <div className="state__item">
          <div className="state__status">
            <div className={isGas ? 'state__alert' : 'state__ok'} />
          </div>
          <div className="state__name">
            Combustible gases
          </div>
        </div>
      </div>
    </div>
  </div>);
};

CurrentPanel.propTypes = {
  isGas: PropTypes.bool,
  isMotion: PropTypes.bool,
};

CurrentPanel.defaultProps = {
  isGas: false,
  isMotion: false,
};

export default CurrentPanel;
