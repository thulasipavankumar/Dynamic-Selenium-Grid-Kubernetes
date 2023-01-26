## Dynamic selenium grid  [![CircleCI](https://circleci.com/gh/thulasipavankumar/Dynamic-Selenium-Grid-Kubernetes/tree/master.svg?style=svg)](https://circleci.com/gh/thulasipavankumar/Dynamic-Selenium-Grid-Kubernetes/tree/master)
- Now able to create selenium session
- Return the created session to request

## Unit testing

To do unit testing use the below command
```bash
go test ./...
```

## Dynamic selenium Grid Tasks
- [ ] Validate session request with required parameters
- [ ] Create Pod ,Serivce and Ingress
- [x] Create session  on the pod and pass on response object
- [x] Delete session and pass on response object to selenium 
- [ ] Delete Pod,Service and Ingress after session delete call

## Paths
Ingress paths to set routes for all `session requests` to `HUB pod` and `delete` calls to `Dynamic Selenium Grid Go program`
```json
"paths": [{
		"path": "<base>/session/<sessionId>/(.+)",
		"pathType": "Prefix",
		"backend": {
			"service": {
				"name": "<Hub-Service>",
				"port": {
					"number": < Hub - Port >
				}
			}
		}
	}, {
		"path": "<base>/session/<sessionId>",
		"pathType": "Prefix",
		"backend": {
			"service": {
				"name": "<Dynamic-Grid-Service>",
				"port": {
					"number": < Dynamic - Grid - port >
				}
			}
		}
	}]
```

