import _ from 'lodash';
import React from 'react';
import { mount } from 'enzyme';
import Scroller from './index';

describe('Scroller', () => {
  it('should render correctly', () => {
    const component = mount(
      <Scroller onScrollChange={_.noop} min={10} max={100} />
    );

    expect(component).toMatchSnapshot();
  });
});
