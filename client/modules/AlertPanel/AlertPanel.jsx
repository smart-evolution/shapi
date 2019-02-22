import _ from 'lodash';
import React from 'react';
import Alert from './Alert';

class AlertPanel extends React.Component {
  render() {
    const { alerts } = this.props;

    return (
      <div className="alert-panel">
        {_.map(alerts, (alert, index) => {
          const { type, message } = alert;
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
  }
}

export default AlertPanel;
