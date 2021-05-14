module github.com/rakamoviz/trymodelwithmongo

go 1.15

replace bitbucket.org/rappinc/gohttp => /home/rcokorda/Projects/rappinc/gohttp
replace bitbucket.org/rappinc/apm => /home/rcokorda/Projects/rappinc/apm/apm

require (
	github.com/sirupsen/logrus v1.4.2
	go.mongodb.org/mongo-driver v1.5.2
	bitbucket.org/rappinc/gohttp v1.2.0
)
