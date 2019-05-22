import { connect } from 'react-redux';
import Jeep from './Jeep';

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
)(Jeep);
