package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	dto "ways/dto/result"
	toppingsdto "ways/dto/toppings"
	"ways/models"
	"ways/repositories"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
)

//Declare handler struct here ....
type handlerTopping struct {
	ToppingRepository repositories.ToppingRepository
}

// Create `path_file` Global variable here ...
var path_filee = "http://localhost:5000/uploads/"

//Declare HandlerProduct function here ...
func HandlerTopping(ToppingRepository repositories.ToppingRepository) *handlerTopping {
	return &handlerTopping{ToppingRepository}
}

//Declare FindTopping method here ...
func (h *handlerTopping) FindTopping(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	toppings, err := h.ToppingRepository.FindTopping()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	// Create Embed Path File on Image property here ...
	for i, p := range toppings {
		toppings[i].Image = path_filee + p.Image
	}

	w.WriteHeader(http.StatusOK)
	//response := dto.SuccessResult{Code: http.StatusOK, Data: products}
	json.NewEncoder(w).Encode(toppings)
}

//Declare GetTopping method here ...
func (h *handlerTopping) GetTopping(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	//menyimpan nilai
	toppings, err := h.ToppingRepository.GetTopping(id)
	//menyimpan nilai
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	// Create Embed Path File on Image property here ...
	toppings.Image = path_filee + toppings.Image

	w.WriteHeader(http.StatusOK)
	//response := dto.SuccessResult{Code: http.StatusOK, Data: products}
	json.NewEncoder(w).Encode(toppings)
}

//Declare CreateTopping method here ...
func (h *handlerTopping) CreateTopping(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// get data user token
	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	userId := int(userInfo["id"].(float64))

	//Get dataFile from middleware and store to filename variable here ...
	dataContex := r.Context().Value("dataFile") // add this code
	filename := dataContex.(string)             // add this code

	price, _ := strconv.Atoi(r.FormValue("price")) //convert. //integer yg diconvert
	//category_id, _ := strconv.Atoi(r.FormValue("category_id"))
	request := toppingsdto.CreateToppingRequest{
		Name:  r.FormValue("name"),
		Price: price,
		//CategoryID: category_id,
	}

	validator := validator.New()
	err := validator.Struct(request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}
	//cara untuk memvalidasi data sudah dikirimkan atau belum, sesuai dengan product_request
	topping := models.Topping{
		Name:   request.Name,
		Price:  request.Price,
		Image:  filename,
		UserID: userId,
	}
	top, err := h.ToppingRepository.CreateTopping(topping)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertToppingResponse(top)}
	json.NewEncoder(w).Encode(response)

}

//Declare DeleteTopping method here ...
func (h *handlerTopping) DeleteTopping(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	topping, err := h.ToppingRepository.GetTopping(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	data, err := h.ToppingRepository.DeleteTopping(topping)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertToppingResponse(data)}
	json.NewEncoder(w).Encode(response)
}

//Declare convertResponse function here ...
func convertToppingResponse(u models.Topping) toppingsdto.ToppingResponse {
	return toppingsdto.ToppingResponse{
		ID:    u.ID,
		Name:  u.Name,
		Price: u.Price,
		Image: u.Image,
	}
}
