import React from 'react';
import PropTypes from 'prop-types';
import Switch from 'components/Switch';

const ControlPanel = props => {
  const { isAlerts, onToggle, sendAlert } = props;

  return (
    <div className="control-panel">
      <div className="control-panel__title">Control Panel</div>
      <div className="control-panel__dashboard">
        <div className="control-panel__control">
          Alerts
          <Switch
            className="control-panel__alerts"
            isOn={isAlerts}
            onToggle={onToggle}
          />
        </div>
        <div className="control-panel__control">
          Send alert
          <button className="control-panel__send-alert" onClick={sendAlert} />
        </div>
      </div>
    </div>
  );
};

ControlPanel.propTypes = {
  isAlerts: PropTypes.bool,
  onToggle: PropTypes.func,
  sendAlert: PropTypes.func,
};

export default ControlPanel;
