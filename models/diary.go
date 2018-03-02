package models

import (
  "errors"
  "time"

  "golang-mood-tracker/db"
  "golang-mood-tracker/forms"
)

//diary ...
type Diary struct {
  ID        int64    `db:"id, primarykey, autoincrement" json:"id"`
  UserID    int64    `db:"user_id" json:"-"`
  Title     string   `db:"title" json:"title"`
  Content   string   `db:"content" json:"content"`
  UpdatedAt int64    `db:"updated_at" json:"updated_at"`
  CreatedAt int64    `db:"created_at" json:"created_at"`
  User      *JSONRaw `db:"user" json:"user"`
}

//DiaryModel ...
type DiaryModel struct{}

//Create ...
func (m DiaryModel) Create(userID int64, form forms.DiaryForm) (diaryID int64, err error) {
  getDb := db.GetDB()

  userModel := new(UserModel)

  checkUser, err := userModel.One(userID)

  if err != nil && checkUser.ID > 0 {
    return 0, errors.New("User doesn't exist")
  }

  _, err = getDb.Exec("INSERT INTO public.diary(user_id, title, content, updated_at, created_at) VALUES($1, $2, $3, $4, $5) RETURNING id", userID, form.Title, form.Content, time.Now().Unix(), time.Now().Unix())

  if err != nil {
    return 0, err
  }

  diaryID, err = getDb.SelectInt("SELECT id FROM public.diary WHERE user_id=$1 ORDER BY id DESC LIMIT 1", userID)

  return diaryID, err
}

// //One ...
// func (m DiaryModel) One(userID, id int64) (diary diary, err error) {
//   err = db.GetDB().SelectOne(&diary, "SELECT a.id, a.title, a.content, a.updated_at, a.created_at, json_build_object('id', u.id, 'name', u.name, 'email', u.email) AS user FROM public.diary a LEFT JOIN public.user u ON a.user_id = u.id WHERE a.user_id=$1 AND a.id=$2 LIMIT 1", userID, id)
//   return diary, err
// }

// //All ...
// func (m DiaryModel) All(userID int64) (diaries []diary, err error) {
//   _, err = db.GetDB().Select(&diarys, "SELECT a.id, a.title, a.content, a.updated_at, a.created_at, json_build_object('id', u.id, 'name', u.name, 'email', u.email) AS user FROM public.diary a LEFT JOIN public.user u ON a.user_id = u.id WHERE a.user_id=$1 ORDER BY a.id DESC", userID)
//   return diarys, err
// }

// //Update ...
// func (m DiaryModel) Update(userID int64, id int64, form forms.DiaryForm) (err error) {
//   _, err = m.One(userID, id)

//   if err != nil {
//     return errors.New("diary not found")
//   }

//   _, err = db.GetDB().Exec("UPDATE public.diary SET title=$1, content=$2, updated_at=$3 WHERE id=$4", form.Title, form.Content, time.Now().Unix(), id)

//   return err
// }

// //Delete ...
// func (m DiaryModel) Delete(userID, id int64) (err error) {
//   _, err = m.One(userID, id)

//   if err != nil {
//     return errors.New("diary not found")
//   }

//   _, err = db.GetDB().Exec("DELETE FROM public.diary WHERE id=$1", id)

//   return err
// }