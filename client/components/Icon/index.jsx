import React from 'react';
import PropTypes from 'prop-types';

const Icon = (props) => {
  const { type, className } = props;

  const classes = `c-icon c-icon--${type} ${className}`;

  return (<span className={classes} />);
};

Icon.propTypes = {
  className: PropTypes.string,
  type: PropTypes.string,
};

export default Icon;

