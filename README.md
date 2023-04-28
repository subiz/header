## quick build go only
./build.sh

For any error, please consume README.md in `protobuf-cpp-3.5.1` or `protobuf-cpp-3.5.1/src`

## Build docker image subiz/protobuild
```sh
docker build  -t live360vn/protobuild:2.5 .
```


### Merge

conversation keep id
but allow to list using new user id

order

note

attribute?


### Phone driver protocal
1. service exposed to conversation
```
  1. /outgoingCall(acc_id, from_device, to_device, timeout): status
  same = n,Set(CALLERID(NUM)=842474943499)
  SIP/0364821895@itel)
  2. /registerNumber(accid, number)
  3. /registerDevice(accid, device) // dial internal to handle incomming call
  4. /unregistryNumber(accid, number)
  5. /unregistryDevice(accid, number)
  6. /getCallStatus
```

2. api
```
 1. convo.UpdatePhoneDevice
 2. convo.IncommingCall
 3. convo.TimeoutCall
 4. convo.StartConversation => start a call
 8. convo.SendMessage /events => update call status
 9. convo.UpdateConversationInfo => /conversations/x?recoding_url=asdf => update recoding url

{
  action: 'dial|hangup|wait',
  account_id,
  conversation_id,
  timeout='40s',
  device='1235,53534'
  music='https://abc.com/gretting.wav'
  recoding=true
}
```

### Lead source

Bot
Agent
Web
WebPlugin
custom-x-y-z
