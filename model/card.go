package model 

type(
	Card struct {
		Id int 
		UserId int
		Picture string
		Name string	
	}

	ReadCardsResponse struct {
		Cards []Card
	}
)
