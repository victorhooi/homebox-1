package repo

import (
	"context"

	"github.com/google/uuid"
	"github.com/hay-kot/content/backend/ent"
)

type EntGroupRepository struct {
	db *ent.Client
}

func (r *EntGroupRepository) Create(ctx context.Context, name string) (*ent.Group, error) {
	return r.db.Group.Create().SetName(name).Save(ctx)
}

func (r *EntGroupRepository) GetOneId(ctx context.Context, id uuid.UUID) (*ent.Group, error) {
	return r.db.Group.Get(ctx, id)
}