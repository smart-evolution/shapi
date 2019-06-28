import { connect } from 'react-redux';
import Jeep from './Jeep';

const mapStateToProps = (state, ownProps) => {
  const { agent } = ownProps;
  const { id, name, isOnline } = agent;

  return {
    id,
    name,
    isOnline,
  };
};

const mapDispatchToProps = () => ({});

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(Jeep);
