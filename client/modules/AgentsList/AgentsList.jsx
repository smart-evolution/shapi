import React from 'react';
import PropTypes from 'prop-types';

const AgentsList = (props) => {
  const { agents } = props;

  return (
    <div className="agents-list">
      {JSON.stringify(agents)}
    </div>
  );
};

AgentsList.propTypes = {
  agents: PropTypes.array,
};

export default AgentsList;
