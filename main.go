package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
	"github.com/google/uuid"
)

/*
stuct
*/
type foobar struct {
    uidFoo string
    uidBar string
}

type stock struct {
    mu sync.Mutex
    foo []string
    bar []string
    foobar []foobar
    cash int
    numberRobots int
}

/*
start
*/
func main() {
    fmt.Println("Hello foobactory!")
    start := time.Now()
    s := stock{}
    s. orchestrator ()
    end := time.Now()
    duration := end.Sub(start)
    fmt.Println("stop")
    fmt.Println("time: ", duration)
}

/*
boss
*/
func (s *stock)  orchestrator () {
    actions := [30]string{}
    s.numberRobots = 2
    for true {
        s.mu.Lock()
        nbRobot := s.numberRobots
        // Finished
        if nbRobot >= 30 {
            fmt.Println("well done")
            break
        }
        s.mu.Unlock()
        for i:=0; i<nbRobot; i++ {
            change := false
            action := s.getAction()
            oldAction := actions[i]
            actions[i] = action
            if oldAction != action && oldAction != "nil"{
                change = true
            }
            // goroutine
            go s.robot(action, change)
        }
    }
}

// Goroutine
 func (s *stock) robot(action string, change bool) {
    if change {
        time.Sleep(5 * time.Second)
    }
    s.execut(action)
 }

// getAction: Choose the right action for the stock
func (s *stock) getAction() string {
    s.mu.Lock()
    var action string 

    if s.cash >= 3 && len(s.foo) >= 6 {
        action = "buyRobot"
    } else if  len(s.foobar) > 4 {
        action = "sellFoobar"
    } else if len(s.foo) > 0 && len(s.bar) > 0 && len(s.foobar) < 5 {
        action = "buildFoobar"
    } else if len(s.foo) < 6 {
        action = "mineFoo"
    } else {
        action = "mineBar"
    }
    s.mu.Unlock()
    return action
}

// execut: Launch the functions
func (s *stock) execut(action string) {
    switch action{
    case "mineBar":
        s.mineBar()
    case "mineFoo":
        s.mineFoo()
    case "buildFoobar":
        s.buildFoobar()
    case "sellFoobar":
        s.sellFoobar()
    case "buyRobot":
        s.buyRobot()
    }
}

/*
factory
*/

// min foo during 1 second 
func (s *stock) mineFoo() {
    fmt.Println("mineFoo")
    time.Sleep(1 * time.Second)
    s.mu.Lock()
    uid := getUid()
    s.foo = append(s.foo, uid)
    s.mu.Unlock()
}

// min bar between 0,5 and 2 seconds
func (s *stock) mineBar() {
    fmt.Println("mineBar")
    
    rand.Seed(time.Now().UnixNano())
    n := rand.Intn(2000-500) + 500
    time.Sleep(time.Duration(n) * time.Millisecond)
    s.mu.Lock()
    uid := getUid()
    s.bar = append(s.bar, uid)
    s.mu.Unlock()
}

// Build foobar during 2 seconds with a probability of 60%
func (s *stock) buildFoobar() {
    fmt.Println("buildFoobar")
    time.Sleep(2 * time.Second)

    rand.Seed(time.Now().UnixNano())
    n := rand.Intn(10-1) + 1
    s.mu.Lock()
    if len(s.bar) > 0 && len(s.foo)  > 0 {
        foo := s.foo[0]
        bar := s.bar[0]
    
        if n < 6 { // success
            s.foo =  RemoveString(s.foo, 0)
            s.foobar = append(s.foobar, foobar{foo, bar})
        }
        s.bar =  RemoveString(s.bar, 0)
    }
    s.mu.Unlock()
}

// Sells foobars and earns 1 euros per foobar
func (s *stock) sellFoobar() {
    fmt.Println("sellFoobar")
    time.Sleep(10 * time.Second)
    s.mu.Lock()
    lenghtFoobars := len(s.foobar)
    if lenghtFoobars > 4 {
        lenghtFoobars = 4
    }

    for i:=0; i<lenghtFoobars; i++ {
        s.foobar = RemoveFoobar(s.foobar, 0)
        s.cash += 1
    }
    s.mu.Unlock()
}

// Buy one robot for 3 euros and 6 foo
func (s *stock) buyRobot() {
    fmt.Println("buyRobot")
    s.mu.Lock()
    
    s.numberRobots += 1
    s.cash = s.cash - 3

    if len(s.foo) >= 6 {
        for i:=0; i<5; i++ {
            s.foo =  RemoveString(s.foo, 0)
        }
    }
    s.mu.Unlock()
}

/*
utils
*/

// getUid: Get uuid
func getUid() string {
    return uuid.New().String()
}

// RemoveString: Remove string value in a slice
func  RemoveString(s []string, index int) []string {
	return append(s[:index], s[index+1:]...)
}

//  RemoveFoobar: Remove foobar value in a slice
func RemoveFoobar(s []foobar, index int) []foobar {
	return append(s[:index], s[index+1:]...)
}
