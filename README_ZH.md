# heartbeat
The heartbeat timer

<h1>Heartbeat简介</h1>

<p>Heartbeat是一个基于Go的多任务的定时响应回调</p>

## 安装
	
	go get -u github.com/noaway/heartbeat

<h1>Heartbeat的简单使用</h1>

<h6>创建任务</h6>

<pre>
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
</pre>

<p>更多使用方法，请参考</p><a href="/heartbeat_test.go">heartbeat_test.go</a>
