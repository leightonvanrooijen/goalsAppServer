package graph

import (
	"fmt"

	"github.com/goalsApp/server/db"
	"github.com/graph-gophers/graphql-go"
)

type RootResolver struct {
	db db.DB
}

type GoalResolver struct {
	db   db.DB
	goal *db.Goal
}

func (r *RootResolver) Goals() ([]*GoalResolver, error) {
	goals, err := r.db.GetAllGoals()
	if err != nil {
		return nil, fmt.Errorf("error [%s]: Could not fetch users", err)
	}

	GoalResolvers := make([]*GoalResolver, len(goals))
	for i := range goals {
		GoalResolvers[i] = &GoalResolver{
			db:   r.db,
			goal: &goals[i],
		}
	}
	return GoalResolvers, nil
}

// Gets goal for user
func (r *RootResolver) Goal(args struct{ ID int32 }) (*GoalResolver, error) {
	goal := r.db.GetGoal(args.ID)

	return &GoalResolver{db: r.db, goal: goal}, nil
}

// Creates a goal
func (g *GoalResolver) Create(args *struct {
	Name        string
	Description string
	ParentID    int32
}) *GoalResolver {
	goal := db.Goal{
		Name:        args.Name,
		Description: args.Description,
		ParentID:    args.ParentID,
	}

	dbGoal := g.db.CreateGoal(&goal)

	return &GoalResolver{db: g.db, goal: dbGoal}
}

// Gets all child goals for an id
func (g *RootResolver) SubGoals(args struct{ ID int32 }) ([]*GoalResolver, error) {
	goals, err := g.db.GetChildren(uint(args.ID))
	if err != nil {
		return nil, fmt.Errorf("error [%s]: Could not fetch child goals", err)
	}

	goalRxs := make([]*GoalResolver, len(goals))
	for i := range goals {
		goalRxs[i] = &GoalResolver{
			db:   g.db,
			goal: &goals[i],
		}
	}
	return goalRxs, nil
}

// Gets name for user
func (g *GoalResolver) Name() *string {
	return &g.goal.Name
}

// Gets name for user
func (g *GoalResolver) Id() *graphql.ID {
	goalID := graphql.ID(fmt.Sprint(&g.goal.ID))

	return &goalID
}

// Gets name for user
func (g *GoalResolver) Description() *string {
	return &g.goal.Description
}

// Gets name for user
func (g *GoalResolver) ParentID() *int32 {
	return &g.goal.ParentID
}

// Resolves child goals from parent
func (g *GoalResolver) Goals() ([]*GoalResolver, error) {
	goals, err := g.db.GetChildren(g.goal.ID)
	if err != nil {
		return nil, fmt.Errorf("error [%s]: Could not fetch child goals", err)
	}

	goalRxs := make([]*GoalResolver, len(goals))
	for i := range goals {
		goalRxs[i] = &GoalResolver{
			db:   g.db,
			goal: &goals[i],
		}
	}
	return goalRxs, nil
}

// dfmn34c5sy
