import _ from 'lodash';
import moment from 'moment';
import React from 'react';
import PropTypes from 'prop-types';

const DATE_FORMAT = 'MMMM Do YYYY, h:mm:ss a';

const PresencePanel = (props) => {
  const { motions } = props;

  return (<div className="presence-panel">
    <div className="presence-panel__title">
      Motion
    </div>
    <ul className="presence-panel__history">
      {_.map(motions, (motion, key) => (
        <li key={key}>
          {moment(motion.time).format(DATE_FORMAT)}
        </li>
      ))}
    </ul>
  </div>);
}

PresencePanel.propTypes = {
  motions: PropTypes.object,
};

PresencePanel.defaultProps = {
  motions: {},
};

export default PresencePanel;
