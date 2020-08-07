package vault

import(
	networkcall "myproject/networkcall"
)

const requestUrl = "https://api.skyflow.dev/v1/vaults"

type Vault struct{
	name string
	description string
	interfaces []string
	queries []string
	vaultType string
}

var vault_request_body = `
	{
		"vault": {
		"name": "EmployeeVault",
		"description": "Stores information about Employees",
		"interfaces": [
		"SQL",
		"REST"
		],
		"type": "ANALYTICAL"
		},
		"notebooks": []
		}`

func SearchVaults(accessToken string) {
	networkcall.MakeGetRequest(requestUrl, "Bearer " + accessToken)
}

func CreateVault(accessToken string){
	networkcall.MakePostRequest(requestUrl, "Bearer " + accessToken, []byte(vault_request_body))
}