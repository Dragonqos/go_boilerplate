package middleware

import (
	"fmt"
	"encoding/json"
	"github.com/Dragonqos/go_boilerplate/api/core/cipher"
	"net/http"
	"os"
)

const tokenKey = "tenant"

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		cryptedToken := r.Header.Get(tokenKey)
		if len(cryptedToken) == 0 {
			cryptedToken = r.URL.Query().Get(tokenKey)
		}

		if len(cryptedToken) > 0 {

			key := os.Getenv("BP_ENCRYPTION_KEY")
			iv := os.Getenv("BP_ENCRYPTION_IV")

			decrypted := cipher.Decrypt([]byte(key), []byte(iv), cryptedToken)

			fmt.Println(decrypted)
			var data map[string]interface{}
			if err := json.Unmarshal([]byte(decrypted), &data); err != nil {
				panic(err)
			}
			tenantId := data["tenant_id"]
			if tenantId != nil {
				next.ServeHTTP(w, r)
				return
			}

			fmt.Println("Wrong data format come with request")
		}

		http.Error(w, "Forbidden", http.StatusForbidden)
	})
}