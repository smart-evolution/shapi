import { connect } from 'react-redux';
import SoundPanel from './SoundPanel';
const mapStateToProps = (state, ownProps) => {
  const { agent } = ownProps;

  return {
    sounds: [],
  };
};

const mapDispatchToProps = () => ({});

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(SoundPanel);
