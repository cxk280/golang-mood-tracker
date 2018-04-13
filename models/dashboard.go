package models

import (
  "errors"
  "time"

  "golang-mood-tracker/db"
  "golang-mood-tracker/forms"
)

// Define Dashboard struct
type Dashboard struct {
  ID        int64    `db:"id, primarykey, autoincrement" json:"id"`
  UserID    int64    `db:"user_id" json:"-"`
  Title     string   `db:"title" json:"title"`
  Content   string   `db:"content" json:"content"`
  UpdatedAt int64    `db:"updated_at" json:"updated_at"`
  CreatedAt int64    `db:"created_at" json:"created_at"`
  User      *JSONRaw `db:"user" json:"user"`
}

// Define Dashboard model struct
type DashboardModel struct{}

// Get all
func (m DashboardModel) All(userID int64) (dashboards []Dashboard, err error) {
  _, err = db.GetDB().Select(&dashboards, "SELECT a.id, a.title, a.content, a.updated_at, a.created_at, json_build_object('id', u.id, 'name', u.name, 'email', u.email) AS user FROM public.dashboard a LEFT JOIN public.user u ON a.user_id = u.id WHERE a.user_id=$1 ORDER BY a.id DESC", userID)
  return dashboards, err
}