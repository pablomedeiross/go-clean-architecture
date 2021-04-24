package configuration

import (
	adapterDB "user-api/adapter/db"
	adapterHttp "user-api/adapter/http"
	"user-api/db/inmemory"
	"user-api/external/http"

	"user-api/external/db"
	"user-api/usecase"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

const (
	profile_local                = "local"
	profile_production           = "production"
	error_unknown_profile        = "Error profile don't exists"
	error_register_prefix        = "Error to create new "
	error_register_sufix         = " in service configuration"
	error_create_nosql_db        = error_register_prefix + "NOSQLDB" + error_register_sufix
	error_create_user_repository = error_register_prefix + "UserRepository" + error_register_sufix
	error_create_createuser      = error_register_prefix + "usecase CreateUser" + error_register_sufix
	error_create_removeuser      = error_register_prefix + "usecase RemoveUser" + error_register_sufix
	error_create_controller      = error_register_prefix + "Controller" + error_register_sufix
	error_create_handler         = error_register_prefix + "Handler" + error_register_sufix
	error_to_read_configuration  = "Error to read configuration file"
	error_to_configure_services  = "Error to configure services in application configuration"
)

//AppStarter represents an application configuration ready for be started
type AppStarter struct {
	engine *gin.Engine
}

//Start start http server application with selected profile
func (init *AppStarter) Start() {
	init.engine.Run()
}

// Create new AppInitializer with profile configuration selected by client
func NewAppStarter(profile string) (*AppStarter, error) {

	var engine *gin.Engine
	var err error

	if profile == profile_local {
		engine, err = configureLocalExecution()

	} else if profile == profile_production {
		engine, err = configureProductionExecution()
	}

	if err != nil {
		return nil, errors.Wrap(err, error_to_configure_services)
	}

	return &AppStarter{engine: engine}, nil
}

// configureLocalExecution start a local mongodb database
// and configure application services for a local execution
func configureLocalExecution() (*gin.Engine, error) {

	memoryDB := inmemory.NewInMemoryMongoDB()
	memoryDB.Start()

	return registerServices(memoryDB.URI(), memoryDB.Name())
}

// configureProductionExecution start reade properties file for find out mongodb database configurations
// and configure application services for a production execution
func configureProductionExecution() (*gin.Engine, error) {

	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()

	if err != nil {
		return nil, errors.Wrap(err, error_to_read_configuration)
	}

	dbUrl := viper.GetString("db.url")
	dbName := viper.GetString("db.name")

	return registerServices(dbUrl, dbName)
}

// registerServices configure dependencies in application services
// and return a gin.Engine with handler configured
func registerServices(dbUrl string, dbName string) (*gin.Engine, error) {

	nosqlDB, err := db.NewNoSQLDB(dbUrl, dbName)

	if err != nil {
		return nil, errors.Wrap(err, error_create_nosql_db)
	}

	userRepository, err := adapterDB.NewUserRepository(&nosqlDB)

	if err != nil {
		return nil, errors.Wrap(err, error_create_user_repository)
	}

	createUserUsecase, err := usecase.NewCreateUser(&userRepository)

	if err != nil {
		return nil, errors.Wrap(err, error_create_createuser)
	}

	removeUsecase, err := usecase.NewRemoveUser(userRepository)

	if err != nil {
		return nil, errors.Wrap(err, error_create_removeuser)
	}

	controller, err := adapterHttp.NewHttpController(&createUserUsecase, &removeUsecase)

	if err != nil {
		return nil, errors.Wrap(err, error_create_controller)
	}

	handler, err := http.NewHandler(&controller)

	if err != nil {
		return nil, errors.Wrap(err, error_create_handler)
	}

	return http.CreateEngineWithRoutes(&handler), nil
}
