package main


import(
	application "myproject/application"
	authentication "myproject/authentication"
	vault "myproject/vault"
	"fmt"
)

var userName string = "bhupathi.raju@zemosolabs.com"
var password string = "Phone@7002184491"


func main(){
	accessToken := authentication.GetAccessToken(userName, password)
	fmt.Println("access", accessToken)
	application.SearchApplications(accessToken)
	vault.SearchVaults(accessToken)
	vault.CreateVault(accessToken)
}