import React from 'react';
import AlertPanel from 'modules/AlertPanel';

const Application = (props) => (
  <div className="application">
    {props.children}
    <AlertPanel />
  </div>
);

export default Application;


