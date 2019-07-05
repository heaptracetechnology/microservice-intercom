# _Intercom_ OMG Microservice

[![Open Microservice Guide](https://img.shields.io/badge/OMG-enabled-brightgreen.svg?style=for-the-badge)](https://microservice.guide)
[![Build Status](https://travis-ci.com/omg-services/intercom.svg?branch=master)](https://travis-ci.com/omg-services/intercom)
[![codecov](https://codecov.io/gh/omg-services/intercom/branch/master/graph/badge.svg)](https://codecov.io/gh/omg-services/intercom)


An OMG service for Intercom, it allows to messaging interaction with client app

## Direct usage in [Storyscript](https://storyscript.io/):

##### Create and save user
```coffee
>>> intercom createUser user_id:'userID' email:'emailAddress' phone:'phoneNumber' name:'userName' custom_attributes:'customAttributes'
{"id": "operationID","email": "userEmail","phone": "phoneNumber","user_id": "userID","anonymous": false,"name": "userName","avatar": {"type": "avatar"},"custom_attributes": {"custom": "attributes"}}
```
##### Send InApp message
```coffee
>>> intercom inappMessage from:'adminID' to:'receiverEmail' body:'messageBody'
{"message_type": "inapp","id": "ID","owner": {"ownerDetails"},"body": "messageBody"}
```
##### Send Email message
```coffee
>>> intercom emailMessage user_id:'userID' to:'receiverEmail' subject:'emailSubject' body:'messageBody'
{"message_type": "email","id": "ID","owner": {"ownerDetails"},"subject": "emailSubject","body": "messageBody"}
```
##### Send User message
```coffee
>>> intercom userMessage -a email:'senderEmail' body:'messageBody'
{"message_type": "inapp","id": "ID","owner": {"ownerDetails"},"body": "messageBody"}
```

Curious to [learn more](https://docs.storyscript.io/)?

✨🍰✨

## Usage with [OMG CLI](https://www.npmjs.com/package/omg)

##### Create and save user
```shell
$ omg run createUser -a user_id=<USER_ID> -a email=<EMAIL_ADDRESS> -a phone=<PHONE_NUMBER> -a name=<USER_NAME> -a custom_attributes=<CUSTOM_ATTRIBUTES> -e ACCESS_TOKEN=<ACCESS_TOKEN>
```
##### Example
```shell
$ omg run createUser -a user_id="001" -a email="testing@demo.com" -a phone=7896541230 -a name="User Name" -a custom_attributes='{"NewCust":"Creating new customer"}' -e ACCESS_TOKEN=<ACCESS_TOKEN>
```
##### Send InApp message
```shell
$ omg run inappMessage -a from=<ADMIN_ID> -a to=<RECIVER_EMAIL> -a body=<MESSAGE_BODY> -e ACCESS_TOKEN=<ACCESS_TOKEN>
```
##### Send Email message
```shell
$ omg run emailMessage -a user_id=<USER_ID> -a to=<RECIVER_EMAIL> -a subject=<EMAIL_SUBJECT> -a body=<MESSAGE_BODY> -e ACCESS_TOKEN=<ACCESS_TOKEN>
```
##### Send User message
```sh
$ omg run userMessage -a email=<SENDER_EMAIL> -a body=<MESSAGE_BODY> -e ACCESS_TOKEN=<ACCESS_TOKEN>
```

**Note**: The OMG CLI requires [Docker](https://docs.docker.com/install/) to be installed.

## License
[MIT License](https://github.com/omg-services/intercom/blob/master/LICENSE).
