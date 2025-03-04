package handler

import (
	"MICROSERVICE-CRUD/data"
	"encoding/json"
	"log"
	"net/http"
)

type Products struct{
	thisLogger *log.Logger
	//log emssages to the output 
}

func NewProducts(thisLogger *log.Logger) *Products{
	return &Products{thisLogger}
}

func (h *Products) GetProducts(w http.ResponseWriter, r *http.Request) {
	h.thisLogger.Println("Handle get products")
    listofProducts := data.GetProducts()
//here we are encoding the req and writing response using w 
	errWhileEncoding := listofProducts.ToJson(w)

	if errWhileEncoding != nil{
		http.Error(w, errWhileEncoding.Error(), http.StatusInternalServerError)
	}


}
func AddSingleProduct(thisProd *Product){
	thisProd.ID = getNextIdFromDb()

	productList = append(productList, thisProd)
}

func getNextIdFromDb()int {
	currListofProdInDB := productList[len(productList)-1]
	return (currListofProdInDB.ID +1)
}


func (h *Products) addProduct(w http.ResponseWriter, r *http.Request){
	h.thisLogger.Println("Handle add product")
	thisprod := &data.Product{}

	errWhileMarshall := thisprod.FromJson(r.Body)

	if errWhileMarshall != nil {
		http.Error(w, errWhileMarshall.Error(), http.StatusBadRequest)
	}
	data.AddSingleProduct(thisprod)
}

func (h *Products) ServeHTTP(w http.ResponseWriter, r *http.Request){
	//method to make use of hhttp Hnalder interface
	if r.Method == http.MethodGet{
		h.GetProducts(w,r)
		return
	} 
	if r.Method == http.MethodPost{
		h.addProduct(w,r)
		return
	}

	h.thisLogger.Println("serving Products List ...")

	 listofProducts := data.GetProducts()
	 jsonRepOfListOfProd, errWhileMarhsal := json.Marshal(listofProducts)

	 if errWhileMarhsal != nil{
		http.Error(w,"unable to marshal the  JSON ", http.StatusInternalServerError)
	 }
	 w.Write(jsonRepOfListOfProd)

	}

