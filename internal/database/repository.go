package database

//go:generate mockgen -source=$GOFILE -destination=./repository_mock.go -package=$GOPACKAGE

type DataReader interface {
	Get() (*Data, error)
}

type DataWriter interface {
	Store(d *Data) error
}

type DataRepository interface {
	DataReader
	DataWriter
}
