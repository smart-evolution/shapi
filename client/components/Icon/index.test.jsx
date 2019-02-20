import React from 'react';
import { mount } from 'enzyme';
import Icon from './index';

describe('Icon', () => {
  it('should render correctly', () => {
    const component = mount(<Icon />);

    expect(component).toMatchSnapshot();
  });
});
