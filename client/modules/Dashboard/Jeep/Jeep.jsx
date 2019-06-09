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
      <div className="dashboard__cell dashboard__cell--full">
        <div className="c-control">
          Device connection
          <div className="c-control__content">
            <Switch
              className=""
              isOn={isDevConnected}
              onToggle={onToggle}
            />
          </div>
        </div>
        <Joystick
          isEnabled={isDevConnected}
          onPositionChange={(left: number, top: number) => {
            onPositionChange(agent, { left, top });
          }}
        />
      </div>
    );
  }
}

export default withRouter(Jeep);
