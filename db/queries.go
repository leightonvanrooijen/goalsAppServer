package db

import "fmt"

// get all goals
// TODO pre-load all comments
func (ctx DB) GetAllGoals() ([]Goal, error) {
	var goals []Goal
	if err := ctx.db.Find(&goals); err.Error != nil {
		return nil, err.Error
	}
	return goals, nil
}

// find goal by id
func (ctx DB) GetGoalById(id int) (*Goal, error) {
	var goal Goal
	if err := ctx.db.First(&goal, "ID = ?", id); err.Error != nil {
		return nil, err.Error
	}
	return &goal, nil
}

// find children by parent id
func (ctx DB) GetChildren(id uint) ([]Goal, error) {
	var goals []Goal
	if err := ctx.db.Find(&goals, "parent_id = ?", id); err.Error != nil {
		return nil, err.Error
	}
	return goals, nil
}

// create Goal associates parent
func (ctx DB) CreateGoal(Goal *Goal) *Goal {
	if err := ctx.db.Create(&Goal); err.Error != nil {
		fmt.Printf("Could not create goal: %s.", Goal.Name)
	}

	return Goal
}

// delete Goal and children
func (ctx DB) GetGoal(id int32) *Goal {
	var goal Goal
	var parent Goal
	ctx.db.First(&goal, "ID = ?", id).Association("Parent").Find(&parent)
	fmt.Println(parent)
	return &goal
}