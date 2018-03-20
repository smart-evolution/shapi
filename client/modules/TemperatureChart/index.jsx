import { connect } from 'react-redux';
import TravelMap from './TravelMap';

const mapStateToProps = state => ({
  countries: state.countries,
});

const mapDispatchToProps = () => ({});

export default connect(
    mapStateToProps,
    mapDispatchToProps
)(TravelMap);
