import { connect } from 'react-redux';
import GasPanel from './GasPanel';

const mapStateToProps = state => ({
  gas: state.gas,
});

const mapDispatchToProps = () => ({});

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(GasPanel);
