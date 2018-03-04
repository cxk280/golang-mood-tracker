package forms

//DiaryForm ...
type DiaryForm struct {
  Title   string `form:"title" json:"title" binding:"required,max=100"`
  Content string `form:"content" json:"content" binding:"required,max=1000"`
}

// //SigninForm ...
// type SigninForm struct {
//   Email    string `form:"email" json:"email" binding:"required,email"`
//   Password string `form:"password" json:"password" binding:"required"`
// }
