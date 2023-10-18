package handler

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/megadata-dev/go-snmp-olt-c320/internal/usecase"
	"github.com/megadata-dev/go-snmp-olt-c320/pkg/utils"
	"net/http"
	"strconv"
)

// Response format for JSON
type Response struct {
	Data []map[string]string `json:"data"`
}

type OnuHandler struct {
	ponUsecase usecase.OnuUseCase
}

func NewOnuHandler(ponUsecase usecase.OnuUseCase) *OnuHandler {
	return &OnuHandler{
		ponUsecase: ponUsecase,
	}
}

func (o *OnuHandler) List(w http.ResponseWriter, r *http.Request) {
	onuInfoList, err := o.ponUsecase.GetAllONUInfo(r.Context())
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Mengonversi hasil ke JSON
	jsonResult, err := json.Marshal(onuInfoList)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(jsonResult)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func (o *OnuHandler) GetByGtGoIDAndPonID(w http.ResponseWriter, r *http.Request) {
	// Ambil nilai parameter "gtgo_id" dan "id" dari URL
	gtGoID := chi.URLParam(r, "gtgo_id")
	ponID := chi.URLParam(r, "id")

	// convert string to int
	gtGoIDInt, err := strconv.Atoi(gtGoID)
	if err != nil {
		utils.ErrorBadRequest(w, fmt.Errorf("gtgo_id must be 0 or 1"))
		return
	}
	if gtGoIDInt != 0 && gtGoIDInt != 1 {
		utils.ErrorBadRequest(w, fmt.Errorf("gtgo_id must be 0 or 1"))
		return
	}

	// convert string to int
	ponIDInt, err := strconv.Atoi(ponID)
	if err != nil {
		utils.ErrorBadRequest(w, fmt.Errorf("pon_id must be between 1 and 8"))
		return
	}
	if ponIDInt < 1 || ponIDInt > 8 {
		utils.ErrorBadRequest(w, fmt.Errorf("pon_id must be between 1 and 8"))
		return
	}

	// Panggil usecase untuk mendapatkan data
	onuInfoList, err := o.ponUsecase.GetByGtGoIDAndPonID(r.Context(), gtGoIDInt, ponIDInt)

	if err != nil {
		utils.ErrorInternalServerError(w, fmt.Errorf("cannot get data from snmp"))
		return
	}

	// Mengonversi hasil ke JSON sesuai dengan struktur WebResponse
	response := utils.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   onuInfoList,
	}

	utils.SendJSONResponse(w, http.StatusOK, response)
}
