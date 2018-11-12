import { connect } from 'react-redux';
import Type2 from './Type2';

const mapStateToProps = (state, ownProps) => {
  const { agent } = ownProps;
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
