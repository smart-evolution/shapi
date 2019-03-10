// @flow
import _ from 'lodash';
import React from 'react';

type Props = {
  isOn: boolean,
  onToggle: () => void,
  className: string,
};

const Switch = (props: Props) => {
  const { isOn, onToggle, className } = props;

  return (
    <button
      role="button"
      className={`c-switch ${className}`}
      onClick={onToggle}
    >
      <input
        className="c-switch__input"
        type="checkbox"
        checked={isOn}
        onChange={_.noop}
      />
      <span className="c-switch__slider" />
    </button>
  );
};

Switch.defaultProps = {
  isOn: false,
  onToggle: _.noop,
};

export default Switch;
