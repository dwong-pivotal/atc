package engine

import (
	"errors"

	garden "github.com/cloudfoundry-incubator/garden/api"
	"github.com/concourse/atc/db"
	"github.com/concourse/turbine"
	"github.com/concourse/turbine/event"
	"github.com/pivotal-golang/lager"
)

var ErrBuildNotFound = errors.New("build not found")
var ErrEndOfStream = errors.New("end of stream")
var ErrReadClosedStream = errors.New("read of closed stream")
var ErrCloseClosedStream = errors.New("close of closed stream")

// intermediate
type BuildPlan turbine.Build

//go:generate counterfeiter . Engine
type Engine interface {
	Name() string

	CreateBuild(db.Build, BuildPlan) (Build, error)
	LookupBuild(db.Build) (Build, error)
}

//go:generate counterfeiter . Build
type Build interface {
	Metadata() string

	Abort() error
	Hijack(garden.ProcessSpec, garden.ProcessIO) error
	Subscribe(from uint) (EventSource, error)
	Resume(lager.Logger) error
}

//go:generate counterfeiter . EventSource
type EventSource interface {
	Next() (event.Event, error)
	Close() error
}