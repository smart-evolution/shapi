// @flow
import React from 'react';
import { withRouter } from 'react-router';
import agentsTypes from 'models/agents/types';
import Joystick from 'components/Joystick';

type Props = {
  agent: agentsTypes.Agent,
  onPositionChange: (agentsTypes.Agent, string) => void,
  setup: agentsTypes.Agent => void,
};

class Jeep extends React.PureComponent<Props> {
  constructor(props) {
    super();
    const { setup, agent } = props;
    setup(agent);
  }

  render() {
    const { agent, onPositionChange } = this.props;

    return (
      <div className="dashboard__cell dashboard__cell--full">
        <Joystick
          onPositionChange={(left: number, top: number) => {
            onPositionChange(agent, { left, top });
          }}
        />
      </div>
    );
  }
}

export default withRouter(Jeep);
