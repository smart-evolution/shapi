import React from 'react';
import { shallow } from 'enzyme';
import Application from './Application';

describe('Application', () => {
  it('should render correctly', () => {
    const component = shallow(
      <Application>
        Content
      </Application>
    );

    expect(component).toMatchSnapshot();
  });
});
