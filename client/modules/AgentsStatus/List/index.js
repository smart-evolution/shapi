import { connect } from 'react-redux';
import _ from 'lodash';
import List from './List';

const mapStateToProps = state => {
  const {
    agents: { agents, isLoading },
  } = state;

  return {
    agents: _.concat(agents, [
      {
        id: 'jeepjimmy',
        name: 'Jimmy',
        data: {},
        type: 'jeep',
      },
    ]),
    isLoading,
  };
};

const mapDispatchToProps = () => ({});

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(List);
