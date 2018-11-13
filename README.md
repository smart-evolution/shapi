[![Build Status](https://travis-ci.org/smart-evolution/smarthome.svg?branch=master)](https://travis-ci.org/smart-evolution/smarthome)

# Smarthome
[[Documentation]]()
[[Tasks board]](https://trello.com/b/QtZlwkhQ/project-smart-home)

### Code architecture overview:

#### Generic
- `client` - place for all the code related with client-side (css / js / elm etc)
- `docs` - place for all kind of documentation changelogs
- `hardware` - list of all available agents that webserver will connect to
- `scripts` - shell scripts required for build process 
- `views` - place for all server-side rendering templates

#### CSS
- `components` - can be extended in any place as they are generic
- `modules` - place for all business logic related modules
-- modules classes should be extended only from one level above in sub-modules
-- if some classes are not used directly in parent component but child component extends them - they should be virtual
-- all sub-modules imports which are more details should be placed in the bottom of the file

#### Go
- `controllers` - handles side-effects related with web rendering
- `services` - place to handle all side effects like database persistence, sending emails etc.
- `models` - represents entities from the business layer
-- models should not contain or handle any side effects, they should perform any side-effect with DI 
- `utils` - place for helpers

#### JavaScript
- `components` - keeps generic components, they should be business-logic agnostic
- `modules` - place for all business logic related modules
