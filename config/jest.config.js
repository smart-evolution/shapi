const path = require('path');

module.exports = {
  verbose: true,
  snapshotSerializers: ['enzyme-to-json/serializer'],
  setupTestFrameworkScriptFile: './setupTest.js',
  roots: [path.resolve(__dirname, '..')],
  modulePaths: [path.resolve(__dirname, '..')],
};
