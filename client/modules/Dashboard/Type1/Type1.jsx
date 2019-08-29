// @flow
import React from 'react';
import * as agentsTypes from 'client/models/agents/types';
import Scroller from 'client/components/Scroller';
import TemperaturePanel from '../TemperaturePanel';
import SoundPanel from '../SoundPanel';
import CurrentPanel from '../CurrentPanel';

type Props = {
  pathname: string,
  agent: agentsTypes.Agent,
  onScroll: () => void,
};

const Type1 = (props: Props) => {
  const { agent, pathname, onScroll } = props;

  return (
    <>
      <div className="dashboard__cell dashboard__cell--full">
        <a className="c-btn c-btn--edit" href={`${pathname}/edit`}>
          Edit
        </a>
      </div>
      <div className="dashboard__cell dashboard__cell--full">
        <Scroller onScrollChange={onScroll} min={30} max={300} />
      </div>
      <div className="dashboard__cell dashboard__cell--full">
        <CurrentPanel agent={agent} />
      </div>
      <div className="dashboard__cell dashboard__cell--full">
        <TemperaturePanel agent={agent} />
      </div>
      <div className="dashboard__cell dashboard__cell--full">
        <SoundPanel agent={agent} />
      </div>
    </>
  );
};

export default Type1;
