package gomock

type IRunner interface {
	Setup() error
	Teardown() error
}

type Runner struct {
	services []*Services
	runners  []IRunner
}

func NewRunner(services []*Services) *Runner {
	return &Runner{
		services: services,
		runners:  make([]IRunner, 0),
	}
}

func (runner *Runner) Setup() error {
	var err error

	if runner.runners, err = runner.createRunners(runner.services); err != nil {
		return err
	}

	return runner.execute()
}

func (runner *Runner) Teardown() error {
	// http
	for _, run := range runner.runners {
		if err := run.Teardown(); err != nil {
			return err
		}
	}

	return nil
}

func (runner *Runner) createRunners(services []*Services) ([]IRunner, error) {
	var httpServices []HttpService
	var sqlServices []SqlService
	var redisServices []RedisService
	var nsqServices []NsqService
	runners := make([]IRunner, 0)

	// load the services for each individual mocking file
	for _, service := range services {
		httpServices = append(httpServices, service.HttpServices...)
		sqlServices = append(sqlServices, service.SqlServices...)
		redisServices = append(redisServices, service.RedisServices...)
		nsqServices = append(nsqServices, service.NsqServices...)
	}

	// create runners to do the job
	httpRunner := NewHttpRunner(httpServices)
	sqlRunner := NewSqlRunner(sqlServices, getDefaultSqlConfig())
	redisRunner := NewRedisRunner(redisServices, getDefaultRedisConfig())
	nsqRunner := NewNsqRunner(nsqServices, getDefaultNsqConfig())

	runners = append(runners, []IRunner{httpRunner, sqlRunner, redisRunner, nsqRunner}...)

	return runners, nil
}

func (runner *Runner) execute() error {
	for _, run := range runner.runners {
		if err := run.Setup(); err != nil {
			return err
		}
	}

	return nil
}
