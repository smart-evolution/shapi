import { connect } from 'react-redux';
import SoundPanel from './SoundPanel';
import { getTicks } from '../queries';

const mapStateToProps = (state) => {
  const sounds = getTicks(state.times, state.sounds);

  return {
    sounds,
  };
};

const mapDispatchToProps = () => ({});

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(SoundPanel);
