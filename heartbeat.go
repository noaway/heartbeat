package heartbeat

import (
	"errors"
	"runtime"
	"sync"
	"time"
)

//状态
const (
	Stop    = "stop"
	Running = "running"
	Pause   = "pause"
)

var (
	maps    = make(map[string]*Task)
	isstart = true
)

func setmaps(name string, t *Task) {
	var self Task
	self.m.Lock()
	defer self.m.Unlock()
	maps[name] = t
}

func getmaps(name string) (t *Task) {
	if maps == nil {
		return nil
	}
	return maps[name]
}

func delmaps(name string) {
	var self Task
	self.m.Lock()
	defer self.m.Unlock()
	delete(maps, name)
}

//初始化一个新任务
func NewTast(name string, spec int) (*Task, error) {
	if name != "" {
		if getmaps(name) != nil {
			isstart = false
			return nil, errors.New("Task has been")
		}
		t := &Task{
			Name:       name,
			Status:     Running,
			Spec:       spec,
			CreateTime: time.Now(),
			Chan:       make(chan string),
		}
		setmaps(name, t)
		return t, nil
	} else {
		isstart = false
		return nil, errors.New("The name cannot be empty")
	}
}

type Task struct {
	Name       string
	Status     string
	Spec       int
	Chan       chan string
	m          sync.Mutex
	CreateTime time.Time
}

func run(self *Task, f func() error) {
	if !isstart {
		return
	}
	timer := time.NewTicker(time.Duration(self.Spec) * time.Second)
	for {
		select {
		case <-timer.C:
			if self.Status == Pause {
				runtime.Gosched()
				continue
			}
			if err := f(); err != nil {
				timer.Stop()
				return
			}
		case status, ok := <-getmaps(self.Name).Chan:
			if !ok {
				if g := getmaps(self.Name); g != nil {
					close(g.Chan)
				}
			}
			switch status {
			case Stop:
				timer.Stop()
				return
			case Running:
				self.Status = Running
			case Pause:
				self.Status = Pause
			}
		}
	}
}

func (self *Task) Start(f func() error) {
	go run(self, f)
}

func GetActivity() (m []interface{}) {
	for _, k := range maps {
		if k != nil {
			dict := make(map[string]interface{})
			dict["Name"] = k.Name
			dict["CreateTime"] = k.CreateTime
			dict["Status"] = k.Status
			dict["Spec"] = k.Spec
			m = append(m, dict)
		}
	}
	return m
}

func ClearTast(name string) error {
	gm := getmaps(name)
	if gm != nil {
		gm.Chan <- Stop
		close(gm.Chan)
		delmaps(name)
		return nil
	}
	return errors.New("The name of the task is invalid")
}

func PauseTast(name string) error {
	gm := getmaps(name)
	if gm != nil {
		gm.Chan <- Pause
		return nil
	}
	return errors.New("The name of the task is invalid")
}

func Restarting(name string) error {
	gm := getmaps(name)
	if gm != nil {
		gm.Chan <- Running
		return nil
	}
	return errors.New("The name of the task is invalid")
}

func RunTast(name string, spec int) error {
	gm := getmaps(name)
	if gm != nil {
		gm.Chan <- Running
		gm.Status = Running
		return nil
	}
	return errors.New("The name of the task is invalid")
}
