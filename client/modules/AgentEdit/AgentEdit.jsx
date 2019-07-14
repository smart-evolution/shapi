// @flow
import React from 'react';
import { withRouter } from 'react-router';
import * as agentsConstants from 'client/models/agents/constants';
import * as agentTypes from 'client/models/agents/types';
import * as agentQueries from 'client/models/agents/queries';
import * as agentConfigTypes from 'client/models/agentConfigs/types';

type Props = {|
  agent: agentTypes.Agent,
  agentConfig: agentConfigTypes.AgentConfig,
  commitConfig: (agentTypes.AgentID, agentConfigTypes.AgentConfig) => void,
  updateProperty: (agentTypes.AgentID, string, string) => void,
|};

class AgentEdit extends React.Component<Props> {
  constructor(props) {
    super(props);
    (this: any).updateTemperature = this.updateTemperature.bind(this);
    (this: any).updateName = this.updateName.bind(this);
    (this: any).updateConfig = this.updateConfig.bind(this);
  }

  updateConfig() {
    const { agent, agentConfig, commitConfig } = this.props;
    commitConfig(agent.id, agentConfig);
  }

  updateTemperature(e) {
    const value = e.target.value;
    const { agent, updateProperty } = this.props;
    updateProperty(agent.id, 'temperature', value);
  }

  updateName(e) {
    const value = e.target.value;
    const { agent, updateProperty } = this.props;
    updateProperty(agent.id, 'name', value);
  }

  render() {
    const { agentConfig, agent } = this.props;

    const rawType = agentQueries.getNoVersionedType(agent);

    const temperatureAdjustment = (
      <div className="c-input c-input__full">
        <div className="c-input__label">Temperature modifier</div>
        <input
          className="c-input__field"
          value={agentConfig.temperature || ''}
          onChange={this.updateTemperature}
        />
      </div>
    );

    return (
      <div>
        <div className="c-input c-input__full">
          <div className="c-input__label">Name</div>
          <input
            className="c-input__field"
            value={agentConfig.name || ''}
            onChange={this.updateName}
          />
        </div>
        {rawType === agentsConstants.Type1 && temperatureAdjustment}
        <button
          className="c-btn c-btn--full c-btn--accept"
          onClick={this.updateConfig}
        >
          UPDATE
        </button>
      </div>
    );
  }
}

export default withRouter(AgentEdit);
