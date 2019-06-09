// @flow
import _ from 'lodash';
import React from 'react';
import ReactDOM from 'react-dom';

type Props = {
  onPositionChange: (left: number, top: number) => void,
  isEnabled: boolean,
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

  static defaultProps = {
    isEnabled: true,
  };

  constructor() {
    super();
    this.state = {
      left: JOYSTICK_RADIUS * 0.5,
      top: JOYSTICK_RADIUS * 0.5,
    };
  }

  onMove(event: SyntheticDragEvent<Element>) {
    event.preventDefault();
    const { clientX, clientY } = event;

    /* eslint-disable react/no-find-dom-node */
    const node: null | Element | Text = ReactDOM.findDOMNode(this);
    /* eslint-enable react/no-find-dom-node */
    if (node === null || node instanceof Text) {
      return;
    }
    const rect: ClientRect = node.getBoundingClientRect();
    const mappedX = clientX - rect.left - JOYSTICK_RADIUS * 0.5;
    const mappedY = clientY - rect.top - JOYSTICK_RADIUS * 0.5;

    this.move(mappedX, mappedY);
  }

  onDragEnd(event: SyntheticDragEvent<Element>) {
    event.preventDefault();
    this.move(JOYSTICK_RADIUS * 0.5, JOYSTICK_RADIUS * 0.5);
  }

  move(x: number, y: number) {
    const left = _.min([_.max([x, 0]), JOYSTICK_RADIUS]);
    const top = _.min([_.max([y, 0]), JOYSTICK_RADIUS]);

    const { onPositionChange } = this.props;

    onPositionChange(left, top);
    this.setState({
      left,
      top,
    });
  }

  render() {
    const { left, top } = this.state;
    const { isEnabled } = this.props;

    const classes = `c-joystick${isEnabled ? '' : ' c-joystick--inactive'}`;

    return (
      <div
        className={classes}
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
          draggable={isEnabled}
          onDrag={(event: SyntheticDragEvent<HTMLDivElement>) => {
            this.onMove(event);
          }}
          onDragEnd={(event: SyntheticDragEvent<HTMLDivElement>) => {
            this.onDragEnd(event);
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
