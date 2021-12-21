package models

type (
	//users
	Users struct {
		User_id  int    `json:"user_id"`
		Username string `json:"username"`
		Password string `json:"password"`
		Role     string `json:"role"`
	}

	//movie
	Movie struct {
		Movie_id     int    `json:"movie_id"`
		Name         string `json:"name"`
		Description  string `json:"description"`
		Release_date string `json:"release_date"`
		Image_url    string `json:"image_url"`
		Creator_id   int    `json:"creator_id"`
		Rate_id      int    `json:"rate_id"`
		Star_id      int    `json:"star_id"`
	}

	//genre_category
	Genre_category struct {
		Movie_id int `json:"movie_id"`
		Genre_id int `json:"genre_id"`
	}

	//genre
	Genre struct {
		Genre_id int    `json:"genre_id"`
		Name     string `json:"name"`
	}

	//creators
	Creators struct {
		Creator_id int    `json:"release_date"`
		Name       string `json:"name"`
	}

	//stars
	Stars struct {
		Star_id int    `json:"star_id"`
		Name    string `json:"name"`
	}

	//rating
	Rating struct {
		Rate_id int     `json:"rate_id"`
		Rate    float64 `json:"rate"`
		User_id int     `json:"user_id"`
	}

	//reviews
	Reviews struct {
		Reviews_id int    `json:"reviews_id"`
		Review     string `json:"review"`
		User_id    int    `json:"user_id"`
		Movie_id   int    `json:"movie_id"`
	}

	//feedback
	Feedback struct {
		Feedback_id int    `json:"feedback_id"`
		Feedback    string `json:"feedback"`
		User_id     int    `json:"user_id"`
		Reviews_id  int    `json:"reviews_id"`
	}
)
