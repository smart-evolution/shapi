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
  static onDragOver(event: SyntheticDragEvent<HTMLDivElement>) {
    event.preventDefault();
  }

  constructor() {
    super();
    this.state = {
      left: 0,
      top: 0,
    };
  }

  onMove(event: SyntheticDragEvent<Element>) {
    event.preventDefault();
    const { clientX, clientY } = event;
    this.move(clientX, clientY);
  }

  onDrop(event: SyntheticDragEvent<Element>) {
    event.preventDefault();
    const { clientX, clientY } = event;
    this.move(clientX, clientY);
  }

  move(x: number, y: number) {
    /* eslint-disable react/no-find-dom-node */
    const node: null | Element | Text = ReactDOM.findDOMNode(this);
    /* eslint-enable react/no-find-dom-node */
    if (node === null || node instanceof Text) {
      return;
    }

    const rect: ClientRect = node.getBoundingClientRect();
    const left = _.min([
      _.max([x - rect.left - JOYSTICK_RADIUS * 0.5, 0]),
      JOYSTICK_RADIUS,
    ]);
    const top = _.min([
      _.max([y - rect.top - JOYSTICK_RADIUS * 0.5, 0]),
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
          onDrag={(event: SyntheticDragEvent<HTMLDivElement>) => {
            this.onMove(event);
          }}
          onDrop={(event: SyntheticDragEvent<HTMLDivElement>) => {
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
