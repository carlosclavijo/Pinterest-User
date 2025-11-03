package handlers

import (
	"context"
	"github.com/carlosclavijo/Pinterest-User/internal/application/user/commands"
	"github.com/carlosclavijo/Pinterest-User/internal/application/user/dto"
	"github.com/carlosclavijo/Pinterest-User/internal/application/user/mappers"
	"github.com/carlosclavijo/Pinterest-User/internal/domain/abstractions"
	"github.com/carlosclavijo/Pinterest-User/internal/domain/shared"
	"github.com/carlosclavijo/Pinterest-User/internal/domain/user"
	"github.com/google/uuid"
)

func (h *UserHandler) HandleUpdate(ctx context.Context, cmd commands.UpdateUserCommand) (*dto.UserResponse, error) {
	var (
		err      error
		username shared.Username
		email    shared.Email
		password shared.Password
		gender   shared.Gender
		birth    shared.BirthDate
		country  shared.Country
		language shared.Language
		phone    *shared.Phone
		webSite  *shared.Website
	)

	if cmd.Id == uuid.Nil {
		return nil, users.ErrIdNilUser
	}

	exist, err := h.repository.ExistsById(ctx, cmd.Id)
	if err != nil {
		return nil, err
	} else if !exist {
		return nil, users.ErrNotFoundUser
	}

	usr, err := h.repository.GetById(ctx, cmd.Id)
	if err != nil {
		return nil, err
	}

	usr.AggregateRoot = abstractions.NewAggregateRoot(cmd.Id)

	if cmd.FirstName != nil {
		err = usr.ChangeFirstName(*cmd.FirstName)
		if err != nil {
			return nil, err
		}
	}

	if cmd.LastName != nil {
		err = usr.ChangeLastName(*cmd.LastName)
		if err != nil {
			return nil, err
		}
	}

	if cmd.UserName != nil {
		if usr.Username().String() != *cmd.UserName {
			username, err = shared.NewUsername(*cmd.UserName)
			if err != nil {
				return nil, err
			}

			err = usr.ChangeUsername(username)
			if err != nil {
				return nil, err
			}
		}
	}

	if cmd.Email != nil {
		if usr.Email().String() != *cmd.Email {
			email, err = shared.NewEmail(*cmd.Email)
			if err != nil {
				return nil, err
			}

			err = usr.ChangeEmail(email)
			if err != nil {
				return nil, err
			}
		}
	}

	if cmd.Password != nil {
		password, err = hashPassword(*cmd.Password)
		if err != nil {
			return nil, err
		}

		err = usr.ChangePassword(password)
		if err != nil {
			return nil, err
		}
	}

	if cmd.Gender != nil {
		gender, err = shared.ParseGender(*cmd.Gender)
		if err != nil {
			return nil, err
		}

		err = usr.ChangeGender(gender)
		if err != nil {
			return nil, err
		}
	}

	if cmd.Birth != nil {
		birth, err = shared.NewBirthDate(*cmd.Birth)
		if err != nil {
			return nil, err
		}

		err = usr.ChangeBirth(birth)
		if err != nil {
			return nil, err
		}
	}

	if cmd.Country != nil {
		country, err = shared.ParseCountry(*cmd.Country)
		if err != nil {
			return nil, err
		}

		err = usr.ChangeCountry(country)
		if err != nil {
			return nil, err
		}
	}

	if cmd.Language != nil {
		language, err = shared.ParseLanguage(*cmd.Language)
		if err != nil {
			return nil, err
		}

		err = usr.ChangeLanguage(language)
		if err != nil {
			return nil, err
		}
	}

	if cmd.Phone != nil {
		if *cmd.Phone != "" {
			phone, err = shared.NewPhone(cmd.Phone)
			if err != nil {
				return nil, err
			}
			usr.ChangePhone(phone)
		} else {
			usr.ChangePhone(nil)
		}
	}

	if cmd.Information != nil {
		if *cmd.Information != "" {
			err = usr.ChangeInformation(cmd.Information)
			if err != nil {
				return nil, err
			}
		} else {
			err = usr.ChangeInformation(nil)
			if err != nil {
				return nil, err
			}
		}
	}

	if cmd.ProfilePic != nil {
		if *cmd.ProfilePic != "" {
			usr.ChangeProfilePic(cmd.ProfilePic)
		} else {
			usr.ChangeProfilePic(nil)
		}
	}

	if cmd.Website != nil {
		if *cmd.Website != "" {
			webSite, err = shared.NewWebSite(cmd.Website)
			if err != nil {
				return nil, err
			}
			usr.ChangeWebSite(webSite)
		} else {
			usr.ChangeWebSite(nil)
		}
	}

	if cmd.Visibility != nil {
		usr.ChangeVisibility(*cmd.Visibility)
	}

	usr.Update()

	if usr, err = h.repository.Update(ctx, usr); err != nil {
		return nil, err
	}

	userDto := mappers.MapToUserDTO(usr)
	userResponse := mappers.MapToUserResponse(userDto, usr.LastLoginAt(), usr.CreatedAt(), usr.UpdatedAt(), usr.DeletedAt())

	return userResponse, nil
}
