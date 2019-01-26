import React from 'react';
import { withRouter } from 'react-router';

class AgentEdit extends React.Component {
  constructor(props) {
    super(props);
    this.updateTemperature = this.updateTemperature.bind(this);
  }

  updateTemperature() {
    const {
      agentID,
      updateConfig,
    } = this.props;

    updateConfig(agentID, {
      temperature: 3,
    });
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
        <label className="c-input__label">Temperature modifier</label>
        <input className="c-input__field"/>
      </div>
      <button onClick={this.updateTemperature}>
        UPDATE
      </button>
    </div>);
  }
}

export default withRouter(AgentEdit);
