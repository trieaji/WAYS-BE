package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	productsdto "ways/dto/products"
	dto "ways/dto/result"
	"ways/models"
	"ways/repositories"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
	// Write this code
)

//Declare handler struct here ....
type handler struct {
	ProductRepository repositories.ProductRepository
}

// Create `path_file` Global variable here ...
var path_file = "http://localhost:5000/uploads/"

//Declare HandlerProduct function here ...
func HandlerProduct(ProductRepository repositories.ProductRepository) *handler {
	return &handler{ProductRepository}
}

//Declare FindProduct method here ...
func (h *handler) FindProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	products, err := h.ProductRepository.FindProduct()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	// Create Embed Path File on Image property here ...
	for i, p := range products {
		products[i].Image = path_file + p.Image
	}

	w.WriteHeader(http.StatusOK)
	//response := dto.SuccessResult{Code: http.StatusOK, Data: products}
	json.NewEncoder(w).Encode(products)
}

//Declare GetProduct method here ...
func (h *handler) GetProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	//menyimpan nilai
	product, err := h.ProductRepository.GetProduct(id)
	//menyimpan nilai
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	// Create Embed Path File on Image property here ...
	product.Image = path_file + product.Image

	w.WriteHeader(http.StatusOK)
	//response := dto.SuccessResult{Code: http.StatusOK, Data: products}
	json.NewEncoder(w).Encode(product)
}

//Declare CreateProduct method here ...
func (h *handler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// get data user token
	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	userId := int(userInfo["id"].(float64))

	//Get dataFile from middleware and store to filename variable here ...
	dataContex := r.Context().Value("dataFile") // add this code
	filename := dataContex.(string)             // add this code

	price, _ := strconv.Atoi(r.FormValue("price")) //convert. //integer yg diconvert
	qty, _ := strconv.Atoi(r.FormValue("qty"))     //convert. //integer yg diconvert
	//category_id, _ := strconv.Atoi(r.FormValue("category_id"))
	request := productsdto.CreateProductRequest{
		Name:  r.FormValue("name"),
		Price: price,
		Qty:   qty,
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
	product := models.Product{
		Name:   request.Name,
		Price:  request.Price,
		Image:  filename,
		Qty:    request.Qty,
		UserID: userId,
	}
	data, err := h.ProductRepository.CreateProduct(product)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponse(data)}
	json.NewEncoder(w).Encode(response)

}

//Declare DeleteProduct method here ...
func (h *handler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	user, err := h.ProductRepository.GetProduct(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	data, err := h.ProductRepository.DeleteProduct(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponse(data)}
	json.NewEncoder(w).Encode(response)
}

//Declare convertResponse function here ...
func convertResponse(u models.Product) productsdto.ProductResponse {
	return productsdto.ProductResponse{
		ID:    u.ID,
		Name:  u.Name,
		Price: u.Price,
		Image: u.Image,
	}
}
