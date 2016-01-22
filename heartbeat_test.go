package heartbeat

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func Test_NewTast(t *testing.T) {
	//Create a new task
	name, spec := 12138, 5
	ht, err := NewTast(strconv.Itoa(name), spec)

	if err != nil {
		fmt.Println(err)
	}

	// Run a new mission
	ht.Start(func() error {
		fmt.Println(name)
		return nil
	})

	time.Sleep(time.Second * 1000)
}

func Test_GetActivity(t *testing.T) {
	t.Log(GetActivity())
}

func Test_ClearTast(t *testing.T) {
	name := "12138"
	if err := ClearTast(name); err != nil {
		t.Error(err)
	}
	t.Log(GetActivity())
}

func Test_PauseTast(t *testing.T) {
	name := "12138"
	if err := PauseTast(name); err != nil {
		t.Error(err)
	}
	t.Log(GetActivity())
}

func Test_RunTast(t *testing.T) {
	name, spec := "12138", 5
	if err := RunTast(name, spec); err != nil {
		t.Error(err)
	}
	t.Log(GetActivity())
}
