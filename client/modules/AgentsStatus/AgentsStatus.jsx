import React from 'react';
import PropTypes from 'prop-types';
import ControlPanel from './ControlPanel';
import List from './List';

const AgentsStatus = (props) => {
  const { error } = props;

  return (
    <div className="agents-status">
      { error && (
        <div className="agents-status__error">
          {error}
        </div>
      )}
      <ControlPanel />
      <List />
    </div>
  );
};

AgentsStatus.defaultProps = {
  error: '',
};

AgentsStatus.propTypes = {
  error: PropTypes.string,
};

export default AgentsStatus;
