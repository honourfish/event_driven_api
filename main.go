package main

import (
	"log"

	// swagger stuff
	"event/server/restapi"
	"event/server/restapi/operations"

	"github.com/go-openapi/loads"
)

func main() {

	// do the server stuff
	var swaggerSpec *loads.Document
	var err error
	if swaggerSpec, err = loads.Analyzed(restapi.SwaggerJSON, ""); err != nil {
		log.Fatalln(err)
	}

	// create a new API which can be used to create a server
	api := operations.NewEventsAPI(swaggerSpec)

	server := restapi.NewServer(api)
	defer server.Shutdown()

	server.Port = 8080

	//bookHandler := handlers.BookHandler{Persister: mongoDb}

	// /book handlers
	//api.PostBookHandler = operations.PostBookHandlerFunc(bookHandler.HandlePostBook)
	//api.GetBookHandler = operations.GetBookHandlerFunc(bookHandler.HandleGetBook)
	//api.GetBookTitleHandler = operations.GetBookTitleHandlerFunc(bookHandler.HandleGetBookTitle)
	//api.PutBookTitleHandler = operations.PutBookTitleHandlerFunc(bookHandler.HandlePutBookTitle)
	//api.DeleteBookTitleHandler = operations.DeleteBookTitleHandlerFunc(bookHandler.HandleDeleteBookTitle) // keep for acceptance tests

	// /librarian/{username}/book handlers
	//seniorLibrarianHandler := handlers.SeniorLibrarianHandler{Persister: mongoDb}
	//librarianHandler := handlers.LibrarianHandler{Persister: mongoDb}

	//api.PostLibrarianUsernameBookHandler = operations.PostLibrarianUsernameBookHandlerFunc(seniorLibrarianHandler.HandlePostBook)
	//api.DeleteLibrarianUsernameBookTitleAuthorCopiesHandler = operations.DeleteLibrarianUsernameBookTitleAuthorCopiesHandlerFunc(seniorLibrarianHandler.HandleDeleteBook)
	//api.GetLibrarianUsernameBookTitleAuthorHandler = operations.GetLibrarianUsernameBookTitleAuthorHandlerFunc(librarianHandler.HandleGetBook)

	// /librarian/{username}/user handlers
	//api.PostLibrarianUsernameUserHandler = operations.PostLibrarianUsernameUserHandlerFunc(librarianHandler.HandlePostUser)
	//api.DeleteLibrarianUsernameUserUserHandler = operations.DeleteLibrarianUsernameUserUserHandlerFunc(librarianHandler.HandleDeleteUser)
	//api.GetLibrarianUsernameUserUserHandler = operations.GetLibrarianUsernameUserUserHandlerFunc(librarianHandler.HandleGetUser)
	//api.PutLibrarianUsernameUserUserCheckoutHandler = operations.PutLibrarianUsernameUserUserCheckoutHandlerFunc(librarianHandler.HandlePutCheckOut)
	//api.PutLibrarianUsernameUserUserCheckinHandler = operations.PutLibrarianUsernameUserUserCheckinHandlerFunc(librarianHandler.HandlePutCheckIn)

	if err = server.Serve(); err != nil {
		log.Fatalln(err)
	}
}
