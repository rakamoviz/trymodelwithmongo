module github.com/rakamoviz/trymodelwithmongo

go 1.15

replace bitbucket.org/rappinc/gohttp => /home/rcokorda/Projects/rappinc/gohttp

replace bitbucket.org/rappinc/apm => /home/rcokorda/Projects/rappinc/apm/apm

require (
	bitbucket.org/rappinc/gohttp v1.2.0
	github.com/dropbox/godropbox v0.0.0-20200228041828-52ad444d3502 // indirect
	github.com/facebookgo/stack v0.0.0-20160209184415-751773369052 // indirect
	github.com/facebookgo/stackerr v0.0.0-20150612192056-c2fcf88613f4 // indirect
	github.com/labstack/echo v3.3.10+incompatible
	github.com/segmentio/encoding v0.2.17 // indirect
	github.com/sirupsen/logrus v1.4.2
	go.mongodb.org/mongo-driver v1.5.2
)
