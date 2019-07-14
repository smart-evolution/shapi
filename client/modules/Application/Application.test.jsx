import _ from 'lodash';
import React from 'react';
import { shallow } from 'enzyme';
import Application from './Application';

describe('Application', () => {
  it('should render correctly', () => {
    const component = shallow(
      <Application mount={_.noop}>Content</Application>
    );

    expect(component).toMatchSnapshot();
  });
});
