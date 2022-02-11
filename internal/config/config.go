//Package config store config information in balance service
package config

//Config structure represents config structure in balance service
type Config struct {
	PostgresURL       string `env:"POSTGRESURL" envDefault:"postgresql://postgres:passwd@localhost:5432/test"`
	BalanceServerPort string `env:"BALANCESERVER" envDefault:"localhost:8085"`
}
