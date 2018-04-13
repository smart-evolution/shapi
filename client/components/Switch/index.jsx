import _ from 'lodash';
import React from 'react';
import PropTypes from 'prop-types';

const Switch = (props) => {
  const { isOn, onToggle, className } = props;

  return (
    <div
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
    </div>);
}

Switch.propTypes = {
  isOn: PropTypes.bool,
  onToggle: PropTypes.func,
};

Switch.defaultProps = {
  isOn: false,
  onToggle: _.noop,
};

export default Switch;

