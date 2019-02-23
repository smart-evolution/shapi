import _ from 'lodash';
import React from 'react';
import PropTypes from 'prop-types';
import Alert from './Alert';

const AlertPanel = (props) => {
  const { alerts } = props;

  return (
    <div className="alert-panel">
      {_.map(alerts, (alert, index) => {
        const {
          type,
          message,
        } = alert;
        const key = `alert-${index}`;

        return (
          <Alert
            key={key}
            type={type}
          >
            {message}
          </Alert>
        );
      })}
    </div>
  );
};

AlertPanel.propTypes = {
  alerts: PropTypes.arrayOf(PropTypes.object),
};

export default AlertPanel;
