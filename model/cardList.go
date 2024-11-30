package model 

type(
	Card struct {
		Id int 
		UserId int
		Picture string
		Name string	
	}

	ReadCardRequest struct {
		UserId int
	}

	ReadCardResponse struct {
		Cards []Card
	}
)
