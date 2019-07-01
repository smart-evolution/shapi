import { connect } from 'react-redux';
import Unknown from './Unknown';

const mapStateToProps = (state, ownProps) => {
  const { agent } = ownProps;
  const { id, name, type, isOnline } = agent;

  return {
    id,
    name,
    type,
    isOnline,
  };
};

const mapDispatchToProps = () => ({});

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(Unknown);
