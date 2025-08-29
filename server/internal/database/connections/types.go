package connections

type Connection interface {
	User() string
	Password() string
	Database() string
	Host() string
}
