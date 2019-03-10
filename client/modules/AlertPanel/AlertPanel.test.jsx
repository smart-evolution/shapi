import React from 'react';
import { mount } from 'enzyme';
import AlertPanel from './AlertPanel';

describe('AlertPanel/AlertPanel', () => {
  it('should render correctly', () => {
    const alerts = [
      {
        type: 'type-class',
        message: 'Message 1',
      },
      {
        type: 'type-class',
        message: 'Message 2',
      },
      {
        type: 'type-class',
        message: 'Message 3',
      },
    ];

    const component = mount(<AlertPanel alerts={alerts} />);

    expect(component).toMatchSnapshot();
  });
});
