import React from 'react';
import PropTypes from 'prop-types';

class Alert extends React.Component {
  constructor() {
    super();

    this.state = {
      show: false,
      hide: false,
    };
  }

  componentDidMount() {
    this.showTimeout = setTimeout(() => {
      this.setState({ show: true });
    }, 2000);

    this.hideTimeout = setTimeout(() => {
      this.setState({ hide: true });
    }, 4000);
  }

  componentWillUnmount() {
    if (this.showTimeout) {
      clearTimeout(this.showTimeout);
    }
    if (this.hideTimeout) {
      clearTimeout(this.hideTimeout);
    }
  }

  render() {
    const { type, children } = this.props;
    const { show, hide } = this.state;

    const showClass = show ? 'show' : '';
    const hideClass = hide ? 'hide' : '';
    const classes = `alert-panel__alert alert-panel__alert--${type} ${showClass} ${hideClass}`;

    return <div className={classes}>{children}</div>;
  }
}

Alert.propTypes = {
  type: PropTypes.string,
  children: PropTypes.oneOfType([
    PropTypes.arrayOf(PropTypes.node),
    PropTypes.node,
  ]).isRequired,
};

export default Alert;
