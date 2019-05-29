// @flow
import _ from 'lodash';
import React from 'react';
import ReactDOM from 'react-dom';

type Props = {
  onPositionChange: (left: number, top: number) => void,
};

type State = {
  left: number,
  top: number,
};

const JOYSTICK_RADIUS = 50;

class Joystick extends React.PureComponent<Props, State> {
  static onDragOver(event) {
    event.preventDefault();
  }

  constructor() {
    super();
    this.state = {
      left: 0,
      top: 0,
    };
  }

  onMove(event) {
    event.preventDefault();
    const { clientX, clientY } = event;
    /* eslint-disable react/no-find-dom-node */
    const rect = ReactDOM.findDOMNode(this).getBoundingClientRect();
    /* eslint-enable react/no-find-dom-node */
    const left = _.min([
      _.max([clientX - rect.left - JOYSTICK_RADIUS * 0.5, 0]),
      JOYSTICK_RADIUS,
    ]);
    const top = _.min([
      _.max([clientY - rect.top - JOYSTICK_RADIUS * 0.5, 0]),
      JOYSTICK_RADIUS,
    ]);

    const { onPositionChange } = this.props;

    onPositionChange(left, top);
    this.setState({
      left,
      top,
    });
  }

  onDrop(event) {
    event.preventDefault();
    const { clientX, clientY } = event;
    /* eslint-disable react/no-find-dom-node */
    const rect = ReactDOM.findDOMNode(this).getBoundingClientRect();
    /* eslint-enable react/no-find-dom-node */
    const left = _.min([
      _.max([clientX - rect.left - JOYSTICK_RADIUS * 0.5, 0]),
      JOYSTICK_RADIUS,
    ]);
    const top = _.min([
      _.max([clientY - rect.top - JOYSTICK_RADIUS * 0.5, 0]),
      JOYSTICK_RADIUS,
    ]);

    const { onPositionChange } = this.props;

    onPositionChange(left, top);
    this.setState({
      left,
      top,
    });
  }

  render() {
    const { left, top } = this.state;

    return (
      <div
        className="c-joystick"
        onDragOver={event => Joystick.onDragOver(event)}
      >
        <div
          className="c-joystick__knob"
          style={{
            left,
            top,
          }}
        />
        <div
          className="c-joystick__drag"
          draggable="true"
          onDrag={event => {
            this.onMove(event);
          }}
          onDrop={event => {
            this.onDrop(event);
          }}
          style={{
            left,
            top,
          }}
        />
      </div>
    );
  }
}

export default Joystick;
