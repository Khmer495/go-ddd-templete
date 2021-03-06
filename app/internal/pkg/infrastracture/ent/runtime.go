// Code generated by entc, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/Khmer495/go-templete/internal/pkg/infrastracture/ent/schema"
	"github.com/Khmer495/go-templete/internal/pkg/infrastracture/ent/team"
	"github.com/Khmer495/go-templete/internal/pkg/infrastracture/ent/teamuser"
	"github.com/Khmer495/go-templete/internal/pkg/infrastracture/ent/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	teamMixin := schema.Team{}.Mixin()
	teamMixinFields1 := teamMixin[1].Fields()
	_ = teamMixinFields1
	teamFields := schema.Team{}.Fields()
	_ = teamFields
	// teamDescCreatedAt is the schema descriptor for created_at field.
	teamDescCreatedAt := teamMixinFields1[0].Descriptor()
	// team.DefaultCreatedAt holds the default value on creation for the created_at field.
	team.DefaultCreatedAt = teamDescCreatedAt.Default.(func() time.Time)
	// teamDescUpdatedAt is the schema descriptor for updated_at field.
	teamDescUpdatedAt := teamMixinFields1[1].Descriptor()
	// team.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	team.DefaultUpdatedAt = teamDescUpdatedAt.Default.(func() time.Time)
	// team.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	team.UpdateDefaultUpdatedAt = teamDescUpdatedAt.UpdateDefault.(func() time.Time)
	teamuserMixin := schema.TeamUser{}.Mixin()
	teamuserMixinFields0 := teamuserMixin[0].Fields()
	_ = teamuserMixinFields0
	teamuserFields := schema.TeamUser{}.Fields()
	_ = teamuserFields
	// teamuserDescCreatedAt is the schema descriptor for created_at field.
	teamuserDescCreatedAt := teamuserMixinFields0[0].Descriptor()
	// teamuser.DefaultCreatedAt holds the default value on creation for the created_at field.
	teamuser.DefaultCreatedAt = teamuserDescCreatedAt.Default.(func() time.Time)
	userMixin := schema.User{}.Mixin()
	userMixinFields1 := userMixin[1].Fields()
	_ = userMixinFields1
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescCreatedAt is the schema descriptor for created_at field.
	userDescCreatedAt := userMixinFields1[0].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the created_at field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(func() time.Time)
	// userDescUpdatedAt is the schema descriptor for updated_at field.
	userDescUpdatedAt := userMixinFields1[1].Descriptor()
	// user.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	user.DefaultUpdatedAt = userDescUpdatedAt.Default.(func() time.Time)
	// user.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	user.UpdateDefaultUpdatedAt = userDescUpdatedAt.UpdateDefault.(func() time.Time)
}
