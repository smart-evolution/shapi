import { connect } from 'react-redux';
import AgentsList from './AgentsList';

const mapStateToProps = state => ({
  agents: state.list,
});

const mapDispatchToProps = () => ({});

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(AgentsList);
