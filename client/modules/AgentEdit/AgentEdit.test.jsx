import React from 'react';
import { mount } from 'enzyme';
import { MemoryRouter as Router } from 'react-router';
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
    const component = mount(
      <Router keyLength={0}>
        <AgentEdit agent={agent} agentConfig={agentConfig} />
      </Router>
    );

    expect(component).toMatchSnapshot();
  });
});
