import _ from 'lodash';
import moment from 'moment';
import React from 'react';
import PropTypes from 'prop-types';

const DATE_FORMAT = 'MMMM Do YYYY, h:mm:ss a';

const PresencePanel = (props) => {
  const { presence, motions } = props;

  const motionLog = _.reverse(_.sortBy(motions));

  return (<div className="presence-panel">
    <div className="presence-panel__title">
      Motion
    </div>
    <div className="presence-panel__current">
      {presence}
    </div>
    <ul className="presence-panel__history">
      {_.map(motionLog, (motion, key) => (
        <li key={key}>
          {moment(motion).format(DATE_FORMAT)}
        </li>
      ))}
    </ul>
  </div>);
}

PresencePanel.propTypes = {
  presence: PropTypes.string,
  motions: PropTypes.object,
};

PresencePanel.defaultProps = {
  presence: "No motion",
};

export default PresencePanel;
