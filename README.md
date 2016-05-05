# heartbeat

[![Join the chat at https://gitter.im/noaway/heartbeat](https://badges.gitter.im/noaway/heartbeat.svg)](https://gitter.im/noaway/heartbeat?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)
The heartbeat timer

[中文文档](README_ZH.md)

<h1>Heartbeat introduction</h1>

<p>Heartbeat is a response to the timing of multi-tasking callback based on Go</p>

## Installation
	
	go get -u github.com/noaway/heartbeat


<h1>Heartbeat simple to use</h1>

## Create task

	name, spec := "12138", 5
	ht, err := heartbeat.NewTast(name, spec)

	if err != nil {
		fmt.Println(err)
	}

	// Run a new mission
	ht.Start(func() error {
		//Call the callback every 5 seconds
		fmt.Println(name)
		return nil
	})

<p>More usage, please reference</p><a href="/heartbeat_test.go">heartbeat_test.go</a>
