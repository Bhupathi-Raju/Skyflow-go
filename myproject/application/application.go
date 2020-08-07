package application 


import(
	networkcall "myproject/networkcall"
)

type apiKey struct{
	name string
	description string
	status string
	basicaudit basicAudit
}

type basicAudit struct{

}

var requestUrl string = "https://api.skyflow.dev/v1/applications"


func SearchApplications(accessToken string)  {
	networkcall.MakeGetRequest(requestUrl, "Bearer " + accessToken)
}

// func CreateApplication(){
// 	networkcall.
// }