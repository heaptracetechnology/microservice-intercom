# _Intercom_ Open Microservice

> This is an intercom service

[![Open Microservice Specification Version](https://img.shields.io/badge/Open%20Microservice-1.0-477bf3.svg)](https://openmicroservices.org)
[![Open Microservices Spectrum Chat](https://withspectrum.github.io/badge/badge.svg)](https://spectrum.chat/open-microservices)
[![Open Microservices Code of Conduct](https://img.shields.io/badge/Contributor%20Covenant-v1.4%20adopted-ff69b4.svg)](https://github.com/oms-services/.github/blob/master/CODE_OF_CONDUCT.md)
[![Open Microservices Commitzen](https://img.shields.io/badge/commitizen-friendly-brightgreen.svg)](http://commitizen.github.io/cz-cli/)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg)](http://makeapullrequest.com)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)

## Introduction

This project is an example implementation of the [Open Microservice Specification](https://openmicroservices.org), a standard
originally created at [Storyscript](https://storyscript.io) for building highly-portable "microservices" that expose the
events, actions, and APIs inside containerized software.

## Getting Started

The `oms` command-line interface allows you to interact with Open Microservices. If you're interested in creating an Open
Microservice the CLI also helps validate, test, and debug your `oms.yml` implementation!

See the [oms-cli](https://github.com/microservices/oms) project to learn more!

### Installation

```
npm install -g @microservices/oms
```

## Usage

### Open Microservices CLI Usage

Once you have the [oms-cli](https://github.com/microservices/oms) installed, you can run any of the following commands from
within this project's root directory:

#### Actions

##### createUser

> Create new user

##### Action Arguments

| Argument Name    | Type     | Required | Default | Description                                                                                     |
| :--------------- | :------- | :------- | :------ | :---------------------------------------------------------------------------------------------- |
| userId           | `string` | `true`   | None    | A unique string identifier for the user. It is required on creation if an email is not supplied |
| email            | `string` | `true`   | None    | The user's email address. It is required on creation if a user_id is not supplied.              |
| phone            | `string` | `false`  | None    | The user's phone number.                                                                        |
| name             | `string` | `false`  | None    | The user's full name                                                                            |
| customAttributes | `map`    | `false`  | None    | A hash of key/value pairs containing any other data about the user you want Intercom to store   |
| ACCESS_TOKEN     | `string` | `true`   | None    | Access token for intercom                                                                       |

```shell
oms run createUser \
    -a userId='*****' \
    -a email='*****' \
    -a phone='*****' \
    -a name='*****' \
    -a customAttributes='*****' \
    -e ACCESS_TOKEN=$ACCESS_TOKEN
```

##### emailMessage

> Send message by email

##### Action Arguments

| Argument Name | Type     | Required | Default | Description                                   |
| :------------ | :------- | :------- | :------ | :-------------------------------------------- |
| userId        | `string` | `true`   | None    | The user id that is created                   |
| to            | `string` | `true`   | None    | The email address of user to send the message |
| subject       | `string` | `true`   | None    | The subject of message                        |
| body          | `string` | `true`   | None    | The body of message                           |
| ACCESS_TOKEN  | `string` | `true`   | None    | Access token for intercom                     |

```shell
oms run emailMessage \
    -a userId='*****' \
    -a to='*****' \
    -a subject='*****' \
    -a body='*****' \
    -e ACCESS_TOKEN=$ACCESS_TOKEN
```

##### inappMessage

> Send message by inapp

##### Action Arguments

| Argument Name | Type     | Required | Default | Description                        |
| :------------ | :------- | :------- | :------ | :--------------------------------- |
| from          | `int`    | `true`   | None    | The form argument will be Admin ID |
| to            | `string` | `true`   | None    | The email address of receiver      |
| body          | `string` | `true`   | None    | The body of message                |
| ACCESS_TOKEN  | `string` | `true`   | None    | Access token for intercom          |

```shell
oms run inappMessage \
    -a from='*****' \
    -a to='*****' \
    -a body='*****' \
    -e ACCESS_TOKEN=$ACCESS_TOKEN
```

##### userMessage

> Send message by user

##### Action Arguments

| Argument Name | Type     | Required | Default | Description                     |
| :------------ | :------- | :------- | :------ | :------------------------------ |
| email         | `string` | `true`   | None    | The object of sender with Email |
| body          | `string` | `true`   | None    | The body of message             |
| ACCESS_TOKEN  | `string` | `true`   | None    | Access token for intercom       |

```shell
oms run userMessage \
    -a email='*****' \
    -a body='*****' \
    -e ACCESS_TOKEN=$ACCESS_TOKEN
```

## Contributing

All suggestions in how to improve the specification and this guide are very welcome. Feel free share your thoughts in the
Issue tracker, or even better, fork the repository to implement your own ideas and submit a pull request.

[![Edit intercom on CodeSandbox](https://codesandbox.io/static/img/play-codesandbox.svg)](https://codesandbox.io/s/github/oms-services/intercom)

This project is guided by [Contributor Covenant](https://github.com/oms-services/.github/blob/master/CODE_OF_CONDUCT.md).
Please read out full [Contribution Guidelines](https://github.com/oms-services/.github/blob/master/CONTRIBUTING.md).

## Additional Resources

- [Install the CLI](https://github.com/microservices/oms) - The OMS CLI helps developers create, test, validate, and build
  microservices.
- [Example OMS Services](https://github.com/oms-services) - Examples of OMS-compliant services written in a variety of
  languages.
- [Example Language Implementations](https://github.com/microservices) - Find tooling & language implementations in Node,
  Python, Scala, Java, Clojure.
- [Storyscript Hub](https://hub.storyscript.io) - A public registry of OMS services.
- [Community Chat](https://spectrum.chat/open-microservices) - Have ideas? Questions? Join us on Spectrum.
