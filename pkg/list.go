package admin

import (
	"fmt"
	"log"
	"os"
	// "os/exec"
	// "net/http"
	// "io/ioutil"
	"encoding/json"

    "github.com/spf13/viper"
	"github.com/r--w/pocketbase"
	"github.com/mitchellh/mapstructure"
	"github.com/olekukonko/tablewriter"

	// tea "github.com/charmbracelet/bubbletea"
)

type User struct {
	Name 		string `mapstructure:"name"`
	Email 		string `mapstructure:"email"`
	Mobile 		string `mapstructure:"mobile"`
	Status   	string `mapstructure:"status"`
	ID          string  `mapstructure:"id"`
}

type PocketbaseConfig struct {
	Admin string
	Passw string
	Addre string
	Colle string
	Filte string
}
var	pbConfig = PocketbaseConfig{}

func CollectionRecords( collectionName string, searchFilter string ) []User {

	// pocketbase
	// pbConfig := PocketbaseConfig{}
	pbConfig.Admin = viper.GetString("pocketbase.admin")
	pbConfig.Passw = viper.GetString("pocketbase.password")
	pbConfig.Addre = viper.GetString("pocketbase.address")
	pbConfig.Colle = collectionName
	pbConfig.Filte = searchFilter

	users := []User{}
	users = pocketbaseList ( pbConfig )

	return users
}

func UpdateCollection( people []User ){
	client := pocketbase.NewClient(pbConfig.Addre, 
	pocketbase.WithAdminEmailPassword( pbConfig.Admin, pbConfig.Passw ))
	// r, err := client.Create("caz", map[string]any{
	// 	"name": "stocaz",
	// })
	for _, p := range people {
		var err = client.Update(pbConfig.Colle, p.ID,map[string]any{
			"status": p.Status,
		})
		if err != nil {
			log.Fatal(err)
		}
	}
}

func BackupCollection( collectionName string, searchFilter string ) {

	people := CollectionRecords( collectionName, searchFilter )
	data, err := json.Marshal(people)
    if err != nil {
        fmt.Println("Error marshaling to JSON:", err)
        return
    }
    fmt.Println(string(data))

}

func ListRecordsFromCollection( collectionName string, searchFilter string ) {

	// pocketbase
	pbConfig.Admin = viper.GetString("pocketbase.admin")
	pbConfig.Passw = viper.GetString("pocketbase.password")
	pbConfig.Addre = viper.GetString("pocketbase.address")
	// pbConfig.Colle = viper.GetString("pocketbase.collection")
	// pbConfig.Filte = viper.GetString("pocketbase.filter")
	pbConfig.Colle = collectionName
	// pbConfig.Filte = "event.name ~ '#4'"
	pbConfig.Filte = searchFilter

	users := []User{}
	users = pocketbaseList ( pbConfig )

	// log.Println(users)

	// Create a new table writer
	table := tablewriter.NewWriter(os.Stdout)
	table.SetBorder(false)                                // Set Border to false

	// Set the table headers
	table.SetHeader([]string{"#", "Name", "Email", "Mobile"})

	// Add the data to the table
	for idx, person := range users{
		// table.Append([]string{person.Name, fmt.Sprintf("%d", person.Age)})
		table.Append([]string{fmt.Sprintf("%d", idx+1), person.Name, person.Email, person.Mobile})
	}

	// // Clear the terminal screen
	// // cmd := exec.Command("reset")
	// cmd := exec.Command("clear")
	// cmd.Stdout = os.Stdout
	// cmd.Run()

	// Render the table
	table.Render()

}

func pocketbaseList( pbConfig PocketbaseConfig ) []User {

	client := pocketbase.NewClient(pbConfig.Addre, 
	pocketbase.WithAdminEmailPassword( pbConfig.Admin, pbConfig.Passw ))

	response, err := client.List(pbConfig.Colle, pocketbase.ParamsList{
		Page: 1, Size: 10000, Sort: "-created", Filters: pbConfig.Filte,
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Total of Emails: %d\n",response.TotalItems)

	users := []User{}
	err = mapstructure.Decode(response.Items, &users)
	if err != nil {
		log.Fatal(err)
	}

	return users
}
////  
////  
////  
////  func UserList() []User{
////  
////  	var table="newsletter"
////  	var secret="eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJzdXBhYmFzZSIsInJlZiI6InFpamdzZ3djZW9lbGtreWNxbnp0Iiwicm9sZSI6InNlcnZpY2Vfcm9sZSIsImlhdCI6MTY0Njc2NzQwNiwiZXhwIjoxOTYyMzQzNDA2fQ.NIQuU-LC8YbBN0MPb2QbbiwsOPNMyjAvx_VLtd4_ElQ"
////  	var url="https://qijgsgwceoelkkycqnzt.supabase.co/rest/v1/"
////  	var sele="verified=is.true&subscribed=is.true"
////  
////  	users := []User{}
////  
////  	// var newUser User 
////  
////  	// if test {
////  	//
////  	// 	log.Println("Sending a test")
////  	// 	if batchTest {
////  	// 		for ii := 1; ii <= 10; ii++ {
////  	// 			newUser.Addr = fmt.Sprintf("ruvido+%d@gmail.com", ii)
////  	// 			users = append(users,newUser)
////  	// 		}
////  	// 	} else {
////  	// 		newUser.Addr = "ruvido@gmail.com"
////  	// 		users = append(users,newUser)
////  	// 	}
////  	//
////  	// } else {
////  
////  	log.Println("fetch email addresses")
////  
////  
////  	params := table+"?"+sele 
////  	path   := url + params
////  
////  	req, err := http.NewRequest(http.MethodGet, path, nil)
////  	if err != nil {
////  		log.Printf("client: could not create request: %s\n", err)
////  		os.Exit(1)
////  	}
////  
////  	bearer := "Bearer "+secret
////  	req.Header.Add("apikey", secret)
////  	req.Header.Add("Authorization", bearer)
////  	res, err := http.DefaultClient.Do(req)
////  	if err != nil {
////  		log.Printf("client: error http request: %s\n", err)
////  		os.Exit(1)
////  	}
////  
////  	log.Printf("client: got response!\n")
////  	log.Printf("client: status code: %d\n", res.StatusCode)
////  
////  	resBody, err := ioutil.ReadAll(res.Body)
////  	if err != nil {
////  		log.Printf("client: error response body: %s\n", err)
////  		os.Exit(1)
////  	}
////  	// fmt.Printf("client: response body: %s\n", resBody)
////  	log.Printf("client: response body: OK")
////  
////  	err = json.Unmarshal(resBody, &users)
////  	if err != nil {
////  		log.Printf("json: Unmarshal error: %s\n", err)
////  		os.Exit(1)
////  	}
////  
////  	// }
////  
////  	return users
////  
////  	// if purge {
////  	// 	log.Printf("purge!!!")
////  	// 	cleanList := []User{}
////  	// 	purgeList := []User{}
////  	// 	content, err := ioutil.ReadFile("./purge.json")
////  	// 	if err != nil {
////  	// 		log.Fatal("Error when opening file: ", err)
////  	// 	}
////  	// 	err = json.Unmarshal(content, &purgeList)
////  	// 	if err != nil {
////  	// 		log.Printf("json: Unmarshal error: %s\n", err)
////  	// 		os.Exit(1)
////  	// 	}
////  	// 	for _,uu := range users {
////  	// 		for _,vv := range purgeList {
////  	// 			if uu.Addr == vv.Addr {
////  	// 				users = append(users,uu)
////  	// 			}
////  	// 		}
////  	// 	}
////  	// 	log.Printf("wellDone!")
////  	// 	return cleanList
////  	// } else {
////  	// 	return users
////  	// }
////  }
