package userModel

//o model define como as nossas entidades serão, posteriormente, adicionarei integração direta com o banco de dados por meio do GORM
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
}