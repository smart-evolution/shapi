// @flow
import _ from 'lodash';
import React from 'react';

type Props = {
  type: string,
  className: string,
};

const Icon = (props: Props) => {
  const { type, className } = props;
  const typeClass = _.isEmpty(type) ? '' : `c-icon--${type}`;

  const classes = `c-icon ${typeClass} ${className}`;

  return <span className={classes} />;
};

Icon.defaultProps = {
  className: '',
};

export default Icon;
