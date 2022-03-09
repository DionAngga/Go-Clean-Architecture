package helper

type CreateUser struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Nasabah string `json:"nasabah"`
	Email   string `json:"email"`
}

// type UserFormatter struct {
// 	ID         int    `json:"id"`
// 	Name       string `json:"name"`
// 	Occupation string `json:"occupation"`
// 	Email      string `json:"email"`
// 	Token      string `json:"token"`
// 	ImageURL   string `json:"image_url"`
// }

// func FormatUser(user entity.User, token string) UserFormatter {
// 	formatter := UserFormatter{
// 		ID:         user.ID,
// 		Name:       user.Name,
// 		Occupation: user.Occupation,
// 		Email:      user.Email,
// 		Token:      token,
// 		ImageURL:   user.AvatarFileName,
// 	}

// 	return formatter
// }
