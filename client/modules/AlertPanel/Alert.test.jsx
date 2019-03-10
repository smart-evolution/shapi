import React from 'react';
import { mount } from 'enzyme';
import Alert from './Alert';

describe('AlertPanel/Alert', () => {
  it('should render correctly', () => {
    const component = mount(<Alert type="type-class">Message</Alert>);

    expect(component).toMatchSnapshot();
  });
});
