package request

import (
	types "belajar_api/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var (
		author = map[int64]*types.Author{}
		book = map[int64]*types.Book{}
		publisher = map[int64]*types.Publisher{}
		borrower = map[int64]*types.Borrower{}
		user = map[int64]*types.User{}
		
	)



func createUser(c *gin.Context) {
	lenUser := int64(len(user)+1)
	name := c.Query("name")
	firstname := c.Query("firstname")
	username := c.Query("username")
	lastname := c.Query("lastname")
	email := c.Query("email")
	password := c.Query("password")

	if name != "" || firstname != "" || lastname != "" || username != "" || email != "" || password != "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "please fill all the required field (Name, FirstName, LastName, Username, Email, Password)"})
	}
	for _, u := range user {
		if u.Username == username {
			c.JSON(http.StatusConflict, gin.H{"error": "Username already exists"})
			return 
		}
	}
	newUser := &types.User {ID: lenUser, Name:name, FirstName:firstname, LastName: lastname, Username: username, Email:email, Password:password}

	user[int64(lenUser)] = newUser


	c.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully", "user": newUser,
	})
}

func createBorrower(c *gin.Context) {
	lenBorrow := int64(len(borrower)+1)
	userId := c.Query("userId")
	endDate := c.Query("endDate")
	startDate := c.Query("startDate")
	bookId  := c.Query("bookId")


	

	userIdInt, err := strconv.ParseInt(userId, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid userId"})
		return
	}

	bookIdInt, err := strconv.ParseInt(bookId, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid bookId"})
		return
	}

	for _, b := range borrower {
		if b.Book.ID == bookIdInt {
			c.JSON(http.StatusConflict, gin.H{
				"message": "Book already borrowed",
			})
			return
		}
	}

	userData, ok := user[userIdInt]
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	bookData, ok := book[bookIdInt]
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	newBorrower := &types.Borrower {
		ID: lenBorrow,
		User: *userData,
		EndDate: endDate,
		StartDate: startDate,
		Book: *bookData,

	}

	borrower[int64(lenBorrow)] = newBorrower

	c.JSON(http.StatusCreated, gin.H{
		"message":  "Borrower created successfully",
		"borrower": newBorrower,
	})
}


func createPubliser(c *gin.Context) {
	lenPublisher := int64(len(publisher)+1)
	name := c.Query("name")
	address := c.Query("address")
	bookId := c.Query("bookId")

	bookIdInt, err := strconv.ParseInt(bookId, 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid type bookId",
		})
	}
	for _, p := range publisher {
		if p.Name == name {
			c.JSON(http.StatusConflict, gin.H{
				"message": "publisher already exists",
			})
			return
		}
	}
	retBook := book[bookIdInt]


	newPublisher := &types.Publisher{
		ID: lenPublisher,
		Name: name,
		Address: address,
		Books: append([]types.Book{}, *retBook),
	}


	publisher[int64(lenPublisher)] = newPublisher

	c.JSON(http.StatusCreated, gin.H{
		"message": "Publisher created successfully",
		"publisher": newPublisher,
	})
	
}

func createBook(c *gin.Context) {
	lenBook := int64(len(book)+1)
	name := c.Query("name")
	description := c.Query("description")
	publisherid := c.Query("publisherId")
	authorid := c.Query("authorId")
	
	publisherIdInt, err := strconv.ParseInt(publisherid, 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid type BookId",
		})
		return
	}

	authorIdInt, err := strconv.ParseInt(authorid, 10, 64)

	
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid type authorId",
		})
	}

	var isborrowed bool
	if c.Query("isborrowed") == "" {
		isborrowed = false
	}else {
		isborrowed = true
	}
	
	for _, b := range book {
		if b.Name == name && b.AuthorID == authorIdInt {
			c.JSON(http.StatusConflict, gin.H{
				"message": "Book already exists",
			})
			return 
		}
	}

	newBook := &types.Book{
		ID: lenBook,
		Name: name,
		Description: description,
		PublisherID: publisherIdInt,
		AuthorID: authorIdInt,
		IsBorrowed: isborrowed,
	}
	book[lenBook] = newBook

	c.JSON(http.StatusOK, gin.H{
		"message": "Book created successfully",
	})
}


func createAuthor(c *gin.Context) {
	lenAuthor := int64(len(book)+1)
	name := c.Query("name")
	email := c.Query("email")
	bookId := c.Query("bookId")


	bookIdInt, err := strconv.ParseInt(bookId, 10, 64)

	if err != nil {
		c.JSON(http.StatusConflict, gin.H{
			"message": "Invalid type bookId",
		})
		return
	}

	getBookData := book[bookIdInt]

	newAuthor := &types.Author{
		ID: lenAuthor,
		Name: name,
		Email: email,
		Book: append([]types.Book{}, *getBookData),
	}

	author[lenAuthor] = newAuthor


	c.JSON(http.StatusOK, gin.H{
		"message": "Author created successfully",
		"author": newAuthor,
	})

}


// Controllers for get requests
func getUsers(c *gin.Context) {
	c.JSON(http.StatusOK, user)
}


func getUser(c *gin.Context) {
	id := c.Query("id")



	ids, err := strconv.ParseInt(id, 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id types"})

		return
	}

	user, ok := user[ids]
	

	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return 
	}

	
	c.JSON(http.StatusOK, user)
}
func getBorrowers(c *gin.Context) {
	c.JSON(http.StatusOK, borrower)
}


func getBorrower(c *gin.Context) {
	id := c.Query("id")



	ids, err := strconv.ParseInt(id, 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id types"})

		return
	}

	borrower, ok := borrower[ids]
	

	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Borrower not found"})
		return 
	}

	
	c.JSON(http.StatusOK, borrower)
}
func getPublishers(c *gin.Context) {
	c.JSON(http.StatusOK, publisher)
}


func getPublisher(c *gin.Context) {
	id := c.Query("id")



	ids, err := strconv.ParseInt(id, 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id types"})

		return
	}

	publisher, ok := publisher[ids]
	

	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "publisher not found"})
		return 
	}

	
	c.JSON(http.StatusOK, publisher)
}
func getBooks(c *gin.Context) {
	c.JSON(http.StatusOK, book)
}


func getBook(c *gin.Context) {
	id := c.Query("id")



	ids, err := strconv.ParseInt(id, 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id types"})

		return
	}

	book, ok := book[ids]
	

	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "book not found"})
		return 
	}

	
	c.JSON(http.StatusOK, book)
}
func getAuthors(c *gin.Context) {
	c.JSON(http.StatusOK, author)
}


func getAuthor(c *gin.Context) {
	id := c.Query("id")



	ids, err := strconv.ParseInt(id, 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id types"})

		return
	}

	Author, ok := author[ids]
	

	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Author not found"})
		return 
	}

	
	c.JSON(http.StatusOK, Author)
}

func updateUser(c *gin.Context) {
	id := c.Query("id")
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid type id",
		})
	}
	name := c.Query("name")
	firstname := c.Query("firstname")
	lastname := c.Query("lastname")
	username := c.Query("username")
	email := c.Query("email")
	password := c.Query("password")

	getUserData := user[idInt]
	



	if name == "" {
		name = getUserData.Name	
	}else if firstname == "" {
		firstname = getUserData.FirstName
	}else if lastname == "" {
		lastname = getUserData.LastName
	}else if username == "" {
		username = getUserData.Username
	}else if email == "" {
		email = getUserData.Email
	}else if password == "" {
		password = getUserData.Password
	}


	newUserdata := &types.User{
		ID: idInt,
		Name: name,
		FirstName: firstname,
		LastName: lastname,
		Username: username,
		Email: email,
		Password: password,
	}


	user[idInt] = newUserdata


	c.JSON(http.StatusOK, gin.H{
		"message": "User update successfully",
	})

}







