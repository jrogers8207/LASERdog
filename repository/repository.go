package repository

import "LASERdog/entities"

type Repository interface {
	CreateTarget()

	CreateOrganization()
	CreatePerson()
	CreateAsset()

	CreateNetwork(network ...entities.Network)
	CreateHost()
	CreatePort()
}
