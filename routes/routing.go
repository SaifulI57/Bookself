package routing

import (
	request "belajar_api/controllers"

	"github.com/gin-gonic/gin"
)


func SetupRoute(router *gin.Engine) {
	authorDB := &request.AuthorDB{}
	bookDB := &request.BookDB{}
	publisherDB := &request.PublisherDB{}
	borrowerDB := &request.BorrowerDB{}
	userDB := &request.UserDB{}


	router.GET("/author", request.GetAuthors(authorDB))
	router.GET("author/:id", request.GetAuthorHTTP(authorDB))
	router.POST("author", request.CreateAuthor(authorDB))
	router.PUT("author/:id", request.UpdateAuthorHTTP(authorDB))
	router.DELETE("author/:id", request.DeleteAuthorHTTP(authorDB))

	
	router.GET("/book", request.GetBooks(bookDB))
	router.GET("book/:id", request.GetBookHTTP(bookDB))
	router.POST("book", request.CreateBook(bookDB))
	router.PUT("book/:id", request.UpdateBookHTTP(bookDB))
	router.DELETE("book/:id", request.DeleteBookHTTP(bookDB))

	router.GET("/publisher", request.GetPublishers(publisherDB))
	router.GET("publisher/:id", request.GetPublisherHTTP(publisherDB))
	router.POST("publisher", request.CreatePublisher(publisherDB))
	router.PUT("publisher/:id", request.UpdatePublisherHTTP(publisherDB))
	router.DELETE("publisher/:id", request.DeletePublisherHTTP(publisherDB))

	router.GET("/borrower", request.GetBorrowers(borrowerDB))
	router.GET("borrower/:id", request.GetBorrowerHTTP(borrowerDB))
	router.POST("borrower", request.CreateBorrower(borrowerDB))
	router.PUT("borrower/:id", request.UpdateBorrowerHTTP(borrowerDB))
	router.DELETE("borrower/:id", request.DeleteBorrowerHTTP(borrowerDB))

	router.GET("/user", request.GetUsers(userDB))
	router.GET("user/:id", request.GetUserHTTP(userDB))
	router.POST("user", request.CreateUser(userDB))
	router.PUT("user/:id", request.UpdateUserHTTP(userDB))
	router.DELETE("user/:id", request.DeleteUserHTTP(userDB))

}