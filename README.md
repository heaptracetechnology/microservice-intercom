# Intercom as a microservice
An OMG service for Intercom, it allows to messaging interaction with client app

[![Open Microservice Guide](https://img.shields.io/badge/OMG-enabled-brightgreen.svg?style=for-the-badge)](https://microservice.guide)
[![Build Status](https://travis-ci.com/heaptracetechnology/microservice-intercom.svg?branch=master)](https://travis-ci.com/heaptracetechnology/microservice-intercom)
[![codecov](https://codecov.io/gh/heaptracetechnology/microservice-intercom/branch/master/graph/badge.svg)](https://codecov.io/gh/heaptracetechnology/microservice-intercom)

## [OMG](hhttps://microservice.guide) CLI

### OMG

* omg validate
```
omg validate
```
* omg build
```
omg build
```
### Test Service

* Test the service by following OMG commands

### CLI

##### Create and save user
```sh
$ omg run createuser -a user_id=<USER_ID> -a email=<EMAIL_ADDRESS> -a phone=<PHONE_NUMBER> -a name=<USER_NAME> -a custom_attributes=<CUSTOM_ATTRIBUTES> -a companies=<COMPANIES_LIST> -e ACCESS_TOKEN=<ACCESS_TOKEN>
```
##### Example
```sh
$ omg run createuser -a user_id="001" -a email="testing@demo.com" -a phone=7896541230 -a name="User Name" -a custom_attributes='{"NewCust":"Creating new customer"}' -a companies='["abc"]' -e ACCESS_TOKEN=<ACCESS_TOKEN>
```
##### Send InApp message
```sh
$ omg run inapp_message -a from=<ADMIN_ID> -a to=<RECIVER_EMAIL> -a body=<MESSAGE_BODY> -e ACCESS_TOKEN=<ACCESS_TOKEN>
```
##### Send Email message
```sh
$ omg run email_message -a user_id=<USER_ID> -a to=<RECIVER_EMAIL> -a subject=<EMAIL_SUBJECT> -a body=<MESSAGE_BODY> -e ACCESS_TOKEN=<ACCESS_TOKEN>
```
##### Send User message
```sh
$ omg run user_message -a email=<SENDER_EMAIL> -a body=<MESSAGE_BODY> -e ACCESS_TOKEN=<ACCESS_TOKEN>
```

## License
### [MIT](https://choosealicense.com/licenses/mit/)

## Docker
### Build
```
docker build -t microservice-intercom .
```
### RUN
```
docker run -p 3000:3000 microservice-intercom
```
