package auth

import (
	"context"
	"net/http"
)

type contextKey string

const (
    userIDKey contextKey = "userID"
    roleKey   contextKey = "role"
)

func AuthMiddleWare(next http.Handler) http.Handler{
	return  http.HandlerFunc(func(w http.ResponseWriter, r* http.Request){
		cookie,err := r.Cookie("access_token")
	if err != nil {
				http.Error(
					w,
					"authentication required",
					http.StatusUnauthorized,
				)
				return
			}

			claims, err := ValidateAccessToken(
				cookie.Value,
			)

			if err != nil {
				http.Error(
					w,
					"invalid or expired token",
					http.StatusUnauthorized,
				)
				return
			}

			ctx := context.WithValue(
				r.Context(),
				userIDKey,
				claims.Subject,
			)
// Context A
//      │
//      │ WithValue(userIDKey, "abc-123")
//      ▼
// Context B
			ctx = context.WithValue(
				ctx,
				roleKey,
				claims.Role,
			)
// 			Context
// ├── userID → abc-123
// └── role   → admin

			next.ServeHTTP(
				w,
				r.WithContext(ctx),
			)
		},
	)
}


func ReqRole(role string) func (next http.Handler) http.Handler{
	// func RequireRole(role string) func(http.Handler) http.Handler -> Give me a role, and I will give you a middleware that requires that role.
	return func(next http.Handler) http.Handler{
		return http.HandlerFunc( func(w http.ResponseWriter, r *http.Request){
			userRole, ok := GetRole(r.Context())
			if !ok {
				http.Error(w,"authenrication required",http.StatusUnauthorized,)
				return
			}
			if userRole != role{
				http.Error(w, "you are not authorized",http.StatusForbidden)
				return
			}
			next.ServeHTTP(w,r)
		},
	)
	}
}

// func RequireRole(role string) func(http.Handler) http.Handler {

//     // STEP 1: role is remembered here

//     middleware := func(next http.Handler) http.Handler {

//         // STEP 2: this is the actual middleware

//         handler := http.HandlerFunc(func(
//             w http.ResponseWriter,
//             r *http.Request,
//         ) {

//             userRole, ok := GetRole(r.Context())

//             if !ok {
//                 http.Error(
//                     w,
//                     "authentication required",
//                     http.StatusUnauthorized,
//                 )
//                 return
//             }

//             if userRole != role {
//                 http.Error(
//                     w,
//                     "you are not authorized",
//                     http.StatusForbidden,
//                 )
//                 return
//             }

//             next.ServeHTTP(w, r)
//         })

//         return handler
//     }

//     return middleware
// }

func GetUserID(
	ctx context.Context,
) (string, bool) {

	userID, ok := ctx.Value(
		userIDKey,
	).(string)

	return userID, ok
}

func GetRole(ctx context.Context)(string,bool){
	role,ok := ctx.Value(roleKey).(string)
	return role,ok
}