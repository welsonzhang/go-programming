package main

import(
"fmt"
)

func main() {

    var countryCapitalMap map[string]string
    countryCapitalMap = make(map[string]string)

    // insert key value
    countryCapitalMap["France"] = "Paris"
   countryCapitalMap["Italy"] = "Rome"
   countryCapitalMap["Japan"] = "Tokyo"
   countryCapitalMap["India"] = "New Delhi"

   // range
   for country := range countryCapitalMap {
       fmt.Println("Capital of",country,"is",countryCapitalMap[country])
   }

   // find
   capital, ok := countryCapitalMap["United States"]
   if(ok) {
    fmt.Println("Capital of United States is", capital) 
   }else {
    fmt.Println("Capital of United States is not present")
   }

   // add
   countryCapitalMap["United States"] = "WastonDC"

   // delete
   delete(countryCapitalMap,"France");
   fmt.Println("Entry for France is deleted")  
   
   // print
   for country := range countryCapitalMap {
      fmt.Println("Capital of",country,"is",countryCapitalMap[country])
   }
}
