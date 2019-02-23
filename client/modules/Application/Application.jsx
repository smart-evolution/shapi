import React from 'react';
import PropTypes from 'prop-types';
import AlertPanel from 'modules/AlertPanel';

const Application = props => (
  <div className="application">
    {props.children}
    <AlertPanel />
  </div>
);

Application.propTypes = {
  children: PropTypes.oneOfType([
    PropTypes.arrayOf(PropTypes.node),
    PropTypes.node
  ]).isRequired,
};

export default Application;
