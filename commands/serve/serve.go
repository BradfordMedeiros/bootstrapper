
// This serves provides an http server implementation for a valid bootstrapper server endpoint
// Servers require the following routes:
/*

/get/path/to/the/topic?tag=sometaghere
-X = GET

/set/path/to/the/topic?tag=sometaghere 
-X = POST
post-body: data to post 

*/

package serve 

import "fmt"

func Start(){
	fmt.Println("start placeholders")
}