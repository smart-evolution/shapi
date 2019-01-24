import React from 'react';
import { withRouter } from 'react-router';


const AgentEdit = (props) => {
  return (<div>
    <div class="c-input">
      <label class="c-input__label">Temperature modifier</label>
      <input class="c-input__field"/>
    </div>
  </div>);
}

export default withRouter(AgentEdit);
