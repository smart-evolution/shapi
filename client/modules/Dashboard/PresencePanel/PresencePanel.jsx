import React from 'react';
import PropTypes from 'prop-types';

const PresencePanel = (props) => {
  const { presence } = props;

  return (<div className="presence-panel">
    <div className="presence-panel__current">
      {presence}
    </div>
  </div>);
}

PresencePanel.propTypes = {
  presence: PropTypes.string,
};

PresencePanel.defaultProps = {
  presence: "No motion",
};

export default PresencePanel;
