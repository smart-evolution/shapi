import React from 'react';
import { withRouter } from 'react-router';

class AgentEdit extends React.Component {
  componentDidMount() {

    console.log(this.props.match.params.agent);
    console.log('--- componentDidMount')
    this.props.fetchConfig(this.props.match.params.agent);
  }

  render(props) {
    return (<div>
      <div className="c-input">
        <label className="c-input__label">Temperature modifier</label>
        <input className="c-input__field"/>
      </div>
    </div>);
  }
}

export default withRouter(AgentEdit);
