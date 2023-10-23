package handler

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/megadata-dev/go-snmp-olt-zte-c320/internal/usecase"
	"github.com/megadata-dev/go-snmp-olt-zte-c320/internal/utils"
	"github.com/megadata-dev/go-snmp-olt-zte-c320/pkg/pagination"
	"net/http"
	"strconv"
)

type OnuHandlerInterface interface {
	GetByBoardIDAndPonID(w http.ResponseWriter, r *http.Request)
	GetByBoardIDAndPonIDWithPaginate(w http.ResponseWriter, r *http.Request)
	GetByBoardIDPonIDAndOnuID(w http.ResponseWriter, r *http.Request)
	GetEmptyOnuID(w http.ResponseWriter, r *http.Request)
	UpdateEmptyOnuID(w http.ResponseWriter, r *http.Request)
	GetEmptyOnuIDQueue(w http.ResponseWriter, r *http.Request)
}

type OnuHandler struct {
	ponUsecase usecase.OnuUseCase
}

func NewOnuHandler(ponUsecase usecase.OnuUseCase) *OnuHandler {
	return &OnuHandler{ponUsecase: ponUsecase}
}

func (o *OnuHandler) GetByBoardIDAndPonID(w http.ResponseWriter, r *http.Request) {
	boardID := chi.URLParam(r, "board_id") // 1 or 2
	ponID := chi.URLParam(r, "pon_id")     // 1 - 8

	boardIDInt, err := strconv.Atoi(boardID) // convert string to int

	// Validate boardIDInt value and return error 400 if boardIDInt is not 1 or 2
	if err != nil || (boardIDInt != 1 && boardIDInt != 2) {
		utils.ErrorBadRequest(w, fmt.Errorf("invalid 'board_id' parameter. It must be 1 or 2")) // error 400
		return
	}

	ponIDInt, err := strconv.Atoi(ponID) // convert string to int

	// Validate ponIDInt value and return error 400 if ponIDInt is not between 1 and 8
	if err != nil || ponIDInt < 1 || ponIDInt > 8 {
		utils.ErrorBadRequest(w, fmt.Errorf("invalid 'pon_id' parameter. It must be between 1 and 8")) // error 400
		return
	}

	query := r.URL.Query() // Get query parameters from the request

	//Validate query parameters and return error 400 if query parameters is not "onu_id" or empty query parameters
	if len(query) > 0 && query["onu_id"] == nil {
		utils.ErrorBadRequest(w, fmt.Errorf("invalid query parameter")) // error 400
		return
	}

	// Call usecase to get data from SNMP
	onuInfoList, err := o.ponUsecase.GetByBoardIDAndPonID(r.Context(), boardIDInt, ponIDInt)
	if err != nil {
		utils.ErrorInternalServerError(w, fmt.Errorf("cannot get data from snmp")) // error 500
		return
	}

	/*
		Validate onuInfoList value
		If onuInfoList is empty, return error 404
	*/

	if len(onuInfoList) == 0 {
		utils.ErrorNotFound(w, fmt.Errorf("data not found")) // error 404
		return
	}

	// Convert result to JSON format according to WebResponse structure
	response := utils.WebResponse{
		Code:   http.StatusOK, // 200
		Status: "OK",          // "OK"
		Data:   onuInfoList,   // data
	}

	utils.SendJSONResponse(w, http.StatusOK, response) // 200

}

func (o *OnuHandler) GetByBoardIDAndPonIDWithPaginate(w http.ResponseWriter, r *http.Request) {

	boardID := chi.URLParam(r, "board_id") // 1 or 2
	ponID := chi.URLParam(r, "pon_id")     // 1 - 8

	// Get page and page size parameters from the request
	pageIndex, pageSize := pagination.GetPaginationParametersFromRequest(r)

	boardIDInt, err := strconv.Atoi(boardID) // convert string to int

	// Validate boardIDInt value and return error 400 if boardIDInt is not 1 or 2
	if err != nil || (boardIDInt != 1 && boardIDInt != 2) {
		utils.ErrorBadRequest(w, fmt.Errorf("invalid 'board_id' parameter. It must be 1 or 2")) // error 400
		return
	}

	ponIDInt, err := strconv.Atoi(ponID) // convert string to int

	// Validate ponIDInt value and return error 400 if ponIDInt is not between 1 and 8
	if err != nil || ponIDInt < 1 || ponIDInt > 8 {
		utils.ErrorBadRequest(w, fmt.Errorf("invalid 'pon_id' parameter. It must be between 1 and 8")) // error 400
		return
	}

	item, count := o.ponUsecase.GetByBoardIDAndPonIDWithPagination(r.Context(), boardIDInt, ponIDInt, pageIndex, pageSize)

	/*
		Validate item value
		If item is empty, return error 404
	*/

	if len(item) == 0 {
		utils.ErrorNotFound(w, fmt.Errorf("data not found")) // error 404
		return
	}

	// Convert result to JSON format according to Pages structure
	pages := pagination.New(pageIndex, pageSize, count)

	// Convert result to JSON format according to WebResponse structure
	responsePagination := pagination.Pages{
		Code:      http.StatusOK, // 200
		Status:    "OK",          // "OK"
		Page:      pages.Page,
		PageSize:  pages.PageSize,
		PageCount: pages.PageCount,
		TotalRows: pages.TotalRows,
		Data:      item,
	}

	utils.SendJSONResponse(w, http.StatusOK, responsePagination) // 200
}

func (o *OnuHandler) GetByBoardIDPonIDAndOnuID(w http.ResponseWriter, r *http.Request) {

	boardID := chi.URLParam(r, "board_id") // 1 or 2
	ponID := chi.URLParam(r, "pon_id")     // 1 - 8
	onuID := chi.URLParam(r, "onu_id")     // 1 - 128

	boardIDInt, err := strconv.Atoi(boardID) // convert string to int

	// Validate boardIDInt value and return error 400 if boardIDInt is not 1 or 2
	if err != nil || (boardIDInt != 1 && boardIDInt != 2) {
		utils.ErrorBadRequest(w, fmt.Errorf("invalid 'board_id' parameter. It must be 1 or 2")) // error 400
		return
	}

	ponIDInt, err := strconv.Atoi(ponID) // convert string to int

	// Validate ponIDInt value and return error 400 if ponIDInt is not between 1 and 8
	if err != nil || ponIDInt < 1 || ponIDInt > 8 {
		utils.ErrorBadRequest(w, fmt.Errorf("invalid 'pon_id' parameter. It must be between 1 and 8")) // error 400
		return
	}

	onuIDInt, err := strconv.Atoi(onuID) // convert string to int

	// Validate onuIDInt value and return error 400 if onuIDInt is not between 1 and 128
	if err != nil || onuIDInt < 1 || onuIDInt > 128 {
		utils.ErrorBadRequest(w, fmt.Errorf("invalid 'onu_id' parameter. It must be between 1 and 128")) // error 400
		return
	}

	// Call usecase to get data from SNMP
	onuInfoList, err := o.ponUsecase.GetByBoardIDPonIDAndOnuID(r.Context(), boardIDInt, ponIDInt, onuIDInt)

	if err != nil {
		utils.ErrorInternalServerError(w, fmt.Errorf("cannot get data from snmp")) // error 500
		return
	}

	/*
		Validate onuInfoList value
		If onuInfoList.Board, onuInfoList.PON, and onuInfoList.ID is 0, return error 404
		example: http://localhost:8080/board/1/pon/1/onu/129
	*/

	if onuInfoList.Board == 0 && onuInfoList.PON == 0 && onuInfoList.ID == 0 {
		utils.ErrorNotFound(w, fmt.Errorf("data not found")) // error 404
		return
	}

	// Convert result to JSON format according to WebResponse structure
	response := utils.WebResponse{
		Code:   http.StatusOK, // 200
		Status: "OK",          // "OK"
		Data:   onuInfoList,   // data
	}

	utils.SendJSONResponse(w, http.StatusOK, response) // 200
}

func (o *OnuHandler) GetEmptyOnuID(w http.ResponseWriter, r *http.Request) {

	boardID := chi.URLParam(r, "board_id") // 1 or 2
	ponID := chi.URLParam(r, "pon_id")     // 1 - 8

	boardIDInt, err := strconv.Atoi(boardID) // convert string to int

	// Validate boardIDInt value and return error 400 if boardIDInt is not 1 or 2
	if err != nil || (boardIDInt != 1 && boardIDInt != 2) {
		utils.ErrorBadRequest(w, fmt.Errorf("invalid 'board_id' parameter. It must be 1 or 2")) // error 400
		return
	}

	ponIDInt, err := strconv.Atoi(ponID) // convert string to int

	// Validate ponIDInt value and return error 400 if ponIDInt is not between 1 and 8
	if err != nil || ponIDInt < 1 || ponIDInt > 8 {
		utils.ErrorBadRequest(w, fmt.Errorf("invalid 'pon_id' parameter. It must be between 1 and 8")) // error 400
		return
	}

	// Call usecase to get data from SNMP
	onuIDEmptyList, err := o.ponUsecase.GetEmptyOnuID(r.Context(), boardIDInt, ponIDInt)

	if err != nil {
		utils.ErrorInternalServerError(w, fmt.Errorf("cannot get data from snmp")) // error 500
		return
	}

	// Convert result to JSON format according to WebResponse structure
	response := utils.WebResponse{
		Code:   http.StatusOK,  // 200
		Status: "OK",           // "OK"
		Data:   onuIDEmptyList, // data
	}

	utils.SendJSONResponse(w, http.StatusOK, response) // 200
}

func (o *OnuHandler) UpdateEmptyOnuID(w http.ResponseWriter, r *http.Request) {
	boardID := chi.URLParam(r, "board_id") // 1 or 2
	ponID := chi.URLParam(r, "pon_id")     // 1 - 8

	boardIDInt, err := strconv.Atoi(boardID) // convert string to int

	// Validate boardIDInt value and return error 400 if boardIDInt is not 1 or 2
	if err != nil || (boardIDInt != 1 && boardIDInt != 2) {
		utils.ErrorBadRequest(w, fmt.Errorf("invalid 'board_id' parameter. It must be 0 or 1")) // error 400
		return
	}

	ponIDInt, err := strconv.Atoi(ponID) // convert string to int

	// Validate ponIDInt value and return error 400 if ponIDInt is not between 1 and 8
	if err != nil || ponIDInt < 1 || ponIDInt > 8 {
		utils.ErrorBadRequest(w, fmt.Errorf("invalid 'pon_id' parameter. It must be between 1 and 8")) // error 400
		return
	}

	// Call usecase to get data from SNMP
	err = o.ponUsecase.UpdateEmptyOnuID(r.Context(), boardIDInt, ponIDInt)

	if err != nil {
		utils.ErrorInternalServerError(w, fmt.Errorf("cannot get data from snmp")) // error 500
		return
	}

	// Convert result to JSON format according to WebResponse structure
	response := utils.WebResponse{
		Code:   http.StatusOK,                 // 200
		Status: "OK",                          // "OK"
		Data:   "Success Update Empty ONU_ID", // data
	}

	utils.SendJSONResponse(w, http.StatusOK, response) // 200
}

func (o *OnuHandler) GetEmptyOnuIDQueue(w http.ResponseWriter, r *http.Request) {

	boardID := chi.URLParam(r, "board_id") // 1 or 2
	ponID := chi.URLParam(r, "pon_id")     // 1 - 8

	boardIDInt, err := strconv.Atoi(boardID) // convert string to int

	// Validate boardIDInt value and return error 400 if boardIDInt is not 1 or 2
	if err != nil || (boardIDInt != 1 && boardIDInt != 2) {
		utils.ErrorBadRequest(w, fmt.Errorf("invalid 'board_id' parameter. It must be 1 or 2")) // error 400
		return
	}

	ponIDInt, err := strconv.Atoi(ponID) // convert string to int

	// Validate ponIDInt value and return error 400 if ponIDInt is not between 1 and 8
	if err != nil || ponIDInt < 1 || ponIDInt > 8 {
		utils.ErrorBadRequest(w, fmt.Errorf("invalid 'pon_id' parameter. It must be between 1 and 8")) // error 400
		return
	}

	// Call usecase to get data from SNMP
	onuIDEmptyList, err := o.ponUsecase.GetEmptyOnuIDQueue(r.Context(), boardIDInt, ponIDInt)

	if err != nil {
		utils.ErrorInternalServerError(w, fmt.Errorf("cannot get data from snmp")) // error 500
		return
	}

	// Convert result to JSON format according to WebResponse structure
	response := utils.WebResponse{
		Code:   http.StatusOK,  // 200
		Status: "OK",           // "OK"
		Data:   onuIDEmptyList, // data
	}

	utils.SendJSONResponse(w, http.StatusOK, response) // 200
}
