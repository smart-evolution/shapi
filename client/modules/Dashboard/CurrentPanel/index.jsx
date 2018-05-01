import { connect } from 'react-redux';
import CurrentPanel from './CurrentPanel';

const mapStateToProps = state => {
  const { motions, gas } = state;

  const isMotion = _.some(_.filter(motions, m => !isNaN(Number(m))), m => m != '0');
  const isGas = _.some(_.filter(gas, g => !isNaN(Number(g))), g => g == '0');

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
