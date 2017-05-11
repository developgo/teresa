package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"
	"strings"

	loads "github.com/go-openapi/loads"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"
	security "github.com/go-openapi/runtime/security"
	spec "github.com/go-openapi/spec"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/luizalabs/teresa-api/restapi/operations/apps"
	"github.com/luizalabs/teresa-api/restapi/operations/deployments"
	"github.com/luizalabs/teresa-api/restapi/operations/teams"
	"github.com/luizalabs/teresa-api/restapi/operations/users"
)

// NewTeresaAPI creates a new Teresa instance
func NewTeresaAPI(spec *loads.Document) *TeresaAPI {
	return &TeresaAPI{
		handlers:        make(map[string]map[string]http.Handler),
		formats:         strfmt.Default,
		defaultConsumes: "application/json",
		defaultProduces: "application/json",
		ServerShutdown:  func() {},
		spec:            spec,
	}
}

/*TeresaAPI The Teresa PaaS API */
type TeresaAPI struct {
	spec            *loads.Document
	context         *middleware.Context
	handlers        map[string]map[string]http.Handler
	formats         strfmt.Registry
	defaultConsumes string
	defaultProduces string
	// JSONConsumer registers a consumer for a "application/json" mime type
	JSONConsumer runtime.Consumer
	// MultipartformConsumer registers a consumer for a "multipart/form-data" mime type
	MultipartformConsumer runtime.Consumer

	// BinProducer registers a producer for a "application/octet-stream" mime type
	BinProducer runtime.Producer
	// JSONProducer registers a producer for a "application/json" mime type
	JSONProducer runtime.Producer

	// TokenHeaderAuth registers a function that takes a token and returns a principal
	// it performs authentication based on an api key Authorization provided in the header
	TokenHeaderAuth func(string) (interface{}, error)

	// APIKeyAuth registers a function that takes a token and returns a principal
	// it performs authentication based on an api key token provided in the query
	APIKeyAuth func(string) (interface{}, error)

	// TeamsAddUserToTeamHandler sets the operation handler for the add user to team operation
	TeamsAddUserToTeamHandler teams.AddUserToTeamHandler
	// AppsCreateAppHandler sets the operation handler for the create app operation
	AppsCreateAppHandler apps.CreateAppHandler
	// DeploymentsCreateDeploymentHandler sets the operation handler for the create deployment operation
	DeploymentsCreateDeploymentHandler deployments.CreateDeploymentHandler
	// TeamsCreateTeamHandler sets the operation handler for the create team operation
	TeamsCreateTeamHandler teams.CreateTeamHandler
	// UsersCreateUserHandler sets the operation handler for the create user operation
	UsersCreateUserHandler users.CreateUserHandler
	// TeamsDeleteTeamHandler sets the operation handler for the delete team operation
	TeamsDeleteTeamHandler teams.DeleteTeamHandler
	// UsersDeleteUserHandler sets the operation handler for the delete user operation
	UsersDeleteUserHandler users.DeleteUserHandler
	// AppsGetAppDetailsHandler sets the operation handler for the get app details operation
	AppsGetAppDetailsHandler apps.GetAppDetailsHandler
	// AppsGetAppLogsHandler sets the operation handler for the get app logs operation
	AppsGetAppLogsHandler apps.GetAppLogsHandler
	// AppsGetAppsHandler sets the operation handler for the get apps operation
	AppsGetAppsHandler apps.GetAppsHandler
	// DeploymentsGetDeploymentsHandler sets the operation handler for the get deployments operation
	DeploymentsGetDeploymentsHandler deployments.GetDeploymentsHandler
	// TeamsGetTeamDetailHandler sets the operation handler for the get team detail operation
	TeamsGetTeamDetailHandler teams.GetTeamDetailHandler
	// TeamsGetTeamsHandler sets the operation handler for the get teams operation
	TeamsGetTeamsHandler teams.GetTeamsHandler
	// UsersGetUserDetailsHandler sets the operation handler for the get user details operation
	UsersGetUserDetailsHandler users.GetUserDetailsHandler
	// UsersGetUsersHandler sets the operation handler for the get users operation
	UsersGetUsersHandler users.GetUsersHandler
	// AppsPartialUpdateAppHandler sets the operation handler for the partial update app operation
	AppsPartialUpdateAppHandler apps.PartialUpdateAppHandler
	// AppsUpdateAppHandler sets the operation handler for the update app operation
	AppsUpdateAppHandler apps.UpdateAppHandler
	// AppsUpdateAppAutoScaleHandler sets the operation handler for the update app auto scale operation
	AppsUpdateAppAutoScaleHandler apps.UpdateAppAutoScaleHandler
	// AppsUpdateAppScaleHandler sets the operation handler for the update app scale operation
	AppsUpdateAppScaleHandler apps.UpdateAppScaleHandler
	// TeamsUpdateTeamHandler sets the operation handler for the update team operation
	TeamsUpdateTeamHandler teams.UpdateTeamHandler
	// UsersUpdateUserHandler sets the operation handler for the update user operation
	UsersUpdateUserHandler users.UpdateUserHandler

	// ServeError is called when an error is received, there is a default handler
	// but you can set your own with this
	ServeError func(http.ResponseWriter, *http.Request, error)

	// ServerShutdown is called when the HTTP(S) server is shut down and done
	// handling all active connections and does not accept connections any more
	ServerShutdown func()

	// Custom command line argument groups with their descriptions
	CommandLineOptionsGroups []swag.CommandLineOptionsGroup

	// User defined logger function.
	Logger func(string, ...interface{})
}

// SetDefaultProduces sets the default produces media type
func (o *TeresaAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *TeresaAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *TeresaAPI) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *TeresaAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *TeresaAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *TeresaAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *TeresaAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the TeresaAPI
func (o *TeresaAPI) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}

	if o.MultipartformConsumer == nil {
		unregistered = append(unregistered, "MultipartformConsumer")
	}

	if o.BinProducer == nil {
		unregistered = append(unregistered, "BinProducer")
	}

	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}

	if o.TokenHeaderAuth == nil {
		unregistered = append(unregistered, "AuthorizationAuth")
	}

	if o.APIKeyAuth == nil {
		unregistered = append(unregistered, "TokenAuth")
	}

	if o.TeamsAddUserToTeamHandler == nil {
		unregistered = append(unregistered, "teams.AddUserToTeamHandler")
	}

	if o.AppsCreateAppHandler == nil {
		unregistered = append(unregistered, "apps.CreateAppHandler")
	}

	if o.DeploymentsCreateDeploymentHandler == nil {
		unregistered = append(unregistered, "deployments.CreateDeploymentHandler")
	}

	if o.TeamsCreateTeamHandler == nil {
		unregistered = append(unregistered, "teams.CreateTeamHandler")
	}

	if o.UsersCreateUserHandler == nil {
		unregistered = append(unregistered, "users.CreateUserHandler")
	}

	if o.TeamsDeleteTeamHandler == nil {
		unregistered = append(unregistered, "teams.DeleteTeamHandler")
	}

	if o.UsersDeleteUserHandler == nil {
		unregistered = append(unregistered, "users.DeleteUserHandler")
	}

	if o.AppsGetAppDetailsHandler == nil {
		unregistered = append(unregistered, "apps.GetAppDetailsHandler")
	}

	if o.AppsGetAppLogsHandler == nil {
		unregistered = append(unregistered, "apps.GetAppLogsHandler")
	}

	if o.AppsGetAppsHandler == nil {
		unregistered = append(unregistered, "apps.GetAppsHandler")
	}

	if o.DeploymentsGetDeploymentsHandler == nil {
		unregistered = append(unregistered, "deployments.GetDeploymentsHandler")
	}

	if o.TeamsGetTeamDetailHandler == nil {
		unregistered = append(unregistered, "teams.GetTeamDetailHandler")
	}

	if o.TeamsGetTeamsHandler == nil {
		unregistered = append(unregistered, "teams.GetTeamsHandler")
	}

	if o.UsersGetUserDetailsHandler == nil {
		unregistered = append(unregistered, "users.GetUserDetailsHandler")
	}

	if o.UsersGetUsersHandler == nil {
		unregistered = append(unregistered, "users.GetUsersHandler")
	}

	if o.AppsPartialUpdateAppHandler == nil {
		unregistered = append(unregistered, "apps.PartialUpdateAppHandler")
	}

	if o.AppsUpdateAppHandler == nil {
		unregistered = append(unregistered, "apps.UpdateAppHandler")
	}

	if o.AppsUpdateAppAutoScaleHandler == nil {
		unregistered = append(unregistered, "apps.UpdateAppAutoScaleHandler")
	}

	if o.AppsUpdateAppScaleHandler == nil {
		unregistered = append(unregistered, "apps.UpdateAppScaleHandler")
	}

	if o.TeamsUpdateTeamHandler == nil {
		unregistered = append(unregistered, "teams.UpdateTeamHandler")
	}

	if o.UsersUpdateUserHandler == nil {
		unregistered = append(unregistered, "users.UpdateUserHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *TeresaAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *TeresaAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {

	result := make(map[string]runtime.Authenticator)
	for name, scheme := range schemes {
		switch name {

		case "token_header":

			result[name] = security.APIKeyAuth(scheme.Name, scheme.In, o.TokenHeaderAuth)

		case "api_key":

			result[name] = security.APIKeyAuth(scheme.Name, scheme.In, o.APIKeyAuth)

		}
	}
	return result

}

// ConsumersFor gets the consumers for the specified media types
func (o *TeresaAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {

	result := make(map[string]runtime.Consumer)
	for _, mt := range mediaTypes {
		switch mt {

		case "application/json":
			result["application/json"] = o.JSONConsumer

		case "multipart/form-data":
			result["multipart/form-data"] = o.MultipartformConsumer

		}
	}
	return result

}

// ProducersFor gets the producers for the specified media types
func (o *TeresaAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {

	result := make(map[string]runtime.Producer)
	for _, mt := range mediaTypes {
		switch mt {

		case "application/octet-stream":
			result["application/octet-stream"] = o.BinProducer

		case "application/json":
			result["application/json"] = o.JSONProducer

		}
	}
	return result

}

// HandlerFor gets a http.Handler for the provided operation method and path
func (o *TeresaAPI) HandlerFor(method, path string) (http.Handler, bool) {
	if o.handlers == nil {
		return nil, false
	}
	um := strings.ToUpper(method)
	if _, ok := o.handlers[um]; !ok {
		return nil, false
	}
	h, ok := o.handlers[um][path]
	return h, ok
}

func (o *TeresaAPI) initHandlerCache() {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["POST"] == nil {
		o.handlers[strings.ToUpper("POST")] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/teams/{team_name}/users"] = teams.NewAddUserToTeam(o.context, o.TeamsAddUserToTeamHandler)

	if o.handlers["POST"] == nil {
		o.handlers[strings.ToUpper("POST")] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/apps"] = apps.NewCreateApp(o.context, o.AppsCreateAppHandler)

	if o.handlers["POST"] == nil {
		o.handlers[strings.ToUpper("POST")] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/apps/{app_name}/deployments"] = deployments.NewCreateDeployment(o.context, o.DeploymentsCreateDeploymentHandler)

	if o.handlers["POST"] == nil {
		o.handlers[strings.ToUpper("POST")] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/teams"] = teams.NewCreateTeam(o.context, o.TeamsCreateTeamHandler)

	if o.handlers["POST"] == nil {
		o.handlers[strings.ToUpper("POST")] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/users"] = users.NewCreateUser(o.context, o.UsersCreateUserHandler)

	if o.handlers["DELETE"] == nil {
		o.handlers[strings.ToUpper("DELETE")] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/teams/{team_id}"] = teams.NewDeleteTeam(o.context, o.TeamsDeleteTeamHandler)

	if o.handlers["DELETE"] == nil {
		o.handlers[strings.ToUpper("DELETE")] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/users/{user_id}"] = users.NewDeleteUser(o.context, o.UsersDeleteUserHandler)

	if o.handlers["GET"] == nil {
		o.handlers[strings.ToUpper("GET")] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/apps/{app_name}"] = apps.NewGetAppDetails(o.context, o.AppsGetAppDetailsHandler)

	if o.handlers["GET"] == nil {
		o.handlers[strings.ToUpper("GET")] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/apps/{app_name}/logs"] = apps.NewGetAppLogs(o.context, o.AppsGetAppLogsHandler)

	if o.handlers["GET"] == nil {
		o.handlers[strings.ToUpper("GET")] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/apps"] = apps.NewGetApps(o.context, o.AppsGetAppsHandler)

	if o.handlers["GET"] == nil {
		o.handlers[strings.ToUpper("GET")] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/teams/{team_id}/apps/{app_id}/deployments"] = deployments.NewGetDeployments(o.context, o.DeploymentsGetDeploymentsHandler)

	if o.handlers["GET"] == nil {
		o.handlers[strings.ToUpper("GET")] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/teams/{team_id}"] = teams.NewGetTeamDetail(o.context, o.TeamsGetTeamDetailHandler)

	if o.handlers["GET"] == nil {
		o.handlers[strings.ToUpper("GET")] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/teams"] = teams.NewGetTeams(o.context, o.TeamsGetTeamsHandler)

	if o.handlers["GET"] == nil {
		o.handlers[strings.ToUpper("GET")] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/users/{user_id}"] = users.NewGetUserDetails(o.context, o.UsersGetUserDetailsHandler)

	if o.handlers["GET"] == nil {
		o.handlers[strings.ToUpper("GET")] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/users"] = users.NewGetUsers(o.context, o.UsersGetUsersHandler)

	if o.handlers["PATCH"] == nil {
		o.handlers[strings.ToUpper("PATCH")] = make(map[string]http.Handler)
	}
	o.handlers["PATCH"]["/apps/{app_name}"] = apps.NewPartialUpdateApp(o.context, o.AppsPartialUpdateAppHandler)

	if o.handlers["PUT"] == nil {
		o.handlers[strings.ToUpper("PUT")] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/apps/{app_name}"] = apps.NewUpdateApp(o.context, o.AppsUpdateAppHandler)

	if o.handlers["PUT"] == nil {
		o.handlers[strings.ToUpper("PUT")] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/apps/{app_name}/autoScale"] = apps.NewUpdateAppAutoScale(o.context, o.AppsUpdateAppAutoScaleHandler)

	if o.handlers["PUT"] == nil {
		o.handlers[strings.ToUpper("PUT")] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/apps/{app_name}/scale"] = apps.NewUpdateAppScale(o.context, o.AppsUpdateAppScaleHandler)

	if o.handlers["PUT"] == nil {
		o.handlers[strings.ToUpper("PUT")] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/teams/{team_id}"] = teams.NewUpdateTeam(o.context, o.TeamsUpdateTeamHandler)

	if o.handlers["PUT"] == nil {
		o.handlers[strings.ToUpper("PUT")] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/users/{user_id}"] = users.NewUpdateUser(o.context, o.UsersUpdateUserHandler)

}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *TeresaAPI) Serve(builder middleware.Builder) http.Handler {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}

	return o.context.APIHandler(builder)
}
