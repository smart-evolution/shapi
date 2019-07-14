// @flow
import React from 'react';
import AlertPanel from 'client/modules/AlertPanel';

type Props = {
  isLoaded: boolean,
  children: React$Element<any>,
  mount: () => void,
};

class Application extends React.PureComponent<Props> {
  componentDidMount() {
    const { mount } = this.props;
    mount();
  }

  render() {
    const { children, isLoaded } = this.props;

    const loader = <div className="c-loader" />;

    return (
      <div className="application">
        {isLoaded && children}
        {!isLoaded && loader}
        <AlertPanel />
      </div>
    );
  }
}

export default Application;
