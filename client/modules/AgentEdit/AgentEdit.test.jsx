// @flow
import _ from 'lodash';
import React from 'react';
import { shallow } from 'enzyme';
import AgentEdit from './AgentEdit';

const agent = {
  id: '90346453',
  name: 'Livingroom',
  data: {},
  type: 'type1-v0.3.3',
  ip: '192.168.1.10',
  isOnline: true,
};

const agentConfig = {
  id: '90346453',
  name: 'Livingroom',
  temperature: 0,
};

describe('AgentEdit', () => {
  it('should render correctly', () => {
    const component = shallow(
      <AgentEdit
        agent={agent}
        agentConfig={agentConfig}
      />
    );

    expect(component).toMatchSnapshot();
  });
});
