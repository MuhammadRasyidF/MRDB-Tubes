package main

import (
	"api-mrdb/creators"
	"api-mrdb/genre"
	"api-mrdb/models"
	"api-mrdb/movie"
	"api-mrdb/rating"
	"api-mrdb/stars"
	"api-mrdb/users"
	"api-mrdb/utils"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	//movie
	router := httprouter.New()
	router.GET("/reviewmovie", GetMovie)
	router.POST("/reviewmovie", Auth(PostMovie))
	router.PUT("/reviewmovie/:id", Auth(UpdateMovie))
	router.DELETE("/reviewmovie/:id", Auth(DeleteMovie))
	////////////////////////////////////////////////////

	//user
	router.GET("/user", Auth(GetUser))
	router.POST("/user", PostUser)
	router.PUT("/user/:id", UpdateUser)
	router.DELETE("/user/:id", Auth(DeleteUser))
	////////////////////////////////////////////////////

	//creator
	router.GET("/creator", Auth(GetCreator))
	router.POST("/creator", Auth(PostCreator))
	router.PUT("/creator/:id", Auth(UpdateCreator))
	router.DELETE("/creator/:id", Auth(DeleteCreator))
	////////////////////////////////////////////////////

	//feedback
	router.GET("/feedback", Auth(GetFeedback))
	router.POST("/feedback", Auth(PostFeedback))
	router.PUT("/feedback)/:id", Auth(UpdateFeedback))
	router.DELETE("/feedback)/:id", Auth(DeleteFeedback))
	////////////////////////////////////////////////////

	//genre
	router.GET("/genre", Auth(GetGenre))
	router.POST("/genre", Auth(PostGenre))
	router.PUT("/genre)/:id", Auth(UpdateGenre))
	router.DELETE("/genre)/:id", Auth(DeleteGenre))
	////////////////////////////////////////////////////

	//genre_category
	router.GET("/genre_category", Auth(GetGenreCategory))
	router.POST("/genre_category", Auth(PostGenreCategory))
	router.PUT("/genre_category)/:id", Auth(UpdateGenreCategory))
	router.DELETE("/genre_category)/:id", Auth(DeleteGenreCategory))
	////////////////////////////////////////////////////

	//rating
	router.GET("/rating", Auth(GetRating))
	router.POST("/rating", Auth(PostRating))
	router.PUT("/rating)/:id", Auth(UpdateRating))
	router.DELETE("/rating)/:id", Auth(DeleteRating))
	////////////////////////////////////////////////////

	//review
	router.GET("/review", Auth(GetReview))
	router.POST("/review", Auth(PostReview))
	router.PUT("/review)/:id", Auth(UpdateReview))
	router.DELETE("/review)/:id", Auth(DeleteReview))
	////////////////////////////////////////////////////

	//star
	router.GET("/star", Auth(GetStar))
	router.POST("/star", Auth(PostStar))
	router.PUT("/star)/:id", Auth(UpdateStar))
	router.DELETE("/star)/:id", Auth(DeleteStar))
	////////////////////////////////////////////////////

	// untuk menampilkan file html di folder public
	router.NotFound = http.FileServer(http.Dir("public"))

	fmt.Println("Server Running at Port 8080")
	log.Fatal(http.ListenAndServe(":"+port, router))
}

//auth
func Auth(h httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		// Get the Basic Authentication credentials
		user, password, hasAuth := r.BasicAuth()

		if hasAuth && user == "admin" && password == "admin" {
			// Delegate request to the given handle
			h(w, r, ps)
		} else {
			// Request Basic Authentication otherwise
			w.Header().Set("WWW-Authenticate", "Basic realm=Restricted")
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		}
	}
}

//------ movie -----//
// Read
// GetMovie
func GetMovie(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	movies, err := movie.GetAll(ctx)

	if err != nil {
		fmt.Println(err)
	}

	utils.ResponseJSON(w, movies, http.StatusOK)
}

// Create
// PostMovie
func PostMovie(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var mov models.Movie

	if err := json.NewDecoder(r.Body).Decode(&mov); err != nil {
		utils.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}

	if err := movie.Insert(ctx, mov); err != nil {
		utils.ResponseJSON(w, err, http.StatusInternalServerError)
		return
	}

	res := map[string]string{
		"status": "Succesfully",
	}

	utils.ResponseJSON(w, res, http.StatusCreated)

}

// UpdateMovie
func UpdateMovie(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var mov models.Movie

	if err := json.NewDecoder(r.Body).Decode(&mov); err != nil {
		utils.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}

	var idMovie = ps.ByName("id")

	if err := movie.Update(ctx, mov, idMovie); err != nil {
		utils.ResponseJSON(w, err, http.StatusInternalServerError)
		return
	}

	res := map[string]string{
		"status": "Succesfully",
	}

	utils.ResponseJSON(w, res, http.StatusCreated)
}

// Delete
// DeleteMovie
func DeleteMovie(w http.ResponseWriter, _ *http.Request, ps httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var idMovie = ps.ByName("id")

	if err := movie.Delete(ctx, idMovie); err != nil {
		kesalahan := map[string]string{
			"error": fmt.Sprintf("%v", err),
		}
		utils.ResponseJSON(w, kesalahan, http.StatusInternalServerError)
		return
	}

	res := map[string]string{
		"status": "Succesfully",
	}

	utils.ResponseJSON(w, res, http.StatusOK)
}

/////////////////////////////////////////////////////////////////////

//--------User----------//
// Read
// GetUser
func GetUser(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	users, err := users.GetAll(ctx)

	if err != nil {
		fmt.Println(err)
	}

	utils.ResponseJSON(w, users, http.StatusOK)
}

// Create
// PostUser
func PostUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var user models.Users

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		utils.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}

	if err := users.Insert(ctx, user); err != nil {
		utils.ResponseJSON(w, err, http.StatusInternalServerError)
		return
	}

	res := map[string]string{
		"status": "Succesfully",
	}

	utils.ResponseJSON(w, res, http.StatusCreated)

}

// UpdateUser
func UpdateUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var user models.Users

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		utils.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}

	var idUser = ps.ByName("uname")

	if err := users.Update(ctx, user, idUser); err != nil {
		utils.ResponseJSON(w, err, http.StatusInternalServerError)
		return
	}

	res := map[string]string{
		"status": "Succesfully",
	}

	utils.ResponseJSON(w, res, http.StatusCreated)
}

// Delete
// DeleteUser
func DeleteUser(w http.ResponseWriter, _ *http.Request, ps httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var idUser = ps.ByName("uname")

	if err := movie.Delete(ctx, idUser); err != nil {
		kesalahan := map[string]string{
			"error": fmt.Sprintf("%v", err),
		}
		utils.ResponseJSON(w, kesalahan, http.StatusInternalServerError)
		return
	}

	res := map[string]string{
		"status": "Succesfully",
	}

	utils.ResponseJSON(w, res, http.StatusOK)
}

/////////////////////////////////////////////////////////////////////

//--------Creator----------//
// Read
// GetCreator
func GetCreator(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	creators, err := creators.GetAll(ctx)

	if err != nil {
		fmt.Println(err)
	}

	utils.ResponseJSON(w, creators, http.StatusOK)
}

// Create
// PostCreator
func PostCreator(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var creator models.Creators

	if err := json.NewDecoder(r.Body).Decode(&creator); err != nil {
		utils.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}

	if err := creators.Insert(ctx, creator); err != nil {
		utils.ResponseJSON(w, err, http.StatusInternalServerError)
		return
	}

	res := map[string]string{
		"status": "Succesfully",
	}

	utils.ResponseJSON(w, res, http.StatusCreated)

}

// UpdateCreator
func UpdateCreator(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var creator models.Creators

	if err := json.NewDecoder(r.Body).Decode(&creator); err != nil {
		utils.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}

	var idCreator = ps.ByName("uname")

	if err := creators.Update(ctx, creator, idCreator); err != nil {
		utils.ResponseJSON(w, err, http.StatusInternalServerError)
		return
	}

	res := map[string]string{
		"status": "Succesfully",
	}

	utils.ResponseJSON(w, res, http.StatusCreated)
}

// Delete
// DeleteCreator
func DeleteCreator(w http.ResponseWriter, _ *http.Request, ps httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var idCreator = ps.ByName("uname")

	if err := movie.Delete(ctx, idCreator); err != nil {
		kesalahan := map[string]string{
			"error": fmt.Sprintf("%v", err),
		}
		utils.ResponseJSON(w, kesalahan, http.StatusInternalServerError)
		return
	}

	res := map[string]string{
		"status": "Succesfully",
	}

	utils.ResponseJSON(w, res, http.StatusOK)
}

/////////////////////////////////////////////////////////////////////

//--------Feedback----------//
// Read
// GetFeedback
func GetFeedback(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	feedback, err := feedbacks.GetAll(ctx)

	if err != nil {
		fmt.Println(err)
	}

	utils.ResponseJSON(w, feedback, http.StatusOK)
}

// Create
// PostFeedback
func PostFeedback(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var feedback models.Feedback

	if err := json.NewDecoder(r.Body).Decode(&feedback); err != nil {
		utils.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}

	if err := feedbacks.Insert(ctx, feedback); err != nil {
		utils.ResponseJSON(w, err, http.StatusInternalServerError)
		return
	}

	res := map[string]string{
		"status": "Succesfully",
	}

	utils.ResponseJSON(w, res, http.StatusCreated)

}

// UpdateFeedback
func UpdateFeedback(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var feedback models.Feedback

	if err := json.NewDecoder(r.Body).Decode(&feedback); err != nil {
		utils.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}

	var idCreator = ps.ByName("id")

	if err := feedbacks.Update(ctx, feedback, idCreator); err != nil {
		utils.ResponseJSON(w, err, http.StatusInternalServerError)
		return
	}

	res := map[string]string{
		"status": "Succesfully",
	}

	utils.ResponseJSON(w, res, http.StatusCreated)
}

// Delete
// DeleteFeedback
func DeleteFeedback(w http.ResponseWriter, _ *http.Request, ps httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var idFeedback = ps.ByName("id")

	if err := movie.Delete(ctx, idFeedback); err != nil {
		kesalahan := map[string]string{
			"error": fmt.Sprintf("%v", err),
		}
		utils.ResponseJSON(w, kesalahan, http.StatusInternalServerError)
		return
	}

	res := map[string]string{
		"status": "Succesfully",
	}

	utils.ResponseJSON(w, res, http.StatusOK)
}

/////////////////////////////////////////////////////////////////////

//--------Genre----------//
// Read
// GetGenre
func GetGenre(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	genre, err := genre.GetAll(ctx)

	if err != nil {
		fmt.Println(err)
	}

	utils.ResponseJSON(w, genre, http.StatusOK)
}

// Create
// PostGenre
func PostGenre(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var genres models.Genre

	if err := json.NewDecoder(r.Body).Decode(&genres); err != nil {
		utils.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}

	if err := genre.Insert(ctx, genres); err != nil {
		utils.ResponseJSON(w, err, http.StatusInternalServerError)
		return
	}

	res := map[string]string{
		"status": "Succesfully",
	}

	utils.ResponseJSON(w, res, http.StatusCreated)

}

// UpdateGenre
func UpdateGenre(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var genres models.Genre

	if err := json.NewDecoder(r.Body).Decode(&genres); err != nil {
		utils.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}

	var idCreator = ps.ByName("id")

	if err := genre.Update(ctx, genres, idCreator); err != nil {
		utils.ResponseJSON(w, err, http.StatusInternalServerError)
		return
	}

	res := map[string]string{
		"status": "Succesfully",
	}

	utils.ResponseJSON(w, res, http.StatusCreated)
}

// Delete
// DeleteGenre
func DeleteGenre(w http.ResponseWriter, _ *http.Request, ps httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var idGenre = ps.ByName("id")

	if err := movie.Delete(ctx, idGenre); err != nil {
		kesalahan := map[string]string{
			"error": fmt.Sprintf("%v", err),
		}
		utils.ResponseJSON(w, kesalahan, http.StatusInternalServerError)
		return
	}

	res := map[string]string{
		"status": "Succesfully",
	}

	utils.ResponseJSON(w, res, http.StatusOK)
}

/////////////////////////////////////////////////////////////////////

//--------Genre-Category----------//
// Read
// GetGenreCategory
func GetGenreCategory(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	genre_category, err := genre_category.GetAll(ctx)

	if err != nil {
		fmt.Println(err)
	}

	utils.ResponseJSON(w, genre_category, http.StatusOK)
}

// Create
// PostGenreCategory
func PostGenreCategory(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var genre_categorys models.Genre_category

	if err := json.NewDecoder(r.Body).Decode(&genre_categorys); err != nil {
		utils.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}

	if err := genre_category.Insert(ctx, genre_categorys); err != nil {
		utils.ResponseJSON(w, err, http.StatusInternalServerError)
		return
	}

	res := map[string]string{
		"status": "Succesfully",
	}

	utils.ResponseJSON(w, res, http.StatusCreated)

}

// UpdateGenreCategory
func UpdateGenreCategory(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var genre_categorys models.Genre_category

	if err := json.NewDecoder(r.Body).Decode(&genre_categorys); err != nil {
		utils.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}

	var idCreator = ps.ByName("id")

	if err := genre_category.Update(ctx, genre_categorys, idCreator); err != nil {
		utils.ResponseJSON(w, err, http.StatusInternalServerError)
		return
	}

	res := map[string]string{
		"status": "Succesfully",
	}

	utils.ResponseJSON(w, res, http.StatusCreated)
}

// Delete
// DeleteGenreCategory
func DeleteGenreCategory(w http.ResponseWriter, _ *http.Request, ps httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var idGenreCategory = ps.ByName("id")

	if err := movie.Delete(ctx, idGenreCategory); err != nil {
		kesalahan := map[string]string{
			"error": fmt.Sprintf("%v", err),
		}
		utils.ResponseJSON(w, kesalahan, http.StatusInternalServerError)
		return
	}

	res := map[string]string{
		"status": "Succesfully",
	}

	utils.ResponseJSON(w, res, http.StatusOK)
}

/////////////////////////////////////////////////////////////////////

//--------Rating----------//
// Read
// GetRating
func GetRating(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	rating, err := rating.GetAll(ctx)

	if err != nil {
		fmt.Println(err)
	}

	utils.ResponseJSON(w, rating, http.StatusOK)
}

// Create
// PostRating
func PostRating(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var ratings models.Rating

	if err := json.NewDecoder(r.Body).Decode(&ratings); err != nil {
		utils.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}

	if err := rating.Insert(ctx, ratings); err != nil {
		utils.ResponseJSON(w, err, http.StatusInternalServerError)
		return
	}

	res := map[string]string{
		"status": "Succesfully",
	}

	utils.ResponseJSON(w, res, http.StatusCreated)

}

// UpdateRating
func UpdateRating(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var ratings models.Rating

	if err := json.NewDecoder(r.Body).Decode(&ratings); err != nil {
		utils.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}

	var idCreator = ps.ByName("id")

	if err := rating.Update(ctx, ratings, idCreator); err != nil {
		utils.ResponseJSON(w, err, http.StatusInternalServerError)
		return
	}

	res := map[string]string{
		"status": "Succesfully",
	}

	utils.ResponseJSON(w, res, http.StatusCreated)
}

// Delete
// DeleteRating
func DeleteRating(w http.ResponseWriter, _ *http.Request, ps httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var idRating = ps.ByName("id")

	if err := movie.Delete(ctx, idRating); err != nil {
		kesalahan := map[string]string{
			"error": fmt.Sprintf("%v", err),
		}
		utils.ResponseJSON(w, kesalahan, http.StatusInternalServerError)
		return
	}

	res := map[string]string{
		"status": "Succesfully",
	}

	utils.ResponseJSON(w, res, http.StatusOK)
}

/////////////////////////////////////////////////////////////////////

//--------Star----------//
// Read
// GetStar
func GetStar(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	star, err := stars.GetAll(ctx)

	if err != nil {
		fmt.Println(err)
	}

	utils.ResponseJSON(w, star, http.StatusOK)
}

// Create
// PostStar
func PostStar(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var star models.Stars

	if err := json.NewDecoder(r.Body).Decode(&star); err != nil {
		utils.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}

	if err := stars.Insert(ctx, star); err != nil {
		utils.ResponseJSON(w, err, http.StatusInternalServerError)
		return
	}

	res := map[string]string{
		"status": "Succesfully",
	}

	utils.ResponseJSON(w, res, http.StatusCreated)

}

// UpdateStar
func UpdateStar(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var star models.Stars

	if err := json.NewDecoder(r.Body).Decode(&star); err != nil {
		utils.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}

	var idStar = ps.ByName("id")

	if err := stars.Update(ctx, star, idStar); err != nil {
		utils.ResponseJSON(w, err, http.StatusInternalServerError)
		return
	}

	res := map[string]string{
		"status": "Succesfully",
	}

	utils.ResponseJSON(w, res, http.StatusCreated)
}

// Delete
// DeleteStar
func DeleteStar(w http.ResponseWriter, _ *http.Request, ps httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var idStar = ps.ByName("id")

	if err := movie.Delete(ctx, idStar); err != nil {
		kesalahan := map[string]string{
			"error": fmt.Sprintf("%v", err),
		}
		utils.ResponseJSON(w, kesalahan, http.StatusInternalServerError)
		return
	}

	res := map[string]string{
		"status": "Succesfully",
	}

	utils.ResponseJSON(w, res, http.StatusOK)
}

/////////////////////////////////////////////////////////////////////

//--------Review----------//
// Read
// GetReview
func GetReview(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	review, err := reviews.GetAll(ctx)

	if err != nil {
		fmt.Println(err)
	}

	utils.ResponseJSON(w, review, http.StatusOK)
}

// Create
// PostReview
func PostReview(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var review models.Reviews

	if err := json.NewDecoder(r.Body).Decode(&review); err != nil {
		utils.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}

	if err := reviews.Insert(ctx, review); err != nil {
		utils.ResponseJSON(w, err, http.StatusInternalServerError)
		return
	}

	res := map[string]string{
		"status": "Succesfully",
	}

	utils.ResponseJSON(w, res, http.StatusCreated)

}

// UpdateReview
func UpdateReview(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var review models.Reviews

	if err := json.NewDecoder(r.Body).Decode(&review); err != nil {
		utils.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}

	var idStar = ps.ByName("id")

	if err := reviews.Update(ctx, review, idStar); err != nil {
		utils.ResponseJSON(w, err, http.StatusInternalServerError)
		return
	}

	res := map[string]string{
		"status": "Succesfully",
	}

	utils.ResponseJSON(w, res, http.StatusCreated)
}

// Delete
// DeleteReview
func DeleteReview(w http.ResponseWriter, _ *http.Request, ps httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var idReview = ps.ByName("id")

	if err := movie.Delete(ctx, idReview); err != nil {
		kesalahan := map[string]string{
			"error": fmt.Sprintf("%v", err),
		}
		utils.ResponseJSON(w, kesalahan, http.StatusInternalServerError)
		return
	}

	res := map[string]string{
		"status": "Succesfully",
	}

	utils.ResponseJSON(w, res, http.StatusOK)
}

/////////////////////////////////////////////////////////////////////
