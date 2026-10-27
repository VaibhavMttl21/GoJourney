package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Welcome to Json Video")
	// Encodejson()
	 DecodeJson()
}

func Encodejson() {
	lcoCourses := []course{
		{"Reactjs", 299, "lco.in", "pass", []string{"web-dev", "js"}},
		{"Mern", 199, "lco.in", "pass12", []string{"web-dev", "js", "fullstack"}},
		{"Angular", 299, "lco.in", "pass123", nil},
	}

	// package this data as json data

	// finalJson,err := json.Marshal(lcoCourses)
	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t")
	if err != nil {
		panic(err)
	}
	// fmt.Printf("%s\n", finalJson) here %s is converting byte data into string
	fmt.Println("Json is", string(finalJson))
}

func DecodeJson() {
	jsondatafromweb := []byte(` 
	    {
                "coursename": "Reactjs",
                "Price": 299,
                "website": "lco.in",
                "tags": ["web-dev","js"]
        }
	`)
	var lcoCourse course
	checkvalid := json.Valid(jsondatafromweb)

	if checkvalid{
		fmt.Println("Json is good to go")
		json.Unmarshal(jsondatafromweb,&lcoCourse)
		fmt.Printf("%#v\n",lcoCourse)
	}else{
		fmt.Println("Json was not valid")
	}

	// just wanted to add that in akey value pair 

	var myonlinedata map[string]interface{}
	json.Unmarshal(jsondatafromweb,&myonlinedata)
	fmt.Printf("%#v\n",myonlinedata)

	for k,v := range myonlinedata{
		fmt.Printf("key is %v and value is %v and Type is :%T\n",k,v,v)
	}
}
