import { connect } from 'react-redux';
import Jeep from './Jeep';

const mapStateToProps = (state, ownProps) => {
  const { agent } = ownProps;
  const { id, name, isOnline, type } = agent;

  return {
    id,
    name,
    isOnline,
    type,
  };
};

const mapDispatchToProps = () => ({});

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(Jeep);
