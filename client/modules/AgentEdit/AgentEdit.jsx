import React from 'react';
import { withRouter } from 'react-router';

class AgentEdit extends React.Component {
  constructor(props) {
    super(props);
    this.updateTemperature = this.updateTemperature.bind(this);
    this.updateConfig = this.updateConfig.bind(this);
  }

  updateTemperature(e) {
    const temperature = e.target.value;
    const {
      agentID,
      updateTemperature,
    } = this.props;

    updateTemperature(agentID, temperature);
  }

  updateConfig() {
    const {
      agentID,
      config,
      updateConfig,
    } = this.props;

    updateConfig(agentID, config);
  }

  componentDidMount() {
    const {
      agentID,
      fetchConfig,
    } = this.props;

    fetchConfig(agentID);
  }

  render() {
    const {
      config: {
        temperature,
      },
    } = this.props;

    return (<div>
      <div className="c-input">
        [{temperature}]
        <label className="c-input__label">
          Temperature modifier
        </label>
        <input
          className="c-input__field"
          onChange={this.updateTemperature}
        />
      </div>
      <button
        className="c-btn c-btn--accept"
        onClick={this.updateConfig}
      >
        UPDATE
      </button>
    </div>);
  }
}

export default withRouter(AgentEdit);
