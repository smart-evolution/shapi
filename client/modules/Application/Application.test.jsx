import _ from 'lodash';
import React from 'react';
import { shallow } from 'enzyme';
import Application from './Application';

describe('Application', () => {
  it('should render with loader', () => {
    const component = shallow(
      <Application isLoaded={false} mount={_.noop}>
        Content
      </Application>
    );

    expect(component).toMatchSnapshot();
  });

  it('should render with children', () => {
    const component = shallow(
      <Application isLoaded mount={_.noop}>
        <p>Some content</p>
      </Application>
    );

    expect(component).toMatchSnapshot();
  });
});
