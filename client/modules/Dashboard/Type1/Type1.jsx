// @flow
import React from 'react';
import { Scroller } from 'graphen';
import * as agentsTypes from 'client/models/agents/types';
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
        <div className="c-panel c-panel--separator">
          <div className="c-panel__content">
            <a className="c-btn c-btn--edit" href={`${pathname}/edit`}>
              Edit
            </a>
          </div>
        </div>
      </div>
      <div className="dashboard__cell dashboard__cell--full">
        <div className="c-panel c-panel--separator">
          <div className="c-panel__content">
            <Scroller onScrollChange={onScroll} min={30} max={300} />
          </div>
        </div>
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
