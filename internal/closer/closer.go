package closer

import (
	"log"
	"os"
	"os/signal"
	"sync"
)

var globalCloser = New()

// Add adds `func() error` callback to the globalCloser
func Add(f ...func() error) {
	globalCloser.Add(f...)
}

// Wait ...
func Wait() {
	globalCloser.Wait()
}

// CloseAll ...
func CloseAll() {
	globalCloser.CloseAll()
}

// Closer управляет списком функций закрытия ресурсов, таких как соединения с базой данных или файлы.
// Он гарантирует, что все зарегистрированные функции будут вызваны при завершении работы программы,
// и каждая функция будет вызвана только один раз, даже если CloseAll() будет вызвано несколько раз.
type Closer struct {
	mu    sync.Mutex
	once  sync.Once
	done  chan struct{}
	funcs []func() error
}

// New возвращает новый Closer.
// Если указан срез []os.Signal, Closer автоматически вызовет
// CloseAll при получении одного из сигналов от операционной системы.
func New(sig ...os.Signal) *Closer {
	c := &Closer{done: make(chan struct{})}

	if len(sig) > 0 {
		go func() {
			ch := make(chan os.Signal, 1)
			signal.Notify(ch, sig...)
			<-ch
			signal.Stop(ch)
			c.CloseAll()
		}()
	}

	return c
}

// Add  добавляет функцию закрытия в closer
func (c *Closer) Add(f ...func() error) {
	c.mu.Lock()
	c.funcs = append(c.funcs, f...)
	c.mu.Unlock()
}

// Wait блокирует выполнение до тех пор, пока все функции закрытия не завершатся
func (c *Closer) Wait() {
	<-c.done
}

// CloseAll вызывает все зарегистрированные функции закрытия ресурсов.
// Гарантирует, что каждая функция будет вызвана только один раз благодаря sync.Once.
// Выполняет все функции закрытия асинхронно, собирая ошибки их выполнения.
// После завершения всех операций закрытия закрывается канал done,
// сигнализируя о завершении процесса закрытия.
func (c *Closer) CloseAll() {
	c.once.Do(func() {
		defer close(c.done)

		c.mu.Lock()

		funcs := c.funcs
		c.funcs = nil

		c.mu.Unlock()

		// Канал для сбора ошибок выполнения функций
		errs := make(chan error, len(funcs))

		// Асинхронный вызов всех зарегистрированных функций закрытия
		for _, f := range funcs {
			go func(f func() error) {
				errs <- f()
			}(f)
		}

		for i := 0; i < cap(errs); i++ {
			if err := <-errs; err != nil {
				log.Println("error returned from Closer")
			}
		}
	})
}
