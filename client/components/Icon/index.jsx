// @flow
import React from 'react';

type Props = {
  type: string,
  className: string,
};

const Icon = (props: Props) => {
  const { type, className } = props;

  const classes = `c-icon c-icon--${type} ${className}`;

  return (<span className={classes} />);
};

Icon.defaultProps = {
  className: '',
};

export default Icon;

