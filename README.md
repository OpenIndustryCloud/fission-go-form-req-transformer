[![Coverage Status](https://coveralls.io/repos/github/OpenIndustryCloud/fission-go-form-req-transformer/badge.svg?branch=master)](https://coveralls.io/github/OpenIndustryCloud/fission-go-form-req-transformer?branch=master)

# TYYPEFORM request transformation API (JSON to JSON)

`form-data-transformer.go` is an API which accepts JSON payload of TYPEFORM and transform the JSON data which can used by other APIs to query weather info or registering a ticket to ZenDesk.

It performs JSON to JSON tranformation by mapping data manually.

## API reference

- [TYPEFORM JSON Response](https://developer.typeform.com/webhooks/example-response/)
- [Zendesk API](https://developer.zendesk.com/rest_api/docs/core/tickets)

## Authentication

This API do not need any authentication data. 


## Error hanlding

Empty Payload or malformed JSON would result in error reponse.

- Technical error : `{"status":400,"message":"EOF"}`


## Sample Input/Output

- Request payload

```
{"event_id":"GQQPPJJtfd","event_type":"form_response","form_response":{"form_id":"krCLZ1","token":"4","submitted_at":"2017-11-06T10:25:13Z","hidden":{"email":"xxxxx","name":"xxxxx","phone":"xxxxx","policy":"xxxxx"},"definition":{"id":"krCLZ1","title":"TV claim","fields":[{"id":"XdVgAjnucXRP","title":"Was there a theft or a deliberate event by a 3rd party?","type":"yes_no","allow_multiple_selections":false,"allow_other_choice":false},{"id":"AZbSqcXTKlED","title":"Please enter a  crime/loss reference number.","type":"short_text","allow_multiple_selections":false,"allow_other_choice":false},{"id":"63241151","title":"Please upload a detailed image of the damage","type":"file_upload","allow_multiple_selections":false,"allow_other_choice":false},{"id":"63241244","title":"Please upload any photos of proof of purchase / ownership / box / receipt","type":"file_upload","allow_multiple_selections":false,"allow_other_choice":false},{"id":"63241572","title":"Are you aware of anything else relevant to your claim that you would like to advise us of at this stage?","type":"long_text","allow_multiple_selections":false,"allow_other_choice":false},{"id":"63241542","title":"If you have any other insurance or warranties covering your TV, please advise us of the company name.","type":"short_text","allow_multiple_selections":false,"allow_other_choice":false},{"id":"63241330","title":"When did the incident happen?","type":"date","allow_multiple_selections":false,"allow_other_choice":false},{"id":"QYabj3uz2gs9","title":"In as much detail as possible, please provide a full written account of what has happened to your TV, including where it happened.","type":"long_text","allow_multiple_selections":false,"allow_other_choice":false},{"id":"WwzqHPb0K9Wv","title":"Make of the TV","type":"short_text","allow_multiple_selections":false,"allow_other_choice":false},{"id":"u2dzBoYFjbRA","title":"Serial number of the TV","type":"short_text","allow_multiple_selections":false,"allow_other_choice":false},{"id":"63241939","title":"We have made the following assumptions about your property, you and anyone living with you","type":"legal","allow_multiple_selections":false,"allow_other_choice":false},{"id":"lGCVB9tse6Re","title":"Model number of the TV","type":"short_text","allow_multiple_selections":false,"allow_other_choice":false},{"id":"XiKbgOrX1OAp","title":"Do you still have the TV in your possession?","type":"yes_no","allow_multiple_selections":false,"allow_other_choice":false},{"id":"63391165","title":"What was the purchase price of the TV?","type":"number","allow_multiple_selections":false,"allow_other_choice":false}]},"answers":[{"type":"boolean","boolean":true,"field":{"id":"XdVgAjnucXRP","type":"yes_no"}},{"type":"text","text":"DU00000000","field":{"id":"AZbSqcXTKlED","type":"short_text"}},{"type":"file_url","file_url":"https://admin.typeform.com/form/results/file/download/krCLZ1/63241151/97750857a9d3-22228398_888332207989782_5676320169463975335_n.jpg","field":{"id":"63241151","type":"file_upload"}},{"type":"file_url","file_url":"https://admin.typeform.com/form/results/file/download/krCLZ1/63241244/e6cc84865de7-22405371_888332421323094_6861338905885136899_n.jpg","field":{"id":"63241244","type":"file_upload"}},{"type":"text","text":"None","field":{"id":"63241572","type":"long_text"}},{"type":"text","text":"None","field":{"id":"63241542","type":"short_text"}},{"type":"date","date":"2017-10-01","field":{"id":"63241330","type":"date"}},{"type":"text","text":"what happened","field":{"id":"QYabj3uz2gs9","type":"long_text"}},{"type":"text","text":"LG","field":{"id":"WwzqHPb0K9Wv","type":"short_text"}},{"type":"text","text":"SerialNo","field":{"id":"u2dzBoYFjbRA","type":"short_text"}},{"type":"boolean","boolean":true,"field":{"id":"63241939","type":"legal"}},{"type":"text","text":"ModelNo","field":{"id":"lGCVB9tse6Re","type":"short_text"}},{"type":"boolean","boolean":true,"field":{"id":"XiKbgOrX1OAp","type":"yes_no"}},{"type":"number","number":20000,"field":{"id":"63391165","type":"number"}}]}}
```

- Response
```
{"status":200,"ticket_details":{"ticket":{"type":"incident","subject":"Storm surge claim","priority":"normal","status":"new","comment":{"html_body":"\u003cp\u003e\u003cb\u003eAre you still have possession of the damage items (i.e. damaged guttering)?\u003c/b\u003e : \u003c/p\u003e\u003chr\u003e\u003cp\u003e\u003cb\u003eIf there has been any recent maintenance carried out on your home, please describe it\u003c/b\u003e : no\u003c/p\u003e\u003chr\u003e\u003cp\u003e\u003cb\u003eIf you have any other insurance or warranties covering your home, please advise us of the company name.\u003c/b\u003e : no\u003c/p\u003e\u003chr\u003e\u003cp\u003e\u003cb\u003eWe have made the following assumptions about your property, you and anyone living with you\u003c/b\u003e : \u003c/p\u003e\u003chr\u003e\u003cp\u003e\u003cb\u003eWhen did the incident happen?\u003c/b\u003e : 2017-11-02\u003c/p\u003e\u003chr\u003e\u003cp\u003e\u003cb\u003eIn as much detail as possible, please use the text box below to describe the full extent of the damage to your home and how you discovered it.\u003c/b\u003e : A cow was blown on top of my roof and destroyed tiles.\u003c/p\u003e\u003chr\u003e\u003cp\u003e\u003cb\u003ePlease describe the details of the condition of your home prior to discovering the damage\u003c/b\u003e : perfect\u003c/p\u003e\u003chr\u003e\u003cp\u003e\u003cb\u003eAre you aware of anything else relevant to your claim that you would like to advise us of at this stage?\u003c/b\u003e : nothing\u003c/p\u003e\u003chr\u003e\u003cp\u003e\u003cb\u003eWhere did the incident happen? (City/town name)\u003c/b\u003e : london\u003c/p\u003e\u003chr\u003e\u003cp\u003e\u003cb\u003eWould you like to upload more images?\u003c/b\u003e : \u003c/p\u003e\u003chr\u003e"},"requester":{"locale_id":1,"name":"Maarten","email":"mectors@gmail.com","phone":"12345678","policy_number":"12345678"},"ticket_form_id":0,"event_id":"SChin7eteE","token":"6","submitted_at":"2017-11-06T14:13:27Z"}},"weather_api_input":{"city":"london","date":"20171102"},"tv_claim_data":{"tv_reciept_image_url":""},"storm_claim_data":{"incident_place":"london","incident_date":"2017-11-02","damage_image_url_2":"https://admin.typeform.com/form/results/file/download/H8mm3s/j79cNctIvogK/aaee97a02b39-1.cow_on_a_roof.jpg"}}
```


## Example Usage

## 1.  Deploy as Fission Functions

First, set up your fission deployment with the go environment.

```
fission env create --name go-env --image fission/go-env:1.8.1
```

To ensure that you build functions using the same version as the
runtime, fission provides a docker image and helper script for
building functions.



- Download the build helper script

```
$ curl https://raw.githubusercontent.com/fission/fission/master/environments/go/builder/go-function-build > go-function-build
$ chmod +x go-function-build
```

- Build the function as a plugin. Outputs result to 'function.so'

`$ go-function-build form-req-transformer.go`

- Upload the function to fission

`$ fission function create --name form-req-transformer --env go-env --package function.so`

- Map /form-req-transformer to the form-req-transformer function

`$ fission route create --method POST --url /form-req-transformer --function form-req-transformer`

- Run the function

```$ curl -d `{"ticket": {"subject": "My printer is on fire!", "comment": {"body": "The smoke is very colorful."}}}` -H "Content-Type: application/json" -X POST http://$FISSION_ROUTER/form-req-transformer```

## 2. Deploy as AWS Lambda

> to be updated