package auth

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{
		service: service,
	}
}

type AuthRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	// Role string `json:"role"`
}

func (h *Handler) Signup(w http.ResponseWriter, r *http.Request) {
	var req AuthRequest
	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		return
	}

	if req.Email == "" || req.Password == "" {
		http.Error(w, "email and password are required", http.StatusBadRequest)
		return
	}

	_, err = h.service.Signup(r.Context(), req.Email, req.Password)
	if err != nil {
		if err == ErrEmailAlreadyExists {
			http.Error(w, "email already exists", http.StatusConflict)
			return
		}

		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode("user created successfully")

}

func (h *Handler) Signin(
	w http.ResponseWriter,
	r *http.Request,
) {

	var req AuthRequest

	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		http.Error(
			w,
			"invalid JSON",
			http.StatusBadRequest,
		)
		return
	}

	accessToken, refreshToken, err := h.service.Signin(
		r.Context(),
		req.Email,
		req.Password,
	)

	if err != nil {
		if err == ErrInvalidCredentials {
			http.Error(
				w,
				"invalid credentials",
				http.StatusUnauthorized,
			)
			return
		}

		http.Error(
			w,
			"internal server error",
			http.StatusInternalServerError,
		)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "access_token",
		Value:    accessToken,
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
		MaxAge:   15 * 60,
	})

	http.SetCookie(w, &http.Cookie{
		Name:     "refresh_token",
		Value:    refreshToken,
		Path:     "/refresh",
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
		MaxAge:   7 * 24 * 60 * 60,
	})

	json.NewEncoder(w).Encode(
		map[string]string{
			"message": "signed in successfully",
		},
	)
}

func (h *Handler) Signout(
	w http.ResponseWriter,
	r *http.Request,
) {

	http.SetCookie(w, &http.Cookie{
		Name:     "access_token",
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
		MaxAge:   -1,
	})

	json.NewEncoder(w).Encode(
		map[string]string{
			"message": "signed out successfully",
		},
	)
}

func (h *Handler) Profile(
	w http.ResponseWriter,
	r *http.Request,
) {

	userID, ok := GetUserID(
		r.Context(),
	)

	if !ok {
		http.Error(
			w,
			"user not authenticated",
			http.StatusUnauthorized,
		)
		return
	}

	json.NewEncoder(w).Encode(
		map[string]string{
			"user_id": userID,
		},
	)
}

func (s *Service) Refresh(
    ctx context.Context,
    refreshToken string,
) (string, error) {

    storedToken, ok := s.refreshTokenStore.Get(refreshToken)

    if !ok {
        return "", errors.New("invalid refresh token")
    }

    foundUser, err := s.userRepo.FindbyID(
        ctx,
        storedToken.UserID,
    )

    if err != nil {
        return "", err
    }

    accessToken, err := GenerateAccessToken(
        foundUser.ID,
        foundUser.Role,
    )

    if err != nil {
        return "", err
    }

    return accessToken, nil
}

func (h *Handler) Refresh(
    w http.ResponseWriter,
    r *http.Request,
) {

    cookie, err := r.Cookie("refresh_token")

    if err != nil {
        http.Error(
            w,
            "refresh token required",
            http.StatusUnauthorized,
        )
        return
    }

    accessToken, err := h.service.Refresh(
        r.Context(),
        cookie.Value,
    )

    if err != nil {
        http.Error(
            w,
            "invalid refresh token",
            http.StatusUnauthorized,
        )
        return
    }

    http.SetCookie(w, &http.Cookie{
        Name:     "access_token",
        Value:    accessToken,
        Path:     "/",
        HttpOnly: true,
        Secure:   false,
        SameSite: http.SameSiteLaxMode,
        MaxAge:   15 * 60,
    })

    json.NewEncoder(w).Encode(map[string]string{
        "message": "access token refreshed",
    })
}