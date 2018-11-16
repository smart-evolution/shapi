import { connect } from 'react-redux';
import SoundPanel from './SoundPanel';

const mapStateToProps = () => ({
  sounds: [],
});

const mapDispatchToProps = () => ({});

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(SoundPanel);
