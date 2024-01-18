package definitions

// Container is a root dependency injection container. It is required to describe
// your services.
type Container struct {
	// put the list of your services here
	// for example
	//  log *log.Logger

	// also, you can describe your services in a separate container
	// repositories RepositoryContainer
}

// this is a separate container
// type RepositoryContainer {
// 	entityRepository domain.EntityRepository
// }
