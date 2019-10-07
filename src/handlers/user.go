package handlers

import (
	"fmt"
	"hilfling-oauth/database"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

const (
	minPasswordLength = 8
	minUsernameLength = 6
	requiredIntegers  = "1234567890"
	requiredSpecial   = "!#$%&/()=?|[]{}*@^-_.:;‚<>"
	requiredLowercase = "qwertyuiopåasdfghjkløæzxcvbnm"
	requiredUppercase = "QWERTYUIOPÅASDFGHJKLØÆZXCVBNM"
)

type baseUser struct {
	Id              int    `json:"id" binding:"required"`
	Username        string `json:"username" binding:"required"`
	FullName        string `json:"full_name" binding:"required"`
	IsDisabled      bool   `json:"is_disabled" binding:"required"`
	SecurityLevelId int    `json:"security_level_id" binding:"required"`
}

type user struct {
	baseUser
	Positions []string `json:"positions" binding:"required"`
}

type securityUser struct {
	baseUser
	PasswordHash string `json:"password_hash" binding:"required"`
	PasswordSalt string `json:"password_salt" binding:"required"`
}

func getUsers() ([]baseUser, error) {
	const q = `SELECT id, username, full_name, is_disabled, security_level_id FROM fg_user;`
	rows := database.Query(q)
	results := make([]baseUser, 0)
	for rows.Next() { // Loop through all DB rows
		var id int
		var username string
		var full_name string
		var is_disabled bool
		var security_level_id int
		err := rows.Scan(&id, &username, &full_name, &is_disabled, &security_level_id)
		if err != nil {
			return nil, err
		}
		results = append(results, baseUser{id, username, full_name, is_disabled, security_level_id})
	}
	return results, nil
}

func GetUsers(ctx *gin.Context) {
	results, err := getUsers()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "internal error: " + err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, results)
}

func hashAndSalt(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func checkPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func nameify(name string) string {
	splitName := strings.Split(name, "%20")
	return strings.Join(splitName, " ")
}

func signup(username, name, password string) (bool, string) {
	// Create query string
	q := `
        INSERT INTO fg_user(username, full_name, password_hash, security_level_id, is_disabled)
          VALUES('%s', '%s', '%s', (SELECT id FROM security_level WHERE level='ALLE'), DEFAULT);
        `
	q = fmt.Sprintf(q, username, name, password)

	// Insert into db
	err := database.Insert(q)
	if err != nil {
		pqErr := err.(*pq.Error) // Convert to *pq.Error to be able to get error code

		errDesc := "Error: " + err.Error()
		fmt.Println(pqErr.Code.Name())
		if pqErr.Code == pq.ErrorCode(23505) { // User already exists
			errDesc = "Error: User already exists"
		}
		return false, errDesc
	}
	return true, "User created"
}

func Signup(ctx *gin.Context) {
	// Read body
	buf := make([]byte, 1024)
	num, _ := ctx.Request.Body.Read(buf) // Ignore error, will most likely get EOF (buffer too big)
	reqBody := string(buf[0:num])

	// Get username and password from body
	s := strings.Split(reqBody, "&")
	ctx.JSON(http.StatusOK, s)
	username, name, password := strings.Split(s[0], "=")[1], strings.Split(s[1], "=")[1], strings.Split(s[2], "=")[1]

	// Fulfill requirements
	reqStr := ""
	if len(username) < minUsernameLength {
		reqStr += fmt.Sprintf("Username must be at least %d characters long \n", minUsernameLength)
	}
	if len(password) < minPasswordLength {
		reqStr += fmt.Sprintf("Password must be at least %d characters long \n", minPasswordLength)
	}
	if !strings.ContainsAny(password, requiredLowercase) {
		reqStr += "Password must contain at least one lowercase letter \n"
	}
	if !strings.ContainsAny(password, requiredUppercase) {
		reqStr += "Password must contain at least one uppercase letter \n"
	}
	if !strings.ContainsAny(password, requiredIntegers) {
		reqStr += "Password must contain at least one number \n"
	}
	if !strings.ContainsAny(password, requiredSpecial) {
		reqStr += fmt.Sprintf("Password must contain at least one of following characters: '%s' \n",
			requiredSpecial)
	}

	// If any requirements are not fulfilled, don't create user
	if len(reqStr) > 0 {
		ctx.JSON(http.StatusNotAcceptable, reqStr)
		return
	}

	// hash password
	password, err := hashAndSalt(password)

	if err != nil {
		log.Println("Error: ", err.Error())
	}

	if signedUp, str := signup(username, nameify(name), password); signedUp == false {
		ctx.JSON(http.StatusInternalServerError, str)
	} else {
		ctx.JSON(http.StatusOK, str)
	}

}
