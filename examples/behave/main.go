package main

import (
	"fmt"
	"time"

	"github.com/woshihaomei/go-behave"
	"github.com/woshihaomei/go-behave/core"
	"github.com/woshihaomei/go-behave/store"
	"github.com/woshihaomei/go-behave/util"

	// Use dot imports to make a tree definition look nice.
	// Be careful when doing this! These packages export
	// common word identifiers such as "Fail" and "Sequence".
	. "github.com/woshihaomei/go-behave/common/action"
	. "github.com/woshihaomei/go-behave/common/composite"
	. "github.com/woshihaomei/go-behave/common/decorator"
)

// var delayingRoot = Repeater(core.Params{"n": 2},
// 	PersistentSequence(
// 		Delayer(core.Params{"ms": 700},
// 			Succeed(nil, nil),
// 		),
// 		Delayer(core.Params{"ms": 400},
// 			Inverter(nil,
// 				Fail(nil, nil),
// 			),
// 		),
// 	),
// )

var someRoot = Sequence(
	Repeater(core.Params{"n": 2}, Succeed()),
	//Delayer(core.Params{"ms":2000}, Succeed()),
	RandomDelayer(core.Params{"msMin":1000, "msMax":3000}, Succeed()),
	//Succeed(),
	RandomSelector(Work(func() {
		fmt.Println("WorkWorkWorkWork111!!!!")
	}), Work(func() {
		fmt.Println("WorkWorkWorkWork222!!!!")
	}),Work(func() {
		fmt.Println("WorkWorkWorkWork333!!!!")
	}),Work(func() {
		fmt.Println("WorkWorkWorkWork444!!!!")
	})),
)

// ID is a simple type only used as tree owner for testing.
// In a real scenario, the owner would be an actual entity
// with some interesting state and functionality.
type ID int

// String returns a string representation of ID.
func (id ID) String() string { return fmt.Sprint(int(id)) }

func main() {
	testTree(someRoot)
}

func testTree(root core.Node) {
	fmt.Println("Testing tree...")

	tree, err := behave.NewBehaviorTree(
		behave.Config{
			Owner: ID(1337),
			Data:  store.NewBlackboard(),
			Root:  root,
		},
	)
	if err != nil {
		panic(err)
	}

	for {
		status := tree.Update()
		util.PrintTreeInColor(tree.Root)
		fmt.Println(time.Now())
		if status == core.StatusSuccess {
			//break
		}
		time.Sleep(1000 * time.Millisecond)
	}

	fmt.Println("Done!")
}
