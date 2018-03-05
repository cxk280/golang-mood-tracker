package forms

//DiaryForm ...
type DiaryForm struct {
  Feeling   string `form:"feeling" json:"feeling" binding:"required,max=100"`
  Notes     string `form:"notes" json:"notes" binding:"required,max=1000"`
}
