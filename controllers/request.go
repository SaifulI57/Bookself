package request

import (
	types "belajar_api/models"
	"net/http"
	"strconv"
	"sync"

	"github.com/gin-gonic/gin"
)

type AuthorDB struct {
	mu     sync.Mutex
	author []types.Author
}
type BookDB struct {
	mu   sync.Mutex
	book []types.Book
}
type PublisherDB struct {
	mu        sync.Mutex
	publisher []types.Publisher
}
type BorrowerDB struct {
	mu       sync.Mutex
	borrower []types.Borrower
}
type UserDB struct {
	mu   sync.Mutex
	user []types.User
}

func (db *AuthorDB) addAuthor(author types.Author) {
	db.mu.Lock()

	defer db.mu.Unlock()

	db.author = append(db.author, author)
}
func (db *BookDB) addBook(book types.Book) {
	db.mu.Lock()

	defer db.mu.Unlock()

	db.book = append(db.book, book)
}
func (db *PublisherDB) addPublisher(publisher types.Publisher) {
	db.mu.Lock()

	defer db.mu.Unlock()

	db.publisher = append(db.publisher, publisher)
}
func (db *BorrowerDB) addBorrower(borrower types.Borrower) {
	db.mu.Lock()

	defer db.mu.Unlock()

	db.borrower = append(db.borrower, borrower)
}
func (db *UserDB) addUser(user types.User) {
	db.mu.Lock()

	defer db.mu.Unlock()

	db.user = append(db.user, user)
}

func (db *AuthorDB) getAuthor(id int64) *types.Author {
	db.mu.Lock()

	defer db.mu.Unlock()

	for _, author := range db.author {
		if author.ID == id {
			return &author
		}
	}
	return nil
}

func (db *BookDB) getBook(id int64) *types.Book {
	db.mu.Lock()
	defer db.mu.Unlock()

	for _, book := range db.book {
		if book.ID == id {
			return &book
		}
	}
	return nil
}

func (db *PublisherDB) getPublisher(id int64) *types.Publisher {
	db.mu.Lock()
	defer db.mu.Unlock()

	for _, publisher := range db.publisher {
		if publisher.ID == id {
			return &publisher
		}
	}
	return nil
}

func (db *BorrowerDB) getBorrower(id int64) *types.Borrower {
	db.mu.Lock()

	defer db.mu.Unlock()

	for _, borrow := range db.borrower {
		if borrow.ID == id {
			return &borrow
		}
	}

	return nil
}

func (db *BorrowerDB) getBorrowerBook(id int64) bool {
	db.mu.Lock()
	defer db.mu.Unlock()

	for _, borrower := range db.borrower {
		if borrower.BookID == id {
			return true
		}
	}
	return false
}

func (db *UserDB) getUser(id int64) *types.User {
	db.mu.Lock()
	defer db.mu.Unlock()

	for _, user := range db.user {
		if user.ID == id {
			return &user
		}
	}
	return nil
}

func (db *AuthorDB) updateAuthor(id int64, updated types.Author) bool {
	db.mu.Lock()
	defer db.mu.Unlock()

	for i, author := range db.author {
		if author.ID == id {
			db.author[i] = updated
			return true
		}
	}
	return false
}

func (db *BookDB) updateBook(id int64, updated types.Book) bool {
	db.mu.Lock()
	defer db.mu.Unlock()

	for i, book := range db.book {
		if book.ID == id {
			db.book[i] = updated
			return true
		}
	}
	return false
}

func (db *PublisherDB) updatePublisherInfo(id int64, updated types.Publisher) bool {
	db.mu.Lock()
	defer db.mu.Unlock()

	for i, publisher := range db.publisher {
		if publisher.ID == id {

			db.publisher[i] = updated
			return true
		}
	}
	return false
}

func (db *BorrowerDB) updateBorrower(id int64, updated types.Borrower) bool {
	db.mu.Lock()
	defer db.mu.Unlock()

	for i, borrow := range db.borrower {
		if borrow.ID == id {
			db.borrower[i] = updated
			return true
		}
	}
	return false
}

func (db *UserDB) updateUser(id int64, updated types.User) bool {
	db.mu.Lock()
	defer db.mu.Unlock()

	for i, user := range db.user {
		if user.ID == id {
			db.user[i] = updated
			return true
		}
	}
	return false
}

func (db *AuthorDB) deleteAuthor(id int64) bool {
	db.mu.Lock()
	defer db.mu.Unlock()

	for i, author := range db.author {
		if author.ID == id {
			updateRestData := db.author[i+1:]
			for j, updatedRestData := range updateRestData {
				updateRestData[j].ID = updatedRestData.ID - 1
			}
			db.author = append(db.author[:i], updateRestData...)
			return true
		}
	}
	return false
}
func (db *BookDB) deleteBook(id int64) bool {
	db.mu.Lock()
	defer db.mu.Unlock()

	for i, book := range db.book {
		if book.ID == id {
			updateRestData := db.book[i+1:]
			for j, updatedRestData := range updateRestData {
				updateRestData[j].ID = updatedRestData.ID - 1
			}
			db.book = append(db.book[:i], updateRestData...)
			return true
		}
	}
	return false
}
func (db *PublisherDB) deletePublisher(id int64) bool {
	db.mu.Lock()
	defer db.mu.Unlock()

	for i, publisher := range db.publisher {
		if publisher.ID == id {
			updateRestData := db.publisher[i+1:]
			for j, updatedRestData := range updateRestData {
				updateRestData[j].ID = updatedRestData.ID - 1
			}
			db.publisher = append(db.publisher[:i], updateRestData...)
			return true
		}
	}
	return false
}
func (db *BorrowerDB) deleteBorrower(id int64) bool {
	db.mu.Lock()
	defer db.mu.Unlock()

	for i, borrower := range db.borrower {
		if borrower.ID == id {
			updateRestData := db.borrower[i+1:]
			for j, updatedRestData := range updateRestData {
				updateRestData[j].ID = updatedRestData.ID - 1
			}
			db.borrower = append(db.borrower[:i], updateRestData...)
			return true
		}
	}
	return false
}
func (db *UserDB) deleteUser(id int64) bool {
	db.mu.Lock()
	defer db.mu.Unlock()

	for i, user := range db.user {
		if user.ID == id {
			updateRestData := db.user[i+1:]
			for j, updatedRestData := range updateRestData {
				updateRestData[j].ID = updatedRestData.ID - 1
			}
			db.user = append(db.user[:i], updateRestData...)
			return true
		}
	}
	return false
}

func CreateUser(db *UserDB) gin.HandlerFunc {
	return func(c *gin.Context) {
		lenUser := int64(len(db.user) + 1)
		name := c.Query("name")
		firstname := c.Query("firstname")
		username := c.Query("username")
		lastname := c.Query("lastname")
		email := c.Query("email")
		password := c.Query("password")

		if name != "" || firstname != "" || lastname != "" || username != "" || email != "" || password != "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "please fill all the required field (Name, FirstName, LastName, Username, Email, Password)"})
		}
		for _, u := range db.user {
			if u.Username == username {
				c.JSON(http.StatusConflict, gin.H{"error": "Username already exists"})
				return
			}
		}
		newUser := types.User{ID: lenUser, Name: name, FirstName: firstname, LastName: lastname, Username: username, Email: email, Password: password}

		db.addUser(newUser)

		c.JSON(http.StatusCreated, gin.H{
			"message": "User created successfully", "user": newUser,
		})
	}
}

func CreateBorrower(db *BorrowerDB) gin.HandlerFunc {
	return func(c *gin.Context) {
		lenBorrow := int64(len(db.borrower) + 1)
		userId := c.Query("userId")
		endDate := c.Query("endDate")
		startDate := c.Query("startDate")
		bookId := c.Query("bookId")

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

		for _, b := range db.borrower {
			if b.BookID == bookIdInt {
				c.JSON(http.StatusConflict, gin.H{
					"message": "Book already borrowed",
				})
				return
			}
		}

		newBorrower := types.Borrower{
			ID:        lenBorrow,
			UserID:    userIdInt,
			EndDate:   endDate,
			StartDate: startDate,
			BookID:    bookIdInt,
		}

		db.addBorrower(newBorrower)

		c.JSON(http.StatusCreated, gin.H{
			"message":  "Borrower created successfully",
			"borrower": newBorrower,
		})
	}
}

func CreatePublisher(db *PublisherDB) gin.HandlerFunc {
	return func(c *gin.Context) {
		lenPublisher := int64(len(db.publisher) + 1)
		name := c.Query("name")
		address := c.Query("address")

		for _, p := range db.publisher {
			if p.Name == name {
				c.JSON(http.StatusConflict, gin.H{
					"message": "publisher already exists",
				})
				return
			}
		}

		newPublisher := types.Publisher{
			ID:      lenPublisher,
			Name:    name,
			Address: address,
		}

		db.addPublisher(newPublisher)

		c.JSON(http.StatusCreated, gin.H{
			"message":   "Publisher created successfully",
			"publisher": newPublisher,
		})

	}
}

func CreateBook(db *BookDB) gin.HandlerFunc {
	return func(c *gin.Context) {
		lenBook := int64(len(db.book) + 1)
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

		for _, b := range db.book {
			if b.Name == name && b.AuthorID == authorIdInt {
				c.JSON(http.StatusConflict, gin.H{
					"message": "Book already exists",
				})
				return
			}
		}

		newBook := types.Book{
			ID:          lenBook,
			Name:        name,
			Description: description,
			PublisherID: publisherIdInt,
			AuthorID:    authorIdInt,
			IsBorrowed:  false,
		}
		db.addBook(newBook)
		c.JSON(http.StatusOK, gin.H{
			"message": "Book created successfully",
		})
	}
}

func CreateAuthor(db *AuthorDB) gin.HandlerFunc {
	return func(c *gin.Context) {
		lenAuthor := int64(len(db.author) + 1)
		name := c.Query("name")
		email := c.Query("email")

		newAuthor := types.Author{
			ID:    lenAuthor,
			Name:  name,
			Email: email,
		}
		for _, author := range db.author {
			if author.Name == name || author.Email == email {
				c.JSON(http.StatusConflict, gin.H{
					"error": "Author already exists",
				})
				return
			}
		}
		db.addAuthor(newAuthor)

		c.JSON(http.StatusOK, gin.H{
			"message": "Author created successfully",
			"author":  newAuthor,
		})

	}
}

// Controllers for get requests
func GetUsers(db *UserDB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "success",
			"status":  200,
			"data":    db.user,
		})
	}
}

func GetUserHTTP(db *UserDB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		ids, err := strconv.ParseInt(id, 10, 64)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id types"})

			return
		}

		dataUser := db.getUser(ids)

		if dataUser == nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}

		c.JSON(http.StatusOK, dataUser)
	}
}

func GetBorrowers(db *BorrowerDB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, db.borrower)
	}
}

func GetBorrowerHTTP(db *BorrowerDB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		ids, err := strconv.ParseInt(id, 10, 64)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id types"})

			return
		}

		dataBorrowers := db.getBorrower(ids)

		if dataBorrowers == nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Borrower not found"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "success",
			"status":  200,
			"data":    db.borrower,
		})
	}
}

func GetPublishers(db *PublisherDB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, db.publisher)
	}
}

func GetPublisherHTTP(db *PublisherDB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		ids, err := strconv.ParseInt(id, 10, 64)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id types"})

			return
		}

		dataPublisher := db.getPublisher(ids)

		if dataPublisher == nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "publisher not found"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "success",
			"status":  200,
			"data":    dataPublisher,
		})
	}
}

func GetBooks(db *BookDB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, db.book)
	}
}

func GetBookHTTP(db *BookDB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		ids, err := strconv.ParseInt(id, 10, 64)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id types"})

			return
		}

		dataBook := db.getBook(ids)

		if dataBook == nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "book not found"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "success",
			"status":  200,
			"data":    dataBook,
		})
	}
}
func GetAuthors(db *AuthorDB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, db.author)
	}
}

func GetAuthorHTTP(db *AuthorDB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		ids, err := strconv.ParseInt(id, 10, 64)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id types"})

			return
		}

		dataAuthor := db.getAuthor(ids)

		if dataAuthor == nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Author not found"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "success",
			"status":  200,
			"data":    dataAuthor,
		})
	}
}

func UpdateUserHTTP(db *UserDB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
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

		getUserData := db.user[idInt]

		if name == "" {
			name = getUserData.Name
		} else if firstname == "" {
			firstname = getUserData.FirstName
		} else if lastname == "" {
			lastname = getUserData.LastName
		} else if username == "" {
			username = getUserData.Username
		} else if email == "" {
			email = getUserData.Email
		} else if password == "" {
			password = getUserData.Password
		}

		newUserdata := types.User{
			ID:        idInt,
			Name:      name,
			FirstName: firstname,
			LastName:  lastname,
			Username:  username,
			Email:     email,
			Password:  password,
		}

		dataUser := db.updateUser(idInt, newUserdata)

		c.JSON(http.StatusOK, gin.H{
			"message": "User update successfully",
			"data":    dataUser,
		})

	}
}

func UpdateBorrowerHTTP(db *BorrowerDB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		idInt, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid type id",
			})
		}

		userid := c.Query("userId")
		enddate := c.Query("endDate")
		startdate := c.Query("startDate")
		bookid := c.Query("bookId")

		useridInt, err := strconv.ParseInt(userid, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid type userId",
			})
		}
		bookidInt, err := strconv.ParseInt(bookid, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid type userId",
			})
		}

		newBorrower := types.Borrower{
			ID:        idInt,
			UserID:    useridInt,
			EndDate:   enddate,
			StartDate: startdate,
			BookID:    bookidInt,
		}

		db.updateBorrower(idInt, newBorrower)
		c.JSON(http.StatusOK, gin.H{
			"message": "Borrower updated successfully",
			"data":    newBorrower,
		})
	}
}

func UpdatePublisherHTTP(db *PublisherDB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		if id == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Please fill the id field",
			})
			return
		}
		idInt, err := strconv.ParseInt(id, 10, 64)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid type id",
			})
		}
		name := c.Query("name")
		address := c.Query("address")

		newPublisher := types.Publisher{
			ID:      idInt,
			Name:    name,
			Address: address,
		}

		db.updatePublisherInfo(idInt, newPublisher)

		c.JSON(http.StatusOK, gin.H{
			"message": "successfully update publisher",
			"status":  200,
			"data":    newPublisher,
		})

	}
}

func UpdateBookHTTP(db *BookDB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		idInt, err := strconv.ParseInt(id, 10, 64)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid type id",
			})
		}

		name := c.Query("name")
		description := c.Query("description")
		publisherId := c.Query("publisherId")
		authorId := c.Query("authorId")

		publisherIdInt, err := strconv.ParseInt(publisherId, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid type publisher id",
			})
		}
		authorIdInt, err := strconv.ParseInt(authorId, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid type author id",
			})
		}

		var dbBorrower *BorrowerDB

		isborrowed := dbBorrower.getBorrowerBook(idInt)

		newBook := types.Book{
			ID:          idInt,
			Name:        name,
			Description: description,
			PublisherID: publisherIdInt,
			AuthorID:    authorIdInt,
			IsBorrowed:  isborrowed,
		}

		db.updateBook(idInt, newBook)

		c.JSON(http.StatusOK, gin.H{
			"message": "successfully update book",
			"status":  200,
			"data":    newBook,
		})
	}
}

func UpdateAuthorHTTP(db *AuthorDB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		idInt, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Invalid type id",
			})
			return
		}

		name := c.Query("name")
		email := c.Query("email")

		newAuthor := types.Author{
			ID:    idInt,
			Name:  name,
			Email: email,
		}

		db.updateAuthor(idInt, newAuthor)

		c.JSON(http.StatusOK, gin.H{
			"message": "successfully update Author",
			"status":  200,
			"data":    newAuthor,
		})
	}
}

func DeleteAuthorHTTP(db *AuthorDB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		idInt, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid type id",
			})
			return
		}
		isDeleted := db.deleteAuthor(idInt)

		if isDeleted {
			c.JSON(http.StatusOK, gin.H{
				"message": "Author deleted",
			})
		} else {
			c.JSON(http.StatusNotModified, gin.H{
				"message": "failed to delete Author",
			})
		}
	}
}
func DeleteBookHTTP(db *BookDB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		idInt, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid type id",
			})
			return
		}
		isDeleted := db.deleteBook(idInt)

		if isDeleted {
			c.JSON(http.StatusOK, gin.H{
				"message": "Book deleted",
			})
		} else {
			c.JSON(http.StatusNotModified, gin.H{
				"message": "failed to delete Book",
			})
		}
	}
}
func DeletePublisherHTTP(db *PublisherDB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		idInt, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid type id",
			})
			return
		}
		isDeleted := db.deletePublisher(idInt)

		if isDeleted {
			c.JSON(http.StatusOK, gin.H{
				"message": "Publisher deleted",
			})
		} else {
			c.JSON(http.StatusNotModified, gin.H{
				"message": "failed to delete Author",
			})
		}
	}
}
func DeleteBorrowerHTTP(db *BorrowerDB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		idInt, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid type id",
			})
			return
		}
		isDeleted := db.deleteBorrower(idInt)

		if isDeleted {
			c.JSON(http.StatusOK, gin.H{
				"message": "Borrower deleted",
			})
		} else {
			c.JSON(http.StatusNotModified, gin.H{
				"message": "failed to delete Borrower",
			})
		}
	}
}
func DeleteUserHTTP(db *UserDB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		idInt, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid type id",
			})
			return
		}
		isDeleted := db.deleteUser(idInt)

		if isDeleted {
			c.JSON(http.StatusOK, gin.H{
				"message": "User deleted",
			})
		} else {
			c.JSON(http.StatusNotModified, gin.H{
				"message": "failed to delete User",
			})
		}
	}
}
