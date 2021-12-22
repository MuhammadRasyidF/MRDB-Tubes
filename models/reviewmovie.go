package models

type (
	//users
	Tb_users struct {
		UserId   int    `json:"userid"`
		Username string `json:"username"`
		Password string `json:"password"`
		Role     string `json:"role"`
	}

	//movie
	Tb_movies struct {
		MovieId     int    `json:"movieid"`
		Name        string `json:"name"`
		Description string `json:"description"`
		ReleaseDate string `json:"releasedate"`
		ImageUrl    string `json:"imageurl"`
	}

	//genre
	Tb_genres struct {
		GenreId int    `json:"genreid"`
		Name    string `json:"name"`
	}

	//creators
	Tb_creators struct {
		CreatorId int    `json:"creatorid"`
		Name      string `json:"name"`
	}

	//stars
	Tb_stars struct {
		StarId int    `json:"starid"`
		Name   string `json:"name"`
	}

	//rating
	Rating struct {
		Rate    int `json:"rate"`
		UserId  int `json:"userid"`
		MovieId int `json:"movieid"`
	}

	//comment
	Tb_comment struct {
		CommentId     int    `json:"commentid"`
		Comment       string `json:"mcomment"`
		UserId        int    `json:"userid"`
		MovieId       int    `json:"movieid"`
		ParentComment int    `json:"parentcomment"`
	}

	//like the comment
	LikeComment struct {
		CommentId int `json:"commentid"`
		UserId    int `json:"userid"`
	}

	//many to many table here
	//movie_creator
	Movie_creator struct {
		MovieId   int `json:"movieid"`
		CreatorId int `json:"creatorid"`
	}

	//movie-genre
	Movie_genre struct {
		MovieId int `json:"movieid"`
		GenreId int `json:"genreid"`
	}

	//movie-star
	Movie_star struct {
		MovieId int `json:"movieid"`
		StarId  int `json:"starid"`
	}
)
