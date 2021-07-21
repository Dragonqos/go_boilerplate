package routes

import (
	"github.com/Dragonqos/go_boilerplate/api/core/cipher"
	"github.com/Dragonqos/go_boilerplate/api/model"
	"encoding/json"
	"net/http"
	"os"
)

func Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var res model.ResponseResult

	key := os.Getenv("BP_ENCRYPTION_KEY")
	iv := os.Getenv("BP_ENCRYPTION_IV")

	crypted := r.URL.Query().Get("tenant")
	decoded := cipher.Decrypt([]byte(key), []byte(iv), crypted)

	res.Result = decoded
	json.NewEncoder(w).Encode(res)
}
