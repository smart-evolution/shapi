import { connect } from 'react-redux';
import List from './List';

const mapStateToProps = (state) => {
  const { agents } = state;

  return {
    agents,
  };
};

const mapDispatchToProps = () => ({});

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(List);
