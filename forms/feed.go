package forms

//FeedForm ...
type FeedForm struct {
  Title   string `form:"title" json:"title" binding:"required,max=100"`
  Content string `form:"content" json:"content" binding:"required,max=1000"`
}