import _ from 'lodash';
import React from 'react';
import PropTypes from 'prop-types';
import Type1 from './Agents/Type1';
import Type2 from './Agents/Type2';

const List = props => {
  const { agents } = props;

  const loader = <div className="c-loader" />;
  const list = (
    <ul className="agents-list__list">
      {_.map(agents, agent => {
        switch (agent.type) {
          case 'type1':
            return <Type1 key={agent.id} agent={agent} />;

          case 'type2':
            return <Type2 key={agent.id} agent={agent} />;

          default:
            return null;
        }
      })}
    </ul>
  );

  return (
    <div className="agents-list">
      <div className="agents-list__title">Available agents</div>
      {_.isEmpty(agents) ? loader : list}
    </div>
  );
};

List.propTypes = {
  agents: PropTypes.arrayOf(PropTypes.object),
};

export default List;
