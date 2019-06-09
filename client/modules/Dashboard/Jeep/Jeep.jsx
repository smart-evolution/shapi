// @flow
import React from 'react';
import { withRouter } from 'react-router';
import * as agentsTypes from 'client/models/agents/types';
import Joystick from 'client/components/Joystick';
import Switch from 'client/components/Switch';

type Props = {
  agent: agentsTypes.Agent,
  onPositionChange: (agentsTypes.Agent, { left: number, top: number }) => void,
  onToggle: () => void,
  setup: () => void,
  isDevConnected: boolean,
};

class Jeep extends React.PureComponent<Props> {
  constructor(props) {
    super();
    const { setup, agent } = props;
    setup(agent);
  }

  render() {
    const { agent, onPositionChange, onToggle, isDevConnected } = this.props;

    return (
      <div className="jeep-panel">
        <div className="jeep-panel__section">
          <div className="c-control">
            Device connection
            <div className="c-control__content">
              <Switch className="" isOn={isDevConnected} onToggle={onToggle} />
            </div>
          </div>
        </div>
        <div className="jeep-panel__section">
          <Joystick
            isEnabled={isDevConnected}
            onPositionChange={(left: number, top: number) => {
              onPositionChange(agent, { left, top });
            }}
          />
        </div>
      </div>
    );
  }
}

export default withRouter(Jeep);
