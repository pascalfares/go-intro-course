package main

import (
	"fmt"
	"os"
	"strings"
	"encoding/json"
	
)

func initData() map[string]string {
	//Read it from a file
	return make(map[string]string)
}

func saveData(File *f, map[string]string) {
	for k,v := range originalMap {
		s := fmt.Sprintf("$v=%v",k,v)
		file.write(s)
	}
}
func main() {
	args := os.Args
	data := initData()
	for i, arg := range args[1:]{
		a := strings.Split(arg,"=")
		fmt.Printf("split result of the args %v is %v\n",i,a )
		if len(a) == 2 {
			data[string(a[0])]=string(a[1])
		}	else {
			if len(a) == 1 {
			    fmt.Printf("\nvalue of %v is %v\n",a[0],data[string(a[0])])
		    }
	    }	
	}
	fmt.Println(json.Marshal(data))
}

