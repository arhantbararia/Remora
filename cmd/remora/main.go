package main

import (
	"fmt"
	"log"
	"remora/pkg/server"
)

func main() {
	//this is remora entrypoint
	// port 1379
	// address 127.0.0.1
	PORT := "1379"
	HOST := "127.0.0.1"

	fmt.Printf("Starting Remora Server on %v:%v \n", HOST, PORT)
	server := server.NewRemoraServer(HOST, PORT)

	err := server.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("Remora Server Error: %v", err))
	}

	//TO DO: Write Graceful shutdown route

	log.Println("Graceful shutdown complete")

}

// // Helper to run a single parse/serialize roundtrip and print verbose output

// func runRESPTest(name, data string) bool {
// 	fmt.Printf("=== Testing %s ===\n", name)
// 	fmt.Println("Input RESP:")
// 	fmt.Print(strings.ReplaceAll(data, "\r\n", "\\r\\n\n"))

// 	// Parse
// 	r := bufio.NewReader(bytes.NewBufferString(data))
// 	val, err := resp.ParseRESP(r)
// 	if err != nil {
// 		log.Printf("Parse Error: %v\n", err)
// 		return false
// 	}
// 	fmt.Println("Parsed Value:")
// 	printValue(val, 0)

// 	// Serialize
// 	var buf bytes.Buffer
// 	w := bufio.NewWriter(&buf)
// 	err = resp.WriteRESP(w, val)
// 	if err != nil {
// 		log.Printf("Serialize Error: %v\n", err)
// 		return false
// 	}
// 	w.Flush()
// 	serialized := buf.String()
// 	fmt.Println("Serialized RESP:")
// 	fmt.Print(strings.ReplaceAll(serialized, "\r\n", "\\r\\n\n"))

// 	// Parse again to check roundtrip
// 	r2 := bufio.NewReader(bytes.NewBufferString(serialized))
// 	val2, err := resp.ParseRESP(r2)
// 	if err != nil {
// 		log.Printf("Parse (Roundtrip) Error: %v\n", err)
// 		return false
// 	}
// 	fmt.Println("Parsed Value (Roundtrip):")
// 	printValue(val2, 0)
// 	fmt.Println()

// 	// Compare values (optional: implement deep equality if needed)
// 	return true
// }

// // printValue recursively prints the RESP value with indentation for nested structures.
// func printValue(val interface{}, indent int) {
// 	prefix := strings.Repeat("  ", indent)
// 	switch v := val.(type) {
// 	case []interface{}:
// 		fmt.Printf("%sArray[%d]:\n", prefix, len(v))
// 		for _, elem := range v {
// 			printValue(elem, indent+1)
// 		}
// 	case map[string]interface{}:
// 		fmt.Printf("%sMap[%d]:\n", prefix, len(v))
// 		for k, elem := range v {
// 			fmt.Printf("%s  Key: %q\n", prefix, k)
// 			printValue(elem, indent+2)
// 		}
// 	case string:
// 		fmt.Printf("%sString: %q\n", prefix, v)
// 	case error:
// 		fmt.Printf("%sError: %v\n", prefix, v)
// 	case nil:
// 		fmt.Printf("%s(nil)\n", prefix)
// 	default:
// 		fmt.Printf("%s%T: %v\n", prefix, v, v)
// 	}
// }

// func main() {
// 	// Test cases for each RESP type and edge cases
// 	tests := []struct {
// 		name string
// 		data string
// 	}{
// 		{
// 			name: "Simple String",
// 			data: "+OK\r\n",
// 		},
// 		{
// 			name: "Simple String (empty)",
// 			data: "+\r\n",
// 		},
// 		{
// 			name: "Error",
// 			data: "-ERR something went wrong\r\n",
// 		},
// 		{
// 			name: "Error (empty)",
// 			data: "-\r\n",
// 		},
// 		{
// 			name: "Integer (zero)",
// 			data: ":0\r\n",
// 		},
// 		{
// 			name: "Integer (positive)",
// 			data: ":12345\r\n",
// 		},
// 		{
// 			name: "Integer (negative)",
// 			data: ":-42\r\n",
// 		},
// 		{
// 			name: "Bulk String",
// 			data: "$6\r\nfoobar\r\n",
// 		},
// 		{
// 			name: "Bulk String (empty)",
// 			data: "$0\r\n\r\n",
// 		},
// 		{
// 			name: "Bulk String (null)",
// 			data: "$-1\r\n",
// 		},
// 		{
// 			name: "Array of Simple Strings",
// 			data: "*2\r\n+foo\r\n+bar\r\n",
// 		},
// 		{
// 			name: "Array of Integers",
// 			data: "*3\r\n:1\r\n:2\r\n:3\r\n",
// 		},
// 		{
// 			name: "Array of Bulk Strings",
// 			data: "*2\r\n$3\r\nfoo\r\n$3\r\nbar\r\n",
// 		},
// 		{
// 			name: "Array (empty)",
// 			data: "*0\r\n",
// 		},
// 		{
// 			name: "Array (null)",
// 			data: "*-1\r\n",
// 		},
// 		{
// 			name: "Nested Array",
// 			data: "*2\r\n*3\r\n:1\r\n:2\r\n:3\r\n*2\r\n+Foo\r\n-Bar\r\n",
// 		},
// 		{
// 			name: "Mixed Array",
// 			data: "*4\r\n+foo\r\n:42\r\n$3\r\nbar\r\n$-1\r\n",
// 		},
// 	}

// 	results := make([]struct {
// 		name string
// 		pass bool
// 	}, len(tests))

// 	for i, test := range tests {
// 		pass := runRESPTest(test.name, test.data)
// 		results[i].name = test.name
// 		results[i].pass = pass
// 	}

// 	fmt.Println("=== Test Summary ===")
// 	for _, r := range results {
// 		status := "PASS"
// 		if !r.pass {
// 			status = "FAIL"
// 		}
// 		fmt.Printf("%-30s : %s\n", r.name, status)
// 	}
// }
