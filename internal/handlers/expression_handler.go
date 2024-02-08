package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/Ki4EH/yndx-distributed_calculator/internal/logger"
	"github.com/Ki4EH/yndx-distributed_calculator/internal/models"
	"github.com/Ki4EH/yndx-distributed_calculator/pkg/database"
	"github.com/Ki4EH/yndx-distributed_calculator/pkg/utils"
	"github.com/rs/xid"
	"net/http"
)

func AddExpressionHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		return
	}
	eh := models.NewExpression()
	log, _ := logger.NewFileLogger()
	if err := json.NewDecoder(r.Body).Decode(&eh); err != nil {
		log.Error(fmt.Sprintf("[JSON/Decode] Error decode json data in expression got: %v", err))
		http.Error(w, "Failed decode JSON", http.StatusInternalServerError)
		return
	}

	//проверяем есть ли ID в бд,
	//либо если пользователь не передал ID генерируем новый
	if eh.ID == "" || (database.Find(eh.ID) == false) {
		if !utils.IsValidExpression(eh.Value) {
			http.Error(w, "Expression invalid", http.StatusBadRequest)
			return
		}

		if eh.ID == "" {
			id := xid.New()
			eh.ID = fmt.Sprintf("%s", id)
		}
		w.WriteHeader(http.StatusOK)
		database.Add(eh.ID, eh)
		//передать данные в оркестор !!!
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Request already computing"))
}
