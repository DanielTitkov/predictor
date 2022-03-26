// Code generated by entc, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/DanielTitkov/predictor/internal/repository/entgo/ent/challenge"
	"github.com/DanielTitkov/predictor/internal/repository/entgo/ent/prediction"
	"github.com/DanielTitkov/predictor/internal/repository/entgo/ent/schema"
	"github.com/DanielTitkov/predictor/internal/repository/entgo/ent/user"
	"github.com/google/uuid"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	challengeMixin := schema.Challenge{}.Mixin()
	challengeMixinFields0 := challengeMixin[0].Fields()
	_ = challengeMixinFields0
	challengeFields := schema.Challenge{}.Fields()
	_ = challengeFields
	// challengeDescCreateTime is the schema descriptor for create_time field.
	challengeDescCreateTime := challengeMixinFields0[0].Descriptor()
	// challenge.DefaultCreateTime holds the default value on creation for the create_time field.
	challenge.DefaultCreateTime = challengeDescCreateTime.Default.(func() time.Time)
	// challengeDescUpdateTime is the schema descriptor for update_time field.
	challengeDescUpdateTime := challengeMixinFields0[1].Descriptor()
	// challenge.DefaultUpdateTime holds the default value on creation for the update_time field.
	challenge.DefaultUpdateTime = challengeDescUpdateTime.Default.(func() time.Time)
	// challenge.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	challenge.UpdateDefaultUpdateTime = challengeDescUpdateTime.UpdateDefault.(func() time.Time)
	// challengeDescContent is the schema descriptor for content field.
	challengeDescContent := challengeFields[1].Descriptor()
	// challenge.ContentValidator is a validator for the "content" field. It is called by the builders before save.
	challenge.ContentValidator = func() func(string) error {
		validators := challengeDescContent.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(content string) error {
			for _, fn := range fns {
				if err := fn(content); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// challengeDescDescription is the schema descriptor for description field.
	challengeDescDescription := challengeFields[2].Descriptor()
	// challenge.DescriptionValidator is a validator for the "description" field. It is called by the builders before save.
	challenge.DescriptionValidator = challengeDescDescription.Validators[0].(func(string) error)
	// challengeDescID is the schema descriptor for id field.
	challengeDescID := challengeFields[0].Descriptor()
	// challenge.DefaultID holds the default value on creation for the id field.
	challenge.DefaultID = challengeDescID.Default.(func() uuid.UUID)
	predictionMixin := schema.Prediction{}.Mixin()
	predictionMixinFields0 := predictionMixin[0].Fields()
	_ = predictionMixinFields0
	predictionFields := schema.Prediction{}.Fields()
	_ = predictionFields
	// predictionDescCreateTime is the schema descriptor for create_time field.
	predictionDescCreateTime := predictionMixinFields0[0].Descriptor()
	// prediction.DefaultCreateTime holds the default value on creation for the create_time field.
	prediction.DefaultCreateTime = predictionDescCreateTime.Default.(func() time.Time)
	// predictionDescUpdateTime is the schema descriptor for update_time field.
	predictionDescUpdateTime := predictionMixinFields0[1].Descriptor()
	// prediction.DefaultUpdateTime holds the default value on creation for the update_time field.
	prediction.DefaultUpdateTime = predictionDescUpdateTime.Default.(func() time.Time)
	// prediction.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	prediction.UpdateDefaultUpdateTime = predictionDescUpdateTime.UpdateDefault.(func() time.Time)
	// predictionDescID is the schema descriptor for id field.
	predictionDescID := predictionFields[0].Descriptor()
	// prediction.DefaultID holds the default value on creation for the id field.
	prediction.DefaultID = predictionDescID.Default.(func() uuid.UUID)
	userMixin := schema.User{}.Mixin()
	userMixinFields0 := userMixin[0].Fields()
	_ = userMixinFields0
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescCreateTime is the schema descriptor for create_time field.
	userDescCreateTime := userMixinFields0[0].Descriptor()
	// user.DefaultCreateTime holds the default value on creation for the create_time field.
	user.DefaultCreateTime = userDescCreateTime.Default.(func() time.Time)
	// userDescUpdateTime is the schema descriptor for update_time field.
	userDescUpdateTime := userMixinFields0[1].Descriptor()
	// user.DefaultUpdateTime holds the default value on creation for the update_time field.
	user.DefaultUpdateTime = userDescUpdateTime.Default.(func() time.Time)
	// user.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	user.UpdateDefaultUpdateTime = userDescUpdateTime.UpdateDefault.(func() time.Time)
	// userDescName is the schema descriptor for name field.
	userDescName := userFields[1].Descriptor()
	// user.NameValidator is a validator for the "name" field. It is called by the builders before save.
	user.NameValidator = userDescName.Validators[0].(func(string) error)
	// userDescEmail is the schema descriptor for email field.
	userDescEmail := userFields[2].Descriptor()
	// user.EmailValidator is a validator for the "email" field. It is called by the builders before save.
	user.EmailValidator = userDescEmail.Validators[0].(func(string) error)
	// userDescAdmin is the schema descriptor for admin field.
	userDescAdmin := userFields[3].Descriptor()
	// user.DefaultAdmin holds the default value on creation for the admin field.
	user.DefaultAdmin = userDescAdmin.Default.(bool)
	// userDescID is the schema descriptor for id field.
	userDescID := userFields[0].Descriptor()
	// user.DefaultID holds the default value on creation for the id field.
	user.DefaultID = userDescID.Default.(func() uuid.UUID)
}
