import { connect } from 'react-redux';
import Type2 from './Type2';

const mapStateToProps = (state) => {
  const { agent } = state;
  const { id, name } = agent;

  return {
    id,
    name,
  };
};

const mapDispatchToProps = () => ({});

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(Type2);
