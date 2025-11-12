package handler

import (
	"encoding/json"
	"gos/domain"
	"gos/service"
	"net/http"
)

type handler struct { // guarda las dependencias que necesita para responder requests
	service service.Service
}

func NewHandler(s service.Service) *handler {
	return &handler{service: s}
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// GET /pokemons -> listar pokemones
	if r.Method == http.MethodGet && r.URL.Path == "/pokemons" {
		pokes, err := h.service.GetPokemons(r.Context())
		if err != nil {
			http.Error(w, "error obteniendo pokemones: "+err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(pokes)
		return
	}

	// POST /pokemons -> crear pokemon
	if r.Method == http.MethodPost && r.URL.Path == "/pokemons" {
		// Mover aquí la lógica de CreatePokemon
		var pokemon domain.Pokemon
		if err := json.NewDecoder(r.Body).Decode(&pokemon); err != nil {
			http.Error(w, "JSON inválido", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		createdPokemon, err := h.service.Post(r.Context(), &pokemon)
		if err != nil {
			http.Error(w, "error creando pokemon", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		_ = json.NewEncoder(w).Encode(createdPokemon)
		return
	}

	http.Error(w, "ruta no encontrada", http.StatusNotFound)
}

/*func (h *handler) GetPersonas(w http.ResponseWriter, r *http.Request) {
	/*
	personas, err := h.service.Get(r.Context())
	if err != nil {
		http.Error(w, "error obteniendo personas", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(personas)
}*/

/*
func (h *handler) CreatePokemon(w http.ResponseWriter, r *http.Request) {

	pokemon, err := h.service.Post(r.Context(), nil)
	if err != nil {
		http.Error(w, "error creando pokemon", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(pokemon)
}*/

func (h *handler) DeletePokemon(w http.ResponseWriter, r *http.Request) {
	// Aquí deberías obtener el ID del Pokémon a eliminar, por ejemplo, desde los parámetros de la URL
	id := 1 // Esto es solo un ejemplo, deberías obtener el ID real
	err := h.service.Delete(r.Context(), id)
	if err != nil {
		http.Error(w, "error eliminando pokemon", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
func (h *handler) PatchPokemon(w http.ResponseWriter, r *http.Request) {
	// Aquí deberías obtener el ID del Pokémon a actualizar, por ejemplo, desde los parámetros de la URL
	id := 1 // Esto es solo un ejemplo, deberías obtener el ID real
	err := h.service.Patch(r.Context(), id, nil)
	if err != nil {
		http.Error(w, "error actualizando pokemon", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
