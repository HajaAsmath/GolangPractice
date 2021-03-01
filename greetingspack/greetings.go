package greetings

import "fmt"
import "errors"
import "math/rand"
import "time"

func Greetings(name string) (string,error) {
	fmt.Println(rand.Int())
	if name == "" {
		return "",errors.New("Empty name")
	}
	message := fmt.Sprintf(randomGreet(),name)
	return message,nil
}

func init() {
		rand.Seed(time.Now().UnixNano())
}

func randomGreet() string {
	formats := [] string {"Hello %v. Welcome!","Great to see you, %v","Have a nice day, %v"}
	return formats[rand.Intn(len(formats))]
}	