package main

import (
	"fmt"
	"os"
	"sync"

	"github.com/subosito/gotenv"
)

type Env struct {
	m sync.Map
}

func (e *Env) Get(key string) string {
	value, ok := e.m.Load(key)
	if !ok {
		return ""
	}
	return value.(string)
}

func (e *Env) Set(key, value string) {
	e.m.Store(key, value)
}

func (e *Env) Delete(key string) {
	e.m.Delete(key)
}

// NewEnv makes a new environment
func NewEnv() *Env {
	return &Env{sync.Map{}}
}

type envFiles []string

func (e *envFiles) String() string {
	return fmt.Sprintf("%s", *e)
}

func (e *envFiles) Set(value string) error {
	*e = append(*e, value)
	return nil
}

func loadEnvs(files []string) (*Env, error) {
	if len(files) == 0 {
		files = []string{".env"}
	}

	env := NewEnv()

	// don't need to lock either environment
	for _, file := range files {
		tmpEnv, err := ReadEnv(file)

		if err != nil {
			return nil, err
		}

		// Merge the file I just read into the env.
		tmpEnv.m.Range(func(key, value interface{}) bool {
			env.m.Store(key, value)
			return true
		})
	}
	return env, nil
}

func ReadEnv(filename string) (*Env, error) {
	env := NewEnv()

	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return env, nil
	}

	fd, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer fd.Close()

	for key, val := range gotenv.Parse(fd) {
		env.Set(key, val)
	}
	return env, nil
}

func (e *Env) asArray() (env []string) {
	env = append(env, os.Environ()...)

	e.m.Range(func(name, val interface{}) bool {
		env = append(env, fmt.Sprintf("%s=%s", name, val))
		return true
	})

	return
}
