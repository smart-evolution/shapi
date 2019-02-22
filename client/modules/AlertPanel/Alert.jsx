import _ from 'lodash';
import React from 'react';


class Alert extends React.Component {
  constructor() {
    super();

    this.state = {
      show: false,
      hide: false,
    };
  }

  componentDidMount () {
    this.showTimeout = setTimeout(function () {
      this.setState({ show: true });
    }.bind(this), 2000);

    this.hideTimeout = setTimeout(function () {
      this.setState({ hide: true });
    }.bind(this), 4000);
  }

  componentWillUnmount () {
    if (this.showTimeout) {
      clearTimeout(this.showTimeout);
    }
    if (this.hideTimeout) {
      clearTimeout(this.hideTimeout);
    }
  }

  render() {
    const {
      type,
      children,
    } = this.props;
    const {
      show,
      hide,
    } = this.state;

    const showClass = show ? 'show' : '';
    const hideClass = hide ? 'hide' : '';
    const classes = `alert-panel__alert alert-panel__alert--${type} ${showClass} ${hideClass}`;

    return (
      <div
        className={classes}
      >
        {children}
      </div>
    );
  }
}

export default Alert;
