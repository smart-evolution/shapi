import _ from 'lodash';
import React from 'react';
import PropTypes from 'prop-types';
import SoundChart from './SoundChart/SoundChart';

const NODATA_SIGN = '-';

const SoundPanel = props => {
  const { sounds } = props;
  const nowSnd = _.head(sounds);
  const value = _.isUndefined(nowSnd) ? NODATA_SIGN : nowSnd.value;

  return (
    <div className="sound-panel">
      <div className="sound-panel__title">Sound</div>
      <div className="sound-panel__current">{value} db</div>
      <div className="sound-panel__chart">
        {sounds.length > 0 ? (
          <SoundChart sounds={sounds} />
        ) : (
          'No data available'
        )}
      </div>
    </div>
  );
};

SoundPanel.propTypes = {
  sounds: PropTypes.arrayOf(PropTypes.string),
};

SoundPanel.defaultProps = {
  sounds: [],
};

export default SoundPanel;
