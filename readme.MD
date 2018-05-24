# Settings Service
The Settings Service is a REST server that stores and retreives groups of JSON files that conform to schemas. 
The basic idea is that you configure a set of JSON schemas, known as a Type. Then you can add 1 or more
instances.

## Object Model

### Reference Model

*TYPE* - A Type represents the specification for a type of a set of preferences or JSON. For instance, users are a type, 
a specific application is a type. All instances of a type share the same JSON schema set for validation, etc. 
Some Types can only have a single instance while others can have many

*SPECIFICATION* - A Specification represents a single JSON schema that will be used to validate against

*CATEGORY* - A category is a logical grouping of specifications. Primarily for visualization

So the cardinality of the reference model is that we can have 1 to * TYPE. A type can have 1 to * categories
and a category can have 1 to * specifications. 

### Instance Model

*INSTANCE* - An Instance is a single instantiation of a Type. 

*SETTING* - A Setting is the JSON document that is validated by a single Specification

## Usage

There are a few primary use cases for this approach.

### Shared Preference Store

The first use case is a shared preference store. This can store user preferences for all apps into a consolidated preferences. 
Now apps can just request the preferences for their application.

### Shared Configuration Store

Application configuration can be stored this way too. In this manner the application would only allow a single instance per type

