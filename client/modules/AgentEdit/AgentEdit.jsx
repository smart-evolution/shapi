// @flow
import React from 'react';
import { withRouter } from 'react-router';

type Props = {
  config: Object,
  agentID: string,
  fetchConfig: string => void,
  updateConfig: (string, Object) => void,
  updateTemperature: (string, string) => void,
};

class AgentEdit extends React.Component<Props> {
  constructor(props) {
    super(props);
    (this: any).updateTemperature = this.updateTemperature.bind(this);
    (this: any).updateConfig = this.updateConfig.bind(this);
  }

  componentDidMount() {
    const { agentID, fetchConfig } = this.props;

    fetchConfig(agentID);
  }

  updateConfig() {
    const { agentID, config, updateConfig } = this.props;

    updateConfig(agentID, config);
  }

  updateTemperature(e) {
    const temperature = e.target.value;
    const { agentID, updateTemperature } = this.props;

    updateTemperature(agentID, temperature);
  }

  render() {
    const {
      config: { temperature },
    } = this.props;

    return (
      <div>
        <div className="c-input">
          <div className="c-input__label">Temperature modifier</div>
          <input
            className="c-input__field"
            value={temperature}
            onChange={this.updateTemperature}
          />
        </div>
        <button className="c-btn c-btn--accept" onClick={this.updateConfig}>
          UPDATE
        </button>
      </div>
    );
  }
}

export default withRouter(AgentEdit);
