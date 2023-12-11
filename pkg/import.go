package admin

import (
	"fmt"
	"log"
	"os"
	"encoding/csv"
	"strconv"
	"time"
	// "encoding/json"
	//
    "github.com/spf13/viper"
	"github.com/r--w/pocketbase"
	// "github.com/mitchellh/mapstructure"
	// "github.com/olekukonko/tablewriter"

)

// type PocketbaseConfig struct {
// 	Admin string
// 	Passw string
// 	Addre string
// 	Colle string
// 	Filte string
// }
// var	pbConfig = PocketbaseConfig{}

func ImportData(collectionName string, dataFilename string ) {

	// pocketbase login
	pbConfig.Admin = viper.GetString("pocketbase.admin")
	pbConfig.Passw = viper.GetString("pocketbase.password")
	pbConfig.Addre = viper.GetString("pocketbase.address")
	pbConfig.Colle = collectionName
	client := pocketbase.NewClient(pbConfig.Addre, 
	pocketbase.WithAdminEmailPassword( pbConfig.Admin, pbConfig.Passw ))

	log.Println(collectionName)
	log.Println(dataFilename)
	log.Println(client)

	csvFile, err := os.Open(dataFilename)
	if err != nil {
		log.Fatal(err)
	}
	defer csvFile.Close()

	// Read the CSV file
	csvReader := csv.NewReader(csvFile)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(records)

	// const layout = "2006-01-02 15:04:05"
	layout := "2006-01-02 15:04:05"
	// timestr := "2022-11-25 18:02:08.052584+00"
	timestr := "2022-11-25 18:02:08"

	for _,p := range records[1:] {
		cr,_ := time.Parse(layout,timestr)
		em   := p[2]
		aa,_ := strconv.ParseBool(p[3])
		bb,_ := strconv.ParseBool(p[4])
		if aa && bb {
			fmt.Println(timestr)
			fmt.Println(cr, em)
		}

		// response, err := client.Create(collectionName, map[string]any{
		// 	"email": p[2],
		// })
		// if err != nil {
		// 	log.Fatal(err)
		// }
	}


	// for _, p := range people {
	// 	var err = client.Update(pbConfig.Colle, p.ID,map[string]any{
	// 		"status": p.Status,
	// 	})
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// }

}

// func ListRecordsFromCollection( collectionName string, searchFilter string ) {
//
// 	// pocketbase
// 	pbConfig.Admin = viper.GetString("pocketbase.admin")
// 	pbConfig.Passw = viper.GetString("pocketbase.password")
// 	pbConfig.Addre = viper.GetString("pocketbase.address")
// 	// pbConfig.Colle = viper.GetString("pocketbase.collection")
// 	// pbConfig.Filte = viper.GetString("pocketbase.filter")
// 	pbConfig.Colle = collectionName
// 	// pbConfig.Filte = "event.name ~ '#4'"
// 	pbConfig.Filte = searchFilter
//
// 	users := []User{}
// 	users = pocketbaseList ( pbConfig )
//
// 	// log.Println(users)
//
// 	// Create a new table writer
// 	table := tablewriter.NewWriter(os.Stdout)
// 	table.SetBorder(false)                                // Set Border to false
//
// 	// Set the table headers
// 	table.SetHeader([]string{"#", "Name", "Email", "Mobile"})
//
// 	// Add the data to the table
// 	for idx, person := range users{
// 		// table.Append([]string{person.Name, fmt.Sprintf("%d", person.Age)})
// 		table.Append([]string{fmt.Sprintf("%d", idx+1), person.Name, person.Email, person.Mobile})
// 	}
//
// 	// // Clear the terminal screen
// 	// // cmd := exec.Command("reset")
// 	// cmd := exec.Command("clear")
// 	// cmd.Stdout = os.Stdout
// 	// cmd.Run()
//
// 	// Render the table
// 	table.Render()
//
// }

// type User struct {
// 	Name 		string `mapstructure:"name"`
// 	Email 		string `mapstructure:"email"`
// 	Mobile 		string `mapstructure:"mobile"`
// 	Extra       map[string]interface{} `mapstructure:"extra"`
// 	Comment 	string `mapstructure:"comment"`
// 	Event 		string `mapstructure:"event"`
// 	Created 	string `mapstructure:"created"`
//
// 	Status   	string `mapstructure:"status"`
// 	ID          string  `mapstructure:"id"`
// }

// func UpdateCollection( people []User ){
// 	client := pocketbase.NewClient(pbConfig.Addre, 
// 	pocketbase.WithAdminEmailPassword( pbConfig.Admin, pbConfig.Passw ))
// 	for _, p := range people {
// 		var err = client.Update(pbConfig.Colle, p.ID,map[string]any{
// 			"status": p.Status,
// 		})
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 	}
// }

// func BackupCollection( collectionName string, searchFilter string ) {
//
// 	people := CollectionRecords( collectionName, searchFilter )
// 	data, err := json.Marshal(people)
//     if err != nil {
//         fmt.Println("Error marshaling to JSON:", err)
//         return
//     }
//     fmt.Println(string(data))
//
// }
// func pocketbaseList( pbConfig PocketbaseConfig ) []User {
//
// 	client := pocketbase.NewClient(pbConfig.Addre, 
// 	pocketbase.WithAdminEmailPassword( pbConfig.Admin, pbConfig.Passw ))
//
// 	response, err := client.List(pbConfig.Colle, pocketbase.ParamsList{
// 		Page: 1, Size: 10000, Sort: "-created", Filters: pbConfig.Filte,
// 	})
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	log.Printf("Total of Emails: %d\n",response.TotalItems)
//
// 	users := []User{}
// 	err = mapstructure.Decode(response.Items, &users)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
//
// 	return users
// }
