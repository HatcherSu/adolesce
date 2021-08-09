package app

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

type RunFunc func(name string) error

type App struct {
	name      string
	shortDesc string
	LongDesc  string
	runFunc   RunFunc
	cmd       *cobra.Command
	sigs      []os.Signal
	stopFunc  []func()
}

func NewApp(name string, opts ...Option) *App {
	a := &App{
		name: name,
		sigs: []os.Signal{os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT},
	}

	for _, o := range opts {
		o(a)
	}
	a.initCommand()

	return a
}

func (a *App) initCommand() {
	cmd := &cobra.Command{
		Use:   a.name,
		Short: a.shortDesc, // short description
		Long:  a.LongDesc,  // long description
	}
	if a.runFunc != nil {
		cmd.RunE = a.runCommand
	}
	a.cmd = cmd
}

func (a *App) runCommand(cmd *cobra.Command, args []string) error {
	// todo 配置命令输入参数等
	if a.runFunc != nil {
		return a.runFunc(a.name)
	}
	return nil
}

func (a *App) Run() {
	// 开始监听关闭
	a.ListenAndClose()

	if err := a.cmd.Execute(); err != nil {
		fmt.Printf("%v %v\n", color.RedString("Error:"), err)
		os.Exit(1)
	}
}

// RegisterStopFunc 注册关闭函数
func (a *App) RegisterStopFunc(fs ...func()) {
	a.stopFunc = fs
}

// ListenAndClose 监听和关闭
func (a *App) ListenAndClose() {
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, a.sigs...)

		<-c

		var wg sync.WaitGroup
		for _, stop := range a.stopFunc {
			wg.Add(1)
			go func(f func()) {
				defer wg.Done()
				f()
			}(stop)
		}
		wg.Wait()
	}()
}
