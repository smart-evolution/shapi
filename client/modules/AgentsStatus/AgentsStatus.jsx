// @flow
import React from 'react';
import ControlPanel from './ControlPanel';
import List from './List';

type Props = {
  error: string,
};

const AgentsStatus = (props: Props) => {
  const { error } = props;

  return (
    <div className="agents-status">
      {error && <div className="agents-status__error">{error}</div>}
      <ControlPanel />
      <List />
    </div>
  );
};

AgentsStatus.defaultProps = {
  error: '',
};

export default AgentsStatus;
