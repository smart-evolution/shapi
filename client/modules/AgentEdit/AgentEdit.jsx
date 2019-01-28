import React from 'react';
import PropTypes from 'prop-types';
import { withRouter } from 'react-router';

class AgentEdit extends React.Component {
  constructor(props) {
    super(props);
    this.updateTemperature = this.updateTemperature.bind(this);
    this.updateConfig = this.updateConfig.bind(this);
  }

  componentDidMount() {
    const {
      agentID,
      fetchConfig,
    } = this.props;

    fetchConfig(agentID);
  }

  updateConfig() {
    const {
      agentID,
      config,
      updateConfig,
    } = this.props;

    updateConfig(agentID, config);
  }

  updateTemperature(e) {
    const temperature = e.target.value;
    const {
      agentID,
      updateTemperature,
    } = this.props;

    updateTemperature(agentID, temperature);
  }

  render() {
    const {
      config: {
        temperature,
      },
    } = this.props;

    return (<div>
      <div className="c-input">
        <div className="c-input__label">
          Temperature modifier
        </div>
        <input
          className="c-input__field"
          value={temperature}
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

AgentEdit.propTypes = {
  config: PropTypes.shape(),
  agentID: PropTypes.string,
  updateTemperature: PropTypes.func,
  fetchConfig: PropTypes.func,
  updateConfig: PropTypes.func,
};

export default withRouter(AgentEdit);
