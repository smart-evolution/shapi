import React from 'react';
import { mount } from 'enzyme';
import Icon from './index';
import '../../../setupTest';

describe('Icon', () => {
  it('should render correctly', () => {
    const component = mount(<Icon />);

    expect(component).toMatchSnapshot();
  });
});
