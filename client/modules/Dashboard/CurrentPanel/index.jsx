import { connect } from 'react-redux';
import CurrentPanel from './CurrentPanel';

const mapStateToProps = state => {
  const { motions, gas } = state;

  const isMotion = _.some(motions, m => m != '0');
  const isGas = _.some(gas, g => g == '0');

  return {
    isMotion,
    isGas,
  }
};

const mapDispatchToProps = () => ({});

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(CurrentPanel);
