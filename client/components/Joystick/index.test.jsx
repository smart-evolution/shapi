import _ from 'lodash';
import React from 'react';
import { mount } from 'enzyme';
import Joystick from './index';

describe('Joystick', () => {
  it('should render correctly', () => {
    const component = mount(<Joystick onPositionChange={_.noop} />);

    expect(component).toMatchSnapshot();
  });

  it('should render correctly in disabled state', () => {
    const component = mount(<Joystick onPositionChange={_.noop} isEnabled={false} />);

    expect(component).toMatchSnapshot();
  });
});
