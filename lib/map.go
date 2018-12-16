package lib

import "fmt"

func InputMap(size int) map[string]string {
	m := make(map[string]string);

	for i:=0; i<size; i++ {
		var key,value string;

		fmt.Scan(&key,&value);
		m[key] = value;
	}

	return m;
}

func PrintMap(m map[string]string)  {
	for key,value :=range m {
		fmt.Printf("%s => %s\n",key,value);
	}
}
