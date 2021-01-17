# Architecture
This section seeks explain a little more about clean architecture.

<br/>

<img style="margin-bottom: 10px" src="architecture.png"/>

<br/>

## Summary
This image is a summary of concepts, relationships and responsibilities of the existenting layers.
With it in mind let's talk about the following topics:

- [Purpose](#Purpose)
- [Principles](#Principles)
- [Concepts](#Concepts)
- [Layers](#Layers)
- [Relationships](#Relationships)
- [Rules](#Rules)
- [More about](#Moreabout)

<br/>


## Purpose 
How objective this architecture seek solve or improve the following points:

- structure expressivity
- maintenance 
- testability
- decoupling of business solve from Frameworks and drives

<br/>

## Principles
The principles that guide this architecture are:

- SoC (separation of concerns)
- Use cases how first class citizens
- Dependency injection
- Testability

## Concepts
There are some concepts that are important we knowing before we following analyzing in the next sections.

### Policies
Software systems are statements of politics. In essence this mean what a program of computer really is. A program of computer is a description detail of a political that coordinates the transformations of input and output.

#### High level
How much more away of inputs and outputs more high is politic's level. Generally this type of politic describe flows more abstracts. It tend change with less frequency but for reasons more substantial.\
Example: Changes in calculate of interest of enterprise.

#### Low level
The policies that cope with inputs and outputs from system are policies of more low level of the system. These policies tend to change more frequently for less substantial reasons because they are closer to I / O.\
Examples: Change of database, changes in UI. 

<br/>

### Enterprise business rules
These rules must are the system's hearth. They're part of high level policies from system and should was the part more reusable of it. They concern the rules that are responsible for generate or save the money from enterprise, they exist even don't exist an automatized system.\
Example: The collection of N% of interest by a loan is a rule of business tha generate receive for the bank and exist without a automatized system.

### Application business rules
Not all business rules are so pure, some of them generate or save money for the company by defining the way an automated system operates. These rules can't be used in a manual environment, because they only make sense as part of an automated system. These policies are low than the business rule by treat of facts not so crucial.\
Example: The control of who could or not view a interface of system.

<br/>

## Layers

### Entity
### Use case
### Adapter of interfaces
### Frameworks and drives

<br/>

## Relationship
### Layers
### Objects 
<br/>

## More about
